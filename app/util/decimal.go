package util

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/shopspring/decimal"
)

type BigDecimal struct {
	value decimal.Decimal
}

// UnmarshalJSON Decimal
func (d *BigDecimal) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	fmt.Println(s)
	if s == "0" {
		d.value = decimal.Decimal{}
		return
	}
	d.value, err = decimal.NewFromString(s)
	return
}

// MarshalJSON Decimal
func (d *BigDecimal) MarshalJSON() ([]byte, error) {
	if d == nil {
		return []byte("0"), nil
	}
	cd := d.value
	return []byte(fmt.Sprintf("%s", cd.String())), nil
}

// GetBSON implements bson.Getter.
func (d BigDecimal) GetBSON() (interface{}, error) {
	f, err := strconv.ParseFloat(d.value.String(), 64)
	if err != nil {
		panic(err)
	}
	return f, nil
}

// SetBSON implements bson.Setter.
func (d BigDecimal) SetBSON(raw bson.Raw) error {
	var num string
	raw.Unmarshal(&num)
	fmt.Println(num)
	return d.UnmarshalJSON([]byte(num))
}
