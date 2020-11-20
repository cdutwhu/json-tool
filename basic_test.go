package jsontool

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestIsValid(t *testing.T) {
	dir := "./examples/"
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if jsonfile := info.Name(); sHasSuffix(jsonfile, ".json") {
			fPln("--->", jsonfile)

			bytes, _ := ioutil.ReadFile(dir + jsonfile)
			jstr := string(bytes)

			if !IsValid(jstr) {
				ioutil.WriteFile(fSf("debug_%s.json", jsonfile), []byte(jstr), 0666)
				panic("error on MkJSON")
			}

			//if jsonfile == "CensusCollection_0.xml" {
			// ioutil.WriteFile(fSf("record_%s.json", jsonfile), []byte(jstr), 0666)
			//}
		}
		return nil
	})
}
