package scheduler

import "zhenaiwang-crawler/engine"


/**
	版本二：并发调度器
		描述：每一个 Worker 独享一个自己的 Channel
 */

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request  //每个 worker 对外的接口是一个 chan engine.Request
}

func (q *QueuedScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

//提交任务
func (q *QueuedScheduler) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request){
	q.workerChan <- w
}

//创建 go routing
func (q *QueuedScheduler) Run()  {
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ  []engine.Request
		var workerQ  []chan engine.Request
		for  {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			//将 request 扔进 request 队列
			case r := <- q.requestChan:
				//TODO: send r to a ?worker
				requestQ = append(requestQ, r)
			//将 worker 扔进 worker 队列
			case w := <- q.workerChan:
				//TODO: send next request to w
				workerQ = append(workerQ, w)
			//request 和 worker 都存在的时候就将 request 交给 worker 去工作
			case activeWorker <- activeRequest:
				//真正执行activeWorker <- activeRequest的时候才将他们从队列中拿掉
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

