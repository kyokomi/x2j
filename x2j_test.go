package x2j

import (
	"fmt"
	"testing"

	"github.com/tealeg/xlsx"
)

func TestConvert(t *testing.T) {
	xFile, err := xlsx.OpenFile("./RogueGameMaster.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	x2j := New()
	res, err := x2j.ToUpper().Convert(xFile)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(res))
}
