package main

func main(){

}
func isMatch(s string, p string) bool {
	index := 0
	xing := 0
	var xingRune rune
	for i := range s{
		if index > len(p){
			return false
		}
		if p[index] == '.'{
			index++
			continue
		}
		
	}
}