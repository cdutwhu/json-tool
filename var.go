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
