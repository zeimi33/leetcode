package main

import (
	"fmt"
)

func main(){
	m := make(map[int]string)
	m[2]="2"
	m[1]="1"
	m[3]="3"
	for k,v :=range m{
		fmt.Println(k,v)
	}
}



//-------------------------------------------------------------------
//var a chan int
//func main (){
//	go func() {
//		log.Println(http.ListenAndServe(":8080",nil))
//	}()
//	c := sync.WaitGroup{}
//	c.Add(1)
//	go trun(&c)
//	c.Wait()
//}
//
//func trun(c *sync.WaitGroup){
//	for  {
//		select {
//		case <-a:fmt.Println(1)
//		default:
//		}
//	}
//	c.Done()
//}