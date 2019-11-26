package taskrunner

import "time"

//timer
//setup
//start{trigger->task->runner}
//
//timer,

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval),
		runner: r,
	}
}

// 也就是每隔3秒就批量删除一次
func (w *Worker) startWorker() {
	for {
		select {
		case <- w.ticker.C:		// 定时一到，就启动runner
			go w.runner.StartAll()
		}
	}
}

func Start() {
	// start video cleaning
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3*time.Second, r)
	go w.startWorker()

}