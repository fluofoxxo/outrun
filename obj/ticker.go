package obj

import "time"

type Ticker struct {
	ID    int64  `json:"id"`
	Start int64  `json:"start"`
	End   int64  `json:"end"`
	Param string `json:"param"`
}

func NewTicker(id, end int64, param string) Ticker {
	start := time.Now().Unix()
	return Ticker{
		id,
		start,
		end,
		param,
	}
}
