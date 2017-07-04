package util

import (
	"fmt"
	"testing"
)

/*
A string is in effect a read-only slice of bytes.
A string holds arbitrary bytes
It is not required to hold Unicode text, UTF-8 text, or any other predefined format
It is exactly equivalent to a slice of bytes.
*/

//String literal that uses the \xNN notation
const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func TestPrint(t *testing.T) {
	// Because some of the bytes in our sample string are not valid ASCII,
	// not even valid UTF-8, printing the string directly will produce ugly output.
	fmt.Println(sample)
	// ��=� ⌘

	// To find out what that string really holds, we need to take it apart
	// and examine the pieces. There are several ways to do this.
	// The most obvious is to loop over its contents and pull out the
	// bytes individually, as in this for loop:

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	// This is the output from the byte-by-byte loop:
	// bd b2 3d bc 20 e2 8c 98
}

func TestStringSymbolPrint(t *testing.T) {
	/*
		To avoid any confusion, we create a "raw string", enclosed by back quotes,
		so it can contain only literal text. (Regular strings, enclosed by double quotes,
		can contain escape sequences as we showed above.)
	*/
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")
}

func TestRunes(t *testing.T) {
	const nihongo = "日本語"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
