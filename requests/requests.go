package requests

type BasicRequest struct {
	SessionID string `json:"sessionId"`
	Version   string `json:"version"`
	Seq       int64  `json:"seq,string"`
}

type LineAuth struct {
	UserID            string `json:"userId"`
	Password          string `json:"password"`
	MigrationPassword string `json:"migrationPassword"`
}

type NotPlayingBase struct {
	Version     string `json:"version"`
	Seq         int64  `json:"seq,string"`
	LineAuth    `json:"lineAuth"`
	Platform    int64  `json:"platform,string"` // what does this do? does 2 mean android?
	Device      string `json:"device"`
	Language    int64  `json:"language,string"`    // 1 means english
	SalesLocate int64  `json:"salesLocate,string"` // ?
	StoreID     int64  `json:"storeId,string"`
	PlatformSNS int64  `json:"platform_sns,string"` // seems to be for statistics
}

type RedStarExchangeListRequest struct {
	BasicRequest
	ItemType int64 `json:"itemType,string"`
}
