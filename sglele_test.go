package jsontool

import (
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestJSONBlkCont(t *testing.T) {
	defer misc.TrackTime(time.Now())
	jstr := `{
		"Activity": {		
		  "-RefId": "C27E1FCF-C163-485F-BEF0-F36F18A0493A",
		  "ActivityTime": {
			"CreationDate": "2002-06-15",
			"DueDate": "2002-09-12",
			"Duration": {
			  "#content": 30,
			  "-Units": "minute"
			},
			"FinishDate": "2002-09-12",
			"StartDate": "2002-09-10"
		  },
		  "ActivityWeight": 5,
		  "AssessmentRefId": "03EDB29E-8116-B450-0435-FA87E42A0AD2",
		  "Evaluation": {
			"-EvaluationType": "Inline",
			"Description": "Students should be able to correctly identify all major characters."
		  },
		  "LearningResources": {
			"LearningResourceRefId": [
			  "B7337698-BF6D-B193-7F79-A07B87211B93"
			]
		  },
		  "LearningStandards": {
			"LearningStandardItemRefId": [
			  "9DB15CEA-B2C5-4F66-94C3-7D0A0CAEDDA4"
			]
		  },
		  "MaxAttemptsAllowed": 3,
		  "Points": 50,
		  "Preamble": "This is a very funny comedy - students should have passing familiarity with Shakespeare",
		  "SourceObjects": {
			"SourceObject": [
			  {
				"#content": "A71ADBD3-D93D-A64B-7166-E420D50EDABC",
				"-SIF_RefObject": "Lesson"
			  }
			]
		  },
		  "Title": "Shakespeare Essay - Much Ado About Nothing"
		}
	  }`

	val, ok := SglEleAttrVal(jstr, "RefId", "-")
	fPln(val, ok)

	name, cont := SglEleBlkCont(jstr)
	fPln("root", name)
	fPln(cont)
	fPln(" ------------------------- ")

	out := MkSglEleBlk(name, "~~~", true)
	fPln(out)

	mav := map[string]interface{}{"a": "b", "c": 12}
	xstr := Cvt2XML(out, mav)
	fPln(xstr)

	// names, values := JSONBreakBlkCont(cont)
	// for i, name := range names {
	// 	fPln(i, name, ":", values[i])
	// }
}
