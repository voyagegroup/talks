package main

import (
	"log"
	"net/http"
	"sync"
)

type Task struct {
	rawurl string
}

var numWorkers = 3

func crawl(t Task) error {
	resp, err := http.Get(t.rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	log.Printf("crawled: %s", t.rawurl)
	// save to file..
	return nil
}

func worker(ch chan Task) {
	for {
		task := <-ch
		crawl(task)
	}
}

func getTasks() []Task {
	return []Task{
		{"https://voyagegroup.com"},
		{"http://ecnavi.jp"},
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan Task, 3)
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ch)
		}()
	}

	tasks := getTasks()
	for _, task := range tasks {
		ch <- task
	}

	wg.Wait()
}
