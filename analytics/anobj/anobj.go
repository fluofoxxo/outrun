package anobj

import "time"

type Aggret struct { // "aggregator"
	PlayerID              string  `json:"playerID"`
	CreationDate          int64   `json:"creationDate"`
	Data                  []int64 `json:"data"`
	HistoricalStoryScores []int64 `json:"historicalStoryScores"`
	HistoricalTimedScores []int64 `json:"historicalTimedScores"`
	LoginTimes            []int64 `json:"loginTimes"`
}

func NewAggret(pid string) Aggret {
	return Aggret{
		pid,
		time.Now().UTC().Unix(),
		make([]int64, 75), // must be longer than length of factors.AnalyticType*
		make([]int64, 20),
		make([]int64, 20),
		[]int64{},
	}
}
