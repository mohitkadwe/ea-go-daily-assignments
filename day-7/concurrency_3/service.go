package executor

import (
	"errors"
	"sync"
)

type Service struct {
	jobChan  chan func()
	isClosed bool
}

func (s *Service) Start(workersCnt int, setup func()) {
	jobChan := make(chan func())
	s.jobChan = jobChan
	wg := sync.WaitGroup{}
	wg.Add(workersCnt)

	for i := 0; i < workersCnt; i++ {
		go worker(setup, &wg, jobChan)
	}

	wg.Wait()
}

func (s Service) Run(job func()) error {
	if s.isClosed {
		return errors.New("closed")
	}
	s.jobChan <- job
	return nil
}

func (s *Service) Close() {
	close(s.jobChan)
	s.isClosed = true
}

func worker(setup func(), wg *sync.WaitGroup, jobChan chan func()) {
	setup()
	wg.Done()
	for {
		job, ok := <-jobChan
		if !ok {
			break
		}
		job()
	}
}

func RunBatch(jobs []func(), workersCnt int) {
	jobChan := make(chan func(), workersCnt)
	wg := sync.WaitGroup{}
	wg.Add(workersCnt)

	for i := 0; i < workersCnt; i++ {
		go workerBatch(&wg, jobChan)
	}

	for _, job := range jobs {
		jobChan <- job
	}

	close(jobChan)
	wg.Wait()
}

func workerBatch(wg *sync.WaitGroup, jobChan chan func()) {
	wg.Done()
	for {
		job, ok := <-jobChan
		if !ok {
			break
		}
		job()
	}
}
