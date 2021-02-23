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
func Fmt(jsonstr, indent string) string {
	// jsonmap := make(map[string]interface{})
	var jsonmap interface{}
	json.Unmarshal([]byte(jsonstr), &jsonmap)
	bytes, err := json.MarshalIndent(&jsonmap, "", indent)
	failOnErr("%v", err)
	return string(bytes)
}

// Cvt2XML :
func Cvt2XML(jsonstr string, mav map[string]interface{}) string {
	var jsonmap interface{}
	json.Unmarshal([]byte(jsonstr), &jsonmap)
	bytes, err := mxj.AnyXmlIndent(jsonmap, "", "    ", "")
	failOnErr("%v", err)
	xmlstr := string(bytes)
	xmlstr = sReplaceAll(xmlstr, "<>", "")
	xmlstr = sReplaceAll(xmlstr, "</>", "")
	xmlstr = sTrim(xmlstr, " \t\n")

	attrs := []string{}
	for a, v := range mav {
		attrs = append(attrs, fSf(`%s="%v"`, a, v))
	}
	if p := sIndex(xmlstr, ">"); len(attrs) > 0 {
		xmlstr = xmlstr[:p] + " " + sJoin(attrs, " ") + xmlstr[p:]
	}

	return xmlstr
}
