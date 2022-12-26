package executor

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	s := Service{}
	expCnt := 3
	var actualCnt int32 = 0

	setup := func() {
		atomic.AddInt32(&actualCnt, 1)
	}

	s.Start(3, setup)

	assert.Equal(t, expCnt, int(actualCnt))
}

func TestRun(t *testing.T) {
	s := Service{}
	s.Start(3, func() {})
	expCnt := 4
	var actualCnt int32 = 0

	wg := sync.WaitGroup{}
	wg.Add(expCnt)
	job := func() {
		wg.Done()
		atomic.AddInt32(&actualCnt, 1)
	}
	for i := 0; i < expCnt; i++ {
		s.Run(job)
	}
	wg.Wait()
	assert.Equal(t, expCnt, int(actualCnt))
}

func TestClose(t *testing.T) {
	s := Service{}
	s.Start(3, func() {})

	s.Close()

	executed := false
	job := func() {
		executed = true
	}
	err := s.Run(job)

	assert.Error(t, err)
	assert.False(t, executed)
}

func TestRunBatch(t *testing.T) {

	s := Service{}
	s.Start(3, func() {})
	expCnt := 4
	var actualCnt int32 = 0

	wg := sync.WaitGroup{}
	wg.Add(expCnt)
	job := func() {
		wg.Done()
		atomic.AddInt32(&actualCnt, 1)
	}
	jobs := make([]func(), 0, expCnt)
	for i := 0; i < expCnt; i++ {
		jobs = append(jobs, job)
	}
	RunBatch(jobs, 3)
	wg.Wait()
	assert.Equal(t, expCnt, int(actualCnt))

}
