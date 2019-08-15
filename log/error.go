package log

import (
	"log"
	"net/http"
)

// TODO: Abandon this file. All functions here are applied within the Helper struct.
const (
	PREFIX_ERR             = "ERR"
	PREFIX_OUT             = "OUT"
	PREFIX_WARN            = "WARN"
	PREFIX_UNCATCHABLE_ERR = "UNCATCHABLE ERR"

	LOGOUT_BASE = "{%s} (%s) %s\n"
	LOGERR_BASE = "{%s} (%s) %s: %s\n"

	INTERNAL_SRV_ERR = "Internal server error"
	BAD_REQUEST      = "Bad request"
)

type Reporter struct {
	CallerName string
	RespW      http.ResponseWriter
}

func MakeReporter(callerName string, r http.ResponseWriter) *Reporter {
	return &Reporter{
		callerName,
		r,
	}
}

func (r *Reporter) Out(msg string) {
	log.Printf(LOGOUT_BASE, PREFIX_OUT, r.CallerName, msg)
}
func (r *Reporter) Warn(msg string) {
	log.Printf(LOGOUT_BASE, PREFIX_WARN, r.CallerName, msg)
}
func (r *Reporter) Uncatchable(msg string) {
	log.Printf(LOGOUT_BASE, PREFIX_OUT, r.CallerName, msg)
}
func (r *Reporter) InternalErr(msg string, err error) {
	log.Printf(LOGERR_BASE, PREFIX_ERR, r.CallerName, msg, err.Error())
	r.RespW.WriteHeader(http.StatusBadRequest)
	r.RespW.Write([]byte(BAD_REQUEST))
}
func (r *Reporter) Err(msg string, err error) {
	log.Printf(LOGERR_BASE, PREFIX_ERR, r.CallerName, msg, err.Error())
	r.RespW.WriteHeader(http.StatusBadRequest)
	r.RespW.Write([]byte(BAD_REQUEST))
}
func (r *Reporter) InternalFatal(msg string, err error) {
	log.Fatalf(LOGERR_BASE, PREFIX_ERR, r.CallerName, msg, err.Error())
	r.RespW.WriteHeader(http.StatusBadRequest)
	r.RespW.Write([]byte(BAD_REQUEST))
}
func (r *Reporter) Fatal(msg string, err error) {
	log.Fatalf(LOGERR_BASE, PREFIX_ERR, r.CallerName, msg, err.Error())
	r.RespW.WriteHeader(http.StatusBadRequest)
	r.RespW.Write([]byte(BAD_REQUEST))
}
