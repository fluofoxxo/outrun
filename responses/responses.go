package responses

import (
	"encoding/json"
	"time"

	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/objects"
	"github.com/jinzhu/now"
)

// TODO: move the loose structs (like Incentive and WheelOptions) to another package

func ToJSON(i interface{}) (string, error) {
	b, err := json.Marshal(i)
	s := string(b)
	return s, err
}

func getBasicIncentives() []objects.Incentive {
	inc := []objects.Incentive{}
	for i := range make([]byte, 7) { // makes 7 duplicates of the same item for the daily challenge
		ix := int64(i + 1)
		incentive := objects.MakeIncentive("900000", 5, ix)
		inc = append(inc, incentive)
	}
	return inc
}

type BaseInfo struct {
	CloseTime    int64        `json:"closeTime"`
	ErrorMessage consts.EMess `json:"errorMessage,string"`
	Seq          int64        `json:"seq,string"`
	ServerTime   int64        `json:"server_time"`
	StatusCode   int64        `json:"statusCode"`
}

func (b BaseInfo) SetErrorMessage(s string) {
	b.ErrorMessage = consts.EMess(s)
}

func NewBaseInfo(errorMessage string, seq, statusCode int64) BaseInfo {
	closeTime := now.EndOfDay().Unix()
	serverTime := time.Now().Unix()
	return BaseInfo{
		closeTime,
		consts.EMess(errorMessage),
		seq,
		serverTime,
		statusCode,
	}
}

type BaseResponse struct {
	BaseInfo
	AssetsVersion     string `json:"assets_version"`
	ClientDataVersion string `json:"client_data_version"`
	DataVersion       string `json:"data_version"`
	InfoVersion       string `json:"info_version"`
	Version           string `json:"version"`
}

func NewBaseResponse(base BaseInfo) BaseResponse {
	br := BaseResponse{
		base,
		"049",
		"2.0.3",
		"15",
		"017",
		"2.0.3",
	}
	return br
}
