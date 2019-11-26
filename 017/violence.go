package main

import "fmt"

func main(){
	a := "DID"
	fmt.Println(numPermsDISequence(a))
}

//func numPermsDISequence(S string) int {
//	num := len(S)
//	var array []int
//	for i:= 0;i<num+1;i++{
//		array =append(array,i)
//	}
//	sum :=0
//	for i:=0;i<num+1;i++{
//		sum += numreturn(S,i,retSlice(array,i))
//	}
//	return sum
//}

func retSlice (array []int,index int)[]int{
	a := []int{}
	for i:=0 ;i<len(array) ;i++{
		if i !=index {
			a = append(a,array[i])
		}
	}
	return a
}
func numreturn (S string,index int,array []int)int{
	if len(array)==1{
		if S[0]=='D' && array[0]<index{
			return 1
		}
		if S[0]=='I' && array[0]>index{
			return 1
		}
		return 0
	}
	num :=0
	if S[0] =='D'{
		for i:=0;i<len(array);i++{
			if array[i]<index {
				num += numreturn(S[1:],array[i],retSlice(array,i))
			}else{
				return num
			}
		}
		return num
	}else {
		for i:=len(array)-1;i>=0;i--{
			if array[i]>index {
				num += numreturn(S[1:],array[i],retSlice(array,i))
			}else{
				return num
			}
		}
		return num
	}
}