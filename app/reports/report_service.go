package reports

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

	"cloud.google.com/go/storage"

	"mime/multipart"

	"github.com/ledongthuc/pdf"
	uuid "github.com/satori/go.uuid"
)

func ReadPdf(path string) (string, error) {
	r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	var textBuilder bytes.Buffer
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		textBuilder.WriteString(p.GetPlainText("\n"))
	}
	return textBuilder.String(), nil
}

func PrintValidReportLines(r Report) {
	ls := r.GetTransactionLines()
	for _, l := range ls {
		t, ld, bn, c, a := r.ParseLineDetail(l)
		fmt.Println(t, " ", ld, " ", bn, " ", c, a)
	}
}

// readFile reads the named file in Google Cloud Storage.
func ReadFile(fileName string, ctx context.Context) (string, error) {
	// io.WriteString(d.w, "\nAbbreviated file content (first line and last 1K):\n")

	rc, err := StorageBucket.Object(fileName).NewReader(ctx)
	if err != nil {
		fmt.Errorf("readFile: unable to open file from bucket %q, file %q: %v", StorageBucketName, fileName, err)
		return "", err
	}
	defer rc.Close()

	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Errorf("readFile: unable to read data from bucket %q, file %q: %v", StorageBucketName, fileName, err)
		return "", err
	}

	var tmpFileName bytes.Buffer
	tmpFileName.WriteString("/tmp/")
	tmpFileName.WriteString(fileName)
	fmt.Println(tmpFileName.String())

	err = ioutil.WriteFile(tmpFileName.String(), slurp, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpFileName.String()) // clean up

	s, _ := ReadPdf(tmpFileName.String())
	//_ := fmt.Sprintf("%s", slurp)

	return s, nil
	// fmt.Fprintf(d.w, "%s\n", bytes.SplitN(slurp, []byte("\n"), 2)[0])
	// if len(slurp) > 1024 {
	// 	fmt.Fprintf(d.w, "...%s\n", slurp[len(slurp)-1024:])
	// } else {
	// 	fmt.Fprintf(d.w, "%s\n", slurp)
	// }
}

func ProcessReport(f multipart.File, fh *multipart.FileHeader) <-chan string {
	out := make(chan string)
	go func() {
		b1 := make([]byte, 5)
		n1, err := f.Read(b1)
		fmt.Printf("%d bytes: %s\n", n1, string(b1))
		if err != nil {
			return
		}
		name := uuid.NewV4().String() + path.Ext(fh.Filename)

		var tmpFileName bytes.Buffer
		tmpFileName.WriteString("/tmp/")
		tmpFileName.WriteString(name)
		fmt.Println(tmpFileName.String())

		err = ioutil.WriteFile(tmpFileName.String(), b1, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer os.Remove(tmpFileName.String()) // clean up

		s, _ := ReadPdf(tmpFileName.String())
		out <- s
	}()
	return out
}

func PushReportToCloudStorage(f multipart.File, fh *multipart.FileHeader) <-chan string {
	out := make(chan string)
	go func() {
		if StorageBucket == nil {
			return
		}
		b1 := make([]byte, 5)
		n1, err := f.Read(b1)
		fmt.Printf("%d bytes: %s\n", n1, string(b1))

		if err != nil {
			return
		}
		// random filename, retaining existing extension.
		name := uuid.NewV4().String() + path.Ext(fh.Filename)

		ctx := context.Background()
		w := StorageBucket.Object(name).NewWriter(ctx)
		w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
		w.ContentType = fh.Header.Get("Content-Type")

		// Entries are immutable, be aggressive about caching (1 day).
		w.CacheControl = "public, max-age=86400"

		if _, err := io.Copy(w, f); err != nil {
			return
		}
		if err := w.Close(); err != nil {
			return
		}
		out <- name
	}()
	return out
}
