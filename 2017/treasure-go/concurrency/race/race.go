package race

import "fmt"

func f() {
	ch := make(chan bool)
	m := make(map[int]bool)
	go func() {
		m[1] = false
		ch <- true
	}()
	_ = m[2]
	<-ch
	for k, v := range m {
		fmt.Println(k, v)
	}
}
