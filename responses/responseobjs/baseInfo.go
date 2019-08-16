package responseobjs

import (
	"time"

	"github.com/jinzhu/now"
)

type BaseInfo struct {
	ErrorMessage ErrorMessage `json:"errorMessage,string"`
	CloseTime    int64        `json:"closeTime"` // end of the day
	Seq          int64        `json:"seq,string"`
	ServerTime   int64        `json:"server_time"`
	StatusCode   int64        `json:"statusCode"`
}

func (b BaseInfo) SetErrorMessage(message string) {
	b.ErrorMessage = ErrorMessage(message)
}

func NewBaseInfo(em string, statusCode int64) BaseInfo {
	// seq is a default 0 for now, since it does not impact gameplay thus far
	closeTime := now.EndOfDay().Unix()
	serverTime := time.Now().Unix()
	seq := int64(0)
	return BaseInfo{
		ErrorMessage(em),
		closeTime,
		seq,
		serverTime,
		statusCode,
	}
}
