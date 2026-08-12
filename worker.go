package main

import (
	"sync"
	"time"
)

type ScanJob struct {
	Index  int
	Target string
}

type ScanResult struct {
	Index  int
	Result *Result
	Err    error
}

func runWorkers(
	jobs []ScanJob,
	workers int,
	timeout time.Duration,
	profile string,
) []*ScanResult {

	jobChan := make(chan ScanJob)
	resultChan := make(chan *ScanResult)

	var wg sync.WaitGroup

	if workers < 1 {
		workers = 1
	}

	for i := 0; i < workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for job := range jobChan {
				r, err := scan(
					job.Target,
					timeout,
					profile,
				)

				resultChan <- &ScanResult{
					Index:  job.Index,
					Result: r,
					Err:    err,
				}
			}
		}()
	}

	go func() {
		for _, job := range jobs {
			jobChan <- job
		}

		close(jobChan)
		wg.Wait()
		close(resultChan)
	}()

	var results []*ScanResult

	for r := range resultChan {
		results = append(results, r)
	}

	return results
}
