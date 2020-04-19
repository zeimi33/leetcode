package main

import (
	"sync"
)

type mmaapp struct {
	m   map[string]struct{}
	mx  sync.Mutex
	w   sync.WaitGroup
	c   chan int
	run bool
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++s
// func main() {
// 	selfmap := mmaapp{}
// 	selfmap.c = make(chan int)
// 	selfmap.run = true
// 	go selfmap.a()
// 	m := make(map[int]struct{})
// 	time1 := time.Now()
// 	for i := 0; i < 100000000; i++ {
// 		m[i] = struct{}{}
// 	}
// 	time2 := time.Now()
// 	fmt.Println(time1, time2)
// 	time1 = time.Now()
// 	for i := 0; i < 100000000; i++ {
// 		if _, ok := m[i]; ok {

// 		}
// 	}
// 	time2 = time.Now()

// 	fmt.Println(time1, time2)
// 	time1 = time.Now()
// 	for i := 0; i < 100000000; i++ {
// 		selfmap.c <- i
// 	}
// 	time2 = time.Now()
// 	fmt.Println(time1, time2)
// }
// func (m *mmaapp) a() {
// 	for m.run {
// 		<-m.c
// 	}
// }

// func main() {
// 	mmaapp := mmaapp{}
// 	mmaapp.m = make(map[string]struct{})
// 	for i := 0; i < 1000; i++ {
// 		mmaapp.w.Add(1)
// 		go mmaapp.rwmap()
// 	}
// 	mmaapp.w.Wait()
// 	fmt.Println(mmaapp.m)
// }

// func (mmaapp *mmaapp) rwmap() {
// 	for i := 0; i < 10000; i++ {
// 		if _, ok := mmaapp.m[fmt.Sprintln(i)]; !ok {
// 			mmaapp.mx.Lock()
// 			mmaapp.m[fmt.Sprintln(i)] = struct{}{}
// 			mmaapp.mx.Unlock()
// 		}
// 	}
// 	mmaapp.w.Done()
// }

// func main() {

// 	a := sync.Map{}
// 	w := sync.WaitGroup{}
// 	timestart := time.Now()
// 	for i := 0; i < 100; i++ {
// 		w.Add(1)
// 		go writeOrRead(&a, &w)
// 	}
// 	w.Wait()
// 	timeend := time.Now()
// 	fmt.Println(timestart, timeend)

// }

// func writeOrRead(a *sync.Map, w *sync.WaitGroup) {
// 	for i := 0; i < 1000000; i++ {
// 		a.LoadOrStore(i, 1)
// 	}
// 	w.Done()
// }
