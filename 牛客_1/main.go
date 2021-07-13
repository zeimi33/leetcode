package main

import "fmt"

func main(){
	fmt.Printf(solve("999999999999","1"))
}
func solve( s string ,  t string ) string {
	if len(s) < len(t){
		return solve(t,s)
	}
	ret := ""
	jin := 0
	l1 := len(s)-1
	l2 := len(t)-1
	for i:=0;i< len(t);i++{
		n1 := int(s[l1-i] - '0')
		n2 := int(t[l2-i] - '0')
		n3 := (n1+n2+jin)%10
		jin = (n1+n2+jin)/10
		ret = fmt.Sprintf("%d%s",n3,ret)
	}
	if l1 == l2   {
		if jin!= 0 {
			return fmt.Sprintf("1%s", ret)
		}else{
			return ret
		}
	}
	for i:= l1-l2-1;i>=0;i--{
		n1 := int(s[i]-'0')
		n3 := (n1+jin) %10
		jin = (n1+jin) /10
		ret = fmt.Sprintf("%d%s",n3,ret)
	}
	if jin != 0 {
		ret = fmt.Sprintf("1%s",ret)
	}
	return ret
}
