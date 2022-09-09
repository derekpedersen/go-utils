package memory

import (
	"runtime"

	"github.com/sirupsen/logrus"
)

func GetMemory() runtime.MemStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m
}

func GetMemoryLogFields() logrus.Fields {
	m := GetMemory()
	return logrus.Fields(map[string]interface{}{
		"Mem Alloc MiB":   bToMb(m.Alloc),
		"Total Alloc MiB": bToMb(m.TotalAlloc),
		"System MiB":      bToMb(m.Sys),
		"Go Routines":     runtime.NumGoroutine(),
	})
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
