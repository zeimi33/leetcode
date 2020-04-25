package main

import "fmt"

func main() {

}

func intToRoman(num int) string {
	m := [4][10]string{{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"}}
	s := ""
	s = fmt.Sprintf("%s%s", m[3][num/1000], s)
	s = fmt.Sprintf("%s%s", s, m[3][(num/100)%10])
	s = fmt.Sprintf("%s%s", s, m[3][(num/10)%10])
	s = fmt.Sprintf("%s%s", s, m[3][num/1000])
	return s
}
