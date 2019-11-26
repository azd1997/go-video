package taskrunner

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {

	// 用于提交任务
	d := func(dc dataChan) error {
		for i:=0; i<30; i++ {
			dc <- i
			log.Printf("Dispatcher sent: %d", i)
		}
		return nil
	}

	// 用于执行任务
	e := func(dc dataChan) error {
	ForLoop:
		for {
			select {
			case d := <- dc:
				log.Printf("Executor received: %v", d)
			default:
				break ForLoop
			}
		}
		//return nil
		return errors.New("executor")
	}

	runner := NewRunner(30, false, d, e)
	go runner.StartAll()
	time.Sleep(3 * time.Second)
}
