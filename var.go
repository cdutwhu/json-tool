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
	sHasPrefix    = strings.HasPrefix
	sTrim         = strings.Trim
	sTrimLeft     = strings.TrimLeft
	sTrimRight    = strings.TrimRight
	sReplaceAll   = strings.ReplaceAll
	sLastIndex    = strings.LastIndex
	rxMustCompile = regexp.MustCompile
	failOnErr     = fn.FailOnErr
	failOnErrWhen = fn.FailOnErrWhen
)
