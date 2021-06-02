package jsontool

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cdutwhu/debog/fn"
)

var (
	fPln          = fmt.Println
	fSf           = fmt.Sprintf
	fEf           = fmt.Errorf
	sJoin         = strings.Join
	sTrim         = strings.Trim
	sTrimLeft     = strings.TrimLeft
	sTrimRight    = strings.TrimRight
	sReplaceAll   = strings.ReplaceAll
	sIndex        = strings.Index
	sLastIndex    = strings.LastIndex
	sHasPrefix    = strings.HasPrefix
	sHasSuffix    = strings.HasSuffix
	rxMustCompile = regexp.MustCompile
	failOnErr     = fn.FailOnErr
	failOnErrWhen = fn.FailOnErrWhen
)

var (
	DEBUG = 0
)

// dropCR drops a terminal \r from the data.
var dropCR = func(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
