package muxhandlers

import (
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/fluofoxxo/outrun/helper"
)

func Handle(f func(*helper.Helper), logExecutionTime bool) func(w http.ResponseWriter, r *http.Request) {
	funcname := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	if logExecutionTime {
		return func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			help := helper.MakeHelper(funcname, w, r)
			f(help, r)
			endTime := time.Now()
			timeDiff := endTime.Sub(startTime)
			help.Out("Done executing in " + timeDiff.String())
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		help := helper.MakeHelper(funcname, w)
		f(help, r)
	}
}
