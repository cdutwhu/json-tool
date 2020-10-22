package jsontool

// BreakMulEleBlk : 'jstr' is LIKE {"1st-element": {...}, "2nd-element": {...}, "3rd-element": [...]}
// return one 'value' is like '{...}', OR like `[{...},{...},...]`
func BreakMulEleBlk(jstr string) (names, values []string) {
	jstr = sTrim(jstr, " \t\n")
	failOnErrWhen(jstr[0] != '{', "%v", fEf("error (format) json"))
	failOnErrWhen(jstr[len(jstr)-1] != '}', "%v", fEf("error (format) json"))
	// jstrhead := jstr
	rName := rxMustCompile(`"[-\w]+":`)   // "name":
	rsVal := rxMustCompile(`^[-\d\.tfn]`) // non-string, simple value start

NEXT:
	if loc := rName.FindStringIndex(jstr); loc != nil { // find attr "name":
		s, e := loc[0], loc[1]
		root := jstr[s+1 : e-2]
		// fPln(root)
		names = append(names, root)
		jstr = sTrimLeft(jstr[e:], " ") // start @ "{" or "[" or simple...

		// Simple Non-String values
		if loc := rsVal.FindStringIndex(jstr); loc != nil {
			// fPln("non-string simple ele")
			for i := 1; i < len(jstr); i++ { // skip the 1st char
				c := jstr[i]
				if c == ',' || c == '\n' {
					values = append(values, jstr[:i])
					jstr = jstr[i+1:]
					goto NEXT
				}
			}
		}

		// Complex, Array or String value
		for i, mark := range []string{"{", "[", "\""} {
			if sHasPrefix(jstr, mark) {
				var m1, m2 byte
				switch i {
				case 0:
					m1, m2 = '{', '}'
				case 1:
					m1, m2 = '[', ']'
				default:
					m1, m2 = '"', '"'
				}
				L := 0
				for i := 0; i < len(jstr); i++ {
					c := jstr[i]
					if m1 != m2 { // Complex, Array
						if c == m1 { // { or [
							L++
						}
						if c == m2 { // } or ]
							L--
							if L == 0 {
								values = append(values, jstr[:i+1])
								jstr = jstr[i+1:]
								goto NEXT
							}
						}
					} else { // String
						if c == m1 { // "***"
							L++
							if L == 2 {
								// values = append(values, jstr[1:i]) // remove '"' at start&end (string & other types mixed)
								values = append(values, jstr[:i+1]) // remove '"' at start&end
								jstr = jstr[i+1:]
								goto NEXT
							}
						}
					}
				}
			}
		}
	}
	return
}

// BreakArr : 'jstr' is like [{...},{...}]
// i.e. [{...},{...}] => {...} AND {...}
// NO ele name could get here
func BreakArr(jstr string) (values []string, ok bool) {
	jstr = sTrim(jstr, " ")
	if jstr[0] != '[' || jstr[len(jstr)-1] != ']' {
		return values, false
	}
	L, S := 0, -1
	for i := 0; i < len(jstr); i++ {
		c := jstr[i]
		if c == '{' {
			L++
			if L == 1 {
				S = i
			}
		}
		if c == '}' {
			L--
			if L == 0 {
				values = append(values, jstr[S:i+1])
			}
		}
	}
	return values, true
}

// BreakMulEleBlkV2 : 'jstr' LIKE { "1st-element": {...}, "2nd-element": {...}, "3rd-element": [...] }
// in return 'values', array types are broken into duplicated names & its single value block
// one 'value' is like '{...}', 'names' may have duplicated names
func BreakMulEleBlkV2(jstr string) (names, values []string) {
	mIndEles := make(map[int][]string)
	Names, Values := BreakMulEleBlk(jstr)
	for i, Val := range Values {
		if eles, ok := BreakArr(Val); ok {
			mIndEles[i] = eles
		}
	}
	for i, Val := range Values {
		if eles, ok := mIndEles[i]; ok {
			for _, ele := range eles {
				names = append(names, Names[i])
				values = append(values, ele)
			}
		} else {
			names = append(names, Names[i])
			values = append(values, Val)
		}
	}
	return
}
