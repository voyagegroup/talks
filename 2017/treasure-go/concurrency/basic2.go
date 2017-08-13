package main

type Task struct {
}

var numWorkers = 3

func process(t Task) {
	// ...
}

func worker(ch chan Task) {
	for {
		task := <-ch
		process(task)
	}
}

func main() {
	ch := make(chan Task, 3)
	for i := 0; i < numWorkers; i++ {
		go worker(ch)
	}

	tasks := getTasks()
	for _, task := range tasks {
		ch <- task
	}

	// brabra..
}
