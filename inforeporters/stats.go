package inforeporters

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type statsInfo struct {
	AllocatedMem   uint64    `json:"allocatedMemory"`
	GoroutineCount int       `json:"goroutineCount"`
	CPUUsages      []float64 `json:"cpuUsages"`
}

func Stats(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	allocated := m.TotalAlloc / 1024 / 1024 // megabytes
	goroutines := runtime.NumGoroutine()
	cpus, err := cpu.Percent(10*time.Millisecond, false)
	if err != nil {
		// ignore for now
		cpus = []float64{0.0}
	}
	stats := statsInfo{
		allocated,
		goroutines,
		cpus,
	}
	jout, err := json.Marshal(stats)
	if err != nil {
		log.Println("error collecting stats: " + err.Error())
	}
	w.Write(jout)
}
