package main


import (
	"fmt"
)
func main(){
	fmt.Println(validPalindrome("deeee"))
}

func validPalindrome(s string) bool {
	length := len(s)
	l := 0 
	r := length-1
	for l <= r {
		if l == r {
			return true
		}
		if s[l] != s[r] {
			break
		}
		l++
		r--
	}
	c := []rune(s)
	f := []rune{}
	t := []rune{}
	f = append(f,c[:l]...)
	f = append(f,c[l+1:]...)
	t = append(t,c[:r]...)
	t = append(t,c[r+1:]...)
	return vaild(f)||vaild(t)
}
func vaild(s []rune)bool{
	fmt.Println(s)
	length := len(s)
	l := 0 
	r := length-1
	for l <= r {
		if l == r {
			return true
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}