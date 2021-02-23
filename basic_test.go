package jsontool

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsValid(t *testing.T) {
	dir := "./data/"
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		failOnErrWhen(info == nil, "%v", err)
		if jsonfile := info.Name(); sHasSuffix(jsonfile, ".json") {
			fPln("--->", jsonfile)

			bytes, _ := os.ReadFile(dir + jsonfile)
			jsonstr := string(bytes)

			if !IsValid(jsonstr) {
				os.WriteFile(fSf("debug_%s.json", jsonfile), []byte(jsonstr), 0666)
				panic("error on MkJSON")
			}

			//if jsonfile == "CensusCollection_0.xml" {
			// os.WriteFile(fSf("record_%s.json", jsonfile), []byte(jsonstr), 0666)
			//}
		}
		return nil
	})
}
