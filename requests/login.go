package requests

type LoginRequest struct {
	Version     string `json:"version"`
	Device      string `json:"device"`
	Seq         int64  `json:"seq,string"`
	Platform    int64  `json:"platform,string"`
	Language    int64  `json:"language,string"`
	SalesLocate int64  `json:"salesLocate,string"`
	StoreID     int64  `json:"storeId,string"`
	PlatformSNS int64  `json:"platform_sns,string"`
	LineAuth    `json:"lineAuth"`
}
