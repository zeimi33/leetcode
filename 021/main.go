package main

func main(){

}

func isPowerOfTwo(n int) bool {
	if n <= 0{
		return false
	}
	for n >=2{
		if n%2 ==0{
			n = n/2
		}
	}
	if n==1 {
		return true
	}
	return false
}