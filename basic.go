package jsontool

import (
	"encoding/json"

	"github.com/clbanning/mxj"
)

// IsValid :
func IsValid(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// Fmt :
func Fmt(jstr, indent string) string {
	// jmap := make(map[string]interface{})
	var jmap interface{}
	json.Unmarshal([]byte(jstr), &jmap)
	bytes, err := json.MarshalIndent(&jmap, "", indent)
	failOnErr("%v", err)
	return string(bytes)
}

// Cvt2XML :
func Cvt2XML(jstr string, mav map[string]interface{}) string {
	var jmap interface{}
	json.Unmarshal([]byte(jstr), &jmap)
	bytes, err := mxj.AnyXmlIndent(jmap, "", "    ", "")
	failOnErr("%v", err)
	xstr := string(bytes)
	xstr = sReplaceAll(xstr, "<>", "")
	xstr = sReplaceAll(xstr, "</>", "")
	xstr = sTrim(xstr, " \t\n")

	attrs := []string{}
	for a, v := range mav {
		attrs = append(attrs, fSf(`%s="%v"`, a, v))
	}
	if p := sIndex(xstr, ">"); len(attrs) > 0 {
		xstr = xstr[:p] + " " + sJoin(attrs, " ") + xstr[p:]
	}

	return xstr
}
