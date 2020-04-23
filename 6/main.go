package main

import "fmt"

func convert(s string, numRows int) string {
	m := [][]rune{}
	for i := 1; i <= numRows; i++ {
		m = append(m, []rune{})
	}
	row := 0
	xiangshang := false
	for _, c := range s {
		m[row] = append(m[row], c)
		if xiangshang {
			if row == 0 {
				xiangshang = false
			}
		} else {
			if row == (numRows - 1) {
				xiangshang = true
			}
		}
		if xiangshang {
			if (row - 1) >= 0 {
				row--
			}
		} else {
			if (row + 1) < numRows {
				row++
			}
		}
	}
	ret := ""
	for _, cs := range m {
		ret = fmt.Sprintf("%s%s", ret, string(cs))
	}
	return ret
}
