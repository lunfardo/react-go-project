package scheduler

import (
	"time"
)

//Execute x every timeSec seconds
func AddTask(x func(), timeSec time.Duration) {
	signal := make(chan bool)
	go func() {
		for true {
			time.Sleep(time.Second * timeSec)
			signal <- true
		}
	}()
	for {
		select {
		case <-signal:
			go x()
		}
	}
}
