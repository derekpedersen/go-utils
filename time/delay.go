package time

import (
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

func Delay() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) + 1 // n will be between 1 and 11
	log.Debugf("Sleeping %d seconds...\n", n)
	time.Sleep(time.Duration(n) * time.Second)
}
