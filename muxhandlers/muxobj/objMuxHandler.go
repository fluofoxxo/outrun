package muxobj

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/fluofoxxo/outrun/helper"
)

func Handle(f func(*helper.Helper), logExecutionTime bool) func(w http.ResponseWriter, r *http.Request) {
	funcnameSplit := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), "/")
	funcname := funcnameSplit[len(funcnameSplit)-1]
	funcnameSplit = strings.Split(funcname, ".") // just get function name
	funcname = funcnameSplit[len(funcnameSplit)-1]
	if logExecutionTime {
		return func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			help := helper.MakeHelper(funcname, w, r)
			f(help)
			endTime := time.Now()
			timeDiff := endTime.Sub(startTime)
			help.Out("Done executing in " + timeDiff.String())
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		help := helper.MakeHelper(funcname, w, r)
		f(help)
	}
}
