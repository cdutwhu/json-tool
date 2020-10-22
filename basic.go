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
func Cvt2XML(jstr string) string {
	var jmap interface{}
	json.Unmarshal([]byte(jstr), &jmap)
	bytes, err := mxj.AnyXmlIndent(jmap, "", "    ", "")
	failOnErr("%v", err)
	xstr := string(bytes)
	xstr = sReplaceAll(xstr, "<>", "")
	xstr = sReplaceAll(xstr, "</>", "")
	return sTrim(xstr, " \t\n")
}
