package jsontool

import (
	"os"
	"testing"
	"time"

	"github.com/cdutwhu/gotil/misc"
)

func TestJSONBreakArrCont(t *testing.T) {
	defer misc.TrackTime(time.Now())

	bytes, err := os.ReadFile("./data/Activities.json")
	failOnErr("%v", err)
	jsonstr := string(bytes)

	values, ok := BreakArr(jsonstr)
	fPln(ok)
	for _, v := range values {
		fPln(v)
	}
}

func TestJSONBreakBlkContV2(t *testing.T) {
	defer misc.TrackTime(time.Now())
	if bytes, err := os.ReadFile("./why.json"); err == nil {
		// jsonstr := JSONBlkFmt(string(bytes), "  ")
		jsonstr := string(bytes)
		_, cont := SglEleBlkCont(jsonstr)
		names, values := BreakMulEleBlkV2(cont)
		for i, name := range names {
			fPln(MkSglEleBlk(name, values[i], false))
			fPln(" ------------------------------------------ ")
		}
	}
}

func Test_detectLCB(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				line: `"ActivityTime": {`,
			},
			want: 1,
		},
		{
			name: "OK",
			args: args{
				line: `}}`,
			},
			want: -2,
		},
		{
			name: "OK",
			args: args{
				line: `"-RefId": "C27E1FCF-C163{{-485F-BEF0-F36F18A0493A`,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := detectLCB(tt.args.line); got != tt.want {
				t.Errorf("detectLCB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanArray2Objects(t *testing.T) {

	file, err := os.OpenFile("./data/data.json", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fPln(err)
	}

	if chRst, ja := ScanArrayObject(file, true); !ja {
		fPln("NOT JSON array")
	} else {

		I := 1
		for rst := range chRst {
			fPln(I)
			fPln(rst.Obj)
			fPln(rst.Err)
			I++
		}

		// for {
		// 	if rst, more := <-chRst; more {
		// 		fPln(I)
		// 		fPln(rst.Obj)
		// 		fPln(rst.Err)
		// 		I++
		// 	} else {
		// 		break
		// 	}
		// }
	}
}
