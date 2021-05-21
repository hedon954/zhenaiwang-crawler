package scheduler

import "zhenaiwang-crawler/engine"

/**
	版本一：简单调度器
		描述：所有 Worker 共用一个 Channel
 */
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(requests chan engine.Request) {
	//Simple版本不做任何事情
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

//提交请求
func (s *SimpleScheduler) Submit(request engine.Request) {
	//开一个 goRouting 去 submit，避免循环等待造成死锁
	go func() {
		s.workerChan <- request
	}()
}


