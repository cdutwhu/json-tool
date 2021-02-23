package jsontool

import (
	"os"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestJSONBreakArrCont(t *testing.T) {
	defer misc.TrackTime(time.Now())

	bytes, err := os.ReadFile("./data/Activities.json")
	failOnErr("%v", err)
	jsonstr := string(bytes)

	values, ok := BreakArr(jsonstr)
	fPln(ok)
	for _, v := range values {
		fPln(v)
	}
}

func TestJSONBreakBlkContV2(t *testing.T) {
	defer misc.TrackTime(time.Now())
	if bytes, err := os.ReadFile("./why.json"); err == nil {
		// jsonstr := JSONBlkFmt(string(bytes), "  ")
		jsonstr := string(bytes)
		_, cont := SglEleBlkCont(jsonstr)
		names, values := BreakMulEleBlkV2(cont)
		for i, name := range names {
			fPln(MkSglEleBlk(name, values[i], false))
			fPln(" ------------------------------------------ ")
		}
	}
}
