package jsontool

// Root :
// LIKE { "only-one-element": { ... } }
func Root(jstr string) string {
	root, _ := SglEleBlkCont(jstr)
	return root
}

// MkSglEleBlk :
// LIKE { "only-one-element": { ... } }
func MkSglEleBlk(name string, value interface{}, fmt bool) string {
	// string type value to be added "quotes"
	switch value.(type) {
	case string:
		sval := value.(string)
		if !(len(sval) >= 2 && (sval[0] == '{' || sval[0] == '[')) {
			value = fSf(`"%s"`, sTrim(sval, `"`))
		}
	}

	jstr := fSf(`{ "%s": %v }`, name, value)
	// failOnErrWhen(!IsValid(jstr), "%v", fEf("Err In Making JSON Block")) // test mode open
	if fmt {
		return Fmt(jstr, "  ")
	}
	return jstr
}

// SglEleBlkCont :
// LIKE { "only-one-element": { ... } }
func SglEleBlkCont(jstr string) (string, string) {
	qtIdx1, qtIdx2 := -1, -1
	for i := 0; i < len(jstr); i++ {
		if qtIdx1 == -1 && jstr[i] == '"' {
			qtIdx1 = i
			continue
		}
		if qtIdx1 != -1 && jstr[i] == '"' {
			qtIdx2 = i
			break
		}
	}
	failOnErrWhen(jstr[qtIdx2+1] != ':', "%v", fEf("error (format) json"))
	failOnErrWhen(jstr[qtIdx2+2] != ' ', "%v", fEf("error (format) json"))
	ebIdx := sLastIndex(jstr, "}")
	return jstr[qtIdx1+1 : qtIdx2], sTrimRight(jstr[qtIdx2+3:ebIdx], " \t\n\r")
}
