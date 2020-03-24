package main

import "sync"

func main(){
	a := sync.Mutex{}
	a.Lock()
}
