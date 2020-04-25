package main

import "fmt"

func main(){
fmt.Println(waysToChange1(25))
}



func waysToChange1(n int) (ans int) {
	const v int=1000000007
	for i:=n/5+1;i>0;i=i-5 {
		if t:=i%2;t==1 {
			t_v:=(i+1)/2
			ans=(ans+(t_v*t_v)%v)%v
		} else {
			t_v:=i/2
			ans=(ans+(t_v*(t_v+1))%v)%v
		}
	}
	return
}