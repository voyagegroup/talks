package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

type info struct {
	size int64
	name string
}

func trace(dir string, ch chan<- info) error {
	f, err := os.Open(dir)
	if err != nil {
		return err
	}
	files, err := f.Readdir(0)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			trace(file.Name(), ch)
			continue
		}
		ch <- info{
			size: file.Size(),
			name: file.Name(),
		}
	}
	return nil
}

func main() {
	flag.Parse()
	info := make(chan info)
	dirs := flag.Args()
	if len(dirs) == 0 {
		dirs = []string{"."}
	}
	var wg sync.WaitGroup
	for _, d := range dirs {
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			if err := trace(d, info); err != nil {
				log.Printf("err: %s", err)
			}
		}(d)
	}

	go func() {
		wg.Wait()
		close(info)
	}()

	var total, files int64
LOOP:
	for {
		s, ok := <-info
		if !ok {
			break LOOP
		}
		// fmt.Printf("%s: %d\n", s.name, s.size)
		total += s.size
		files++
	}

	fmt.Printf("%d files, %d bytes", files, total)
}
