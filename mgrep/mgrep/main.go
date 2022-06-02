package main

import (
	"fmt"
	"os"
	"mgrep/worklist"
	"mgrep/worker"
	"path/filepath"
	"sync"

	"github.com/alexflint/go-arg"
)

func discoverDirs(wl *worklist.Worklist, path string) {
	entires, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Readdir error:", err)
		return
	}

	for _, entry := range entires {
		if entry.IsDir() {
			nextPath := filepath.Join(path, entry.Name())
			discoverDirs(wl, nextPath)
		} else {
			wl.Add(worklist.NewJob(filepath.Join(path, entry.Name())))
		}
	}
}

var args struct {
	SearchTerm string	`args:"positional,required"`
	SearchDir string 	`arg:"positional"`
}

func main() {
	arg.MustParse(&args)

	var workersWg sync.WaitGroup
	// capacity 100
	wl := worklist.New(100)

	resultsChannel := make(chan worker.Result, 100)

	numWorkers := 10

	workersWg.Add(1)
	// goroutine -- 1 to discover files
	go func() {
		defer workersWg.Done()
		discoverDirs(&wl, args.SearchDir)
		wl.Finalize(numWorkers)
	}()

	for i := 0 ; i < numWorkers ; i++ {
		workersWg.Add(1)
		go func() {
			defer workersWg.Done()
			for {
				workEntry := wl.Next()
				if workEntry.Path != "" {
					workerResult := worker.FindInFile(workEntry.Path, args.SearchTerm)
					if workerResult != nil {
						for _, r := range workerResult.Inner {
							resultsChannel <- r
						}
					}
				} else {
					return
				}
 			}
		}()
	}

	// Wait for the workers to finish
	blockWorkersWg := make(chan struct{})
	go func() {
		workersWg.Wait()
		close(blockWorkersWg)
	}()

	var displayWg sync.WaitGroup
	displayWg.Add(1)
	go func() {
		for {
			select {
			case r := <- resultsChannel:
				fmt.Printf("%v[%v]:%v\n", r.Path, r.LineNum, r.Line)
			case <- blockWorkersWg:
				if len(resultsChannel) == 0 {
					displayWg.Done()
					return
				}
			}
		}
	}()
	displayWg.Wait()
}