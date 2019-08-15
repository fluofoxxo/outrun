package bgtasks

import (
	"time"

	"github.com/fluofoxxo/outrun/db"
)

func MainTask() {
	for true {
		time.Sleep(10 * time.Minute)
		db.PurgeAllExpiredSessionIDs()
	}
}
