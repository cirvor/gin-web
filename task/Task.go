package task

import (
	"log"
	"time"

	"github.com/jasonlvhit/gocron"
)

func Task() {
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(run)

	<-s.Start()
}

func run() {
	log.Println("I am runnning task.", time.Now())
}
