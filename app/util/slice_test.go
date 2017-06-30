package util_test

import (
	"testing"

	"github.com/agustin-sarasua/pimbay/app/util"
)

func filter(l string) bool {
	return l == "1234"
}

func TestFilterStringSlice(t *testing.T) {
	var lines = []string{"1234", "asdfdsaf", "1234", "asdfasdf"}
	linesFiltered := util.FilterStringSlice(lines, filter)

	if len(linesFiltered) != 2 {
		t.Error("expected 2")
	}
}
