package analytics

import (
	"encoding/json"
	"log"
	"time"

	"github.com/fluofoxxo/outrun/analytics/anobj"
	"github.com/fluofoxxo/outrun/analytics/factors"
	"github.com/fluofoxxo/outrun/config"
	"github.com/fluofoxxo/outrun/consts"
	"github.com/fluofoxxo/outrun/db/dbaccess"
)

func Store(pid string, atype int, value ...int64) (bool, error) {
	if !config.CFile.EnableAnalytics {
		return true, nil
	}
	if config.CFile.DebugPrints {
		log.Println("analytics.Store", pid, atype, value)
	}
	found := true
	data, err := dbaccess.Get(consts.DBBucketAnalytics, pid)
	var aggret anobj.Aggret
	if err != nil {
		// player isn't found, so make a new Aggret
		log.Println("analytics.Store warning: error (", err.Error(), "), using default")
		aggret = anobj.NewAggret(pid)
		found = false
	} else {
		if config.CFile.DebugPrints {
			log.Println("analytics.Store fetched data: ", string(data))
		}
		if len(data) != 0 { // there's data here
			err = json.Unmarshal(data, &aggret)
			if err != nil {
				return found, err
			}
		} else {
			log.Println("analytics.Store warning: data found, but empty, using default")
			aggret = anobj.NewAggret(pid)
			found = false
		}
	}

	if atype == factors.AnalyticTypeAverageStoryScore {
		// add score to historical scores for averaging
		aggret.HistoricalStoryScores = append(aggret.HistoricalStoryScores, value[0])[1:]
		// average scores
		total := int64(0)
		count := 0
		for _, score := range aggret.HistoricalStoryScores {
			if score > 0 {
				count++
				total += score
			}
		}
		avg := total / int64(count)
		aggret.Data[atype] = avg
	} else if atype == factors.AnalyticTypeAverageTimedScore {
		// add score to historical scores for averaging
		aggret.HistoricalTimedScores = append(aggret.HistoricalTimedScores, value[0])[1:]
		// average scores
		total := int64(0)
		count := 0
		for _, score := range aggret.HistoricalTimedScores {
			if score > 0 {
				count++
				total += score
			}
		}
		avg := total / int64(count)
		aggret.Data[atype] = avg
	} else if atype == factors.AnalyticTypeLogins {
		aggret.Data[atype]++
		aggret.LoginTimes = append(aggret.LoginTimes, time.Now().UTC().Unix())
	} else {
		if len(value) == 0 {
			aggret.Data[atype]++
		} else {
			aggret.Data[atype] += value[0]
		}
	}

	jdata, err := json.Marshal(aggret)
	if err != nil {
		return found, err
	}
	err = dbaccess.Set(consts.DBBucketAnalytics, pid, jdata)
	if err != nil {
		return found, err
	}

	return found, nil
}

func Get(pid string, atype int) (int64, bool, error) {
	found := true
	data, err := dbaccess.Get(consts.DBBucketAnalytics, pid)
	var aggret anobj.Aggret
	if err != nil {
		// player isn't found, so make a new Aggret
		aggret = anobj.NewAggret(pid)
		found = false
	}
	err = json.Unmarshal(data, &aggret)
	if err != nil {
		return 0, found, err
	}
	return aggret.Data[atype], found, nil
}
