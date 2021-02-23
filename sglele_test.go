package jsontool

import (
	"os"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestJSONBlkCont(t *testing.T) {
	defer misc.TrackTime(time.Now())

	bytes, err := os.ReadFile("./data/Activity.json")
	failOnErr("%v", err)
	jsonstr := string(bytes)

	val, ok := SglEleAttrVal(jsonstr, "RefId", "-")
	fPln(val, ok)

	name, cont := SglEleBlkCont(jsonstr)
	fPln("root", name)
	fPln(cont)
	fPln(" ------------------------- ")

	out := MkSglEleBlk(name, "~~~", true)
	fPln(out)

	mav := map[string]interface{}{"a": "b", "c": 12}
	xmlstr := Cvt2XML(out, mav)
	fPln(xmlstr)

	// names, values := JSONBreakBlkCont(cont)
	// for i, name := range names {
	// 	fPln(i, name, ":", values[i])
	// }
}
