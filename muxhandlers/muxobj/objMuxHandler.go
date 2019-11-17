package muxobj

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fluofoxxo/outrun/config"
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
			if config.CFile.LogAllRequests {
				nano := time.Now().UnixNano()
				nanoStr := strconv.Itoa(int(nano))
				filename := help.Request.RequestURI + "--" + nanoStr
				filename = strings.ReplaceAll(filename, ".", "-")
				filename = strings.ReplaceAll(filename, "/", "-") + ".txt"
				filepath := "logging/all_requests/" + filename
				help.Out("DEBUG: Saving request to " + filepath)
				err := ioutil.WriteFile(filepath, help.GetGameRequest(), 0644)
				if err != nil {
					help.Out("DEBUG ERROR: Unable to write file '" + filepath + "'")
				}
			}
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
