package main

import (
	"fmt"
	"sort"
)

	func main() {
		fmt.Println(peopleIndexes([][]string{{"arrtztkotazhufrsfczr","knzgidixqgtnahamebxf","zibvccaoayyihidztflj"},{"cffiqfviuwjowkppdajm","owqvnrhuzwqohquamvsz"},{"knzgidixqgtnahamebxf","owqvnrhuzwqohquamvsz","qzeqyrgnbplsrgqnplnl"},{"arrtztkotazhufrsfczr","cffiqfviuwjowkppdajm"},{"arrtztkotazhufrsfczr","knzgidixqgtnahamebxf","owqvnrhuzwqohquamvsz","qzeqyrgnbplsrgqnplnl","zibvccaoayyihidztflj"}}))
	}
//	{{"arrtztkotazhufrsfczr","knzgidixqgtnahamebxf","zibvccaoayyihidztflj"},{"cffiqfviuwjowkppdajm","owqvnrhuzwqohquamvsz"],{"knzgidixqgtnahamebxf","owqvnrhuzwqohquamvsz","qzeqyrgnbplsrgqnplnl"},{"arrtztkotazhufrsfczr","cffiqfviuwjowkppdajm"},{"arrtztkotazhufrsfczr","knzgidixqgtnahamebxf","owqvnrhuzwqohquamvsz","qzeqyrgnbplsrgqnplnl","zibvccaoayyihidztflj"}}
type so struct{
	fCP [][]string
	indexM []int
}
func (a so) Len() int           { return len(a.fCP) }
func (a so) Swap(i, j int){ 
	a.fCP[i], a.fCP[j] = a.fCP[j], a.fCP[i] 
	a.indexM[i],a.indexM[j] = a.indexM[j],a.indexM[i]
}
func (a so) Less(i, j int) bool { return len(a.fCP[i]) > len(a.fCP[j]) }
func peopleIndexes(favoriteCompanies [][]string) []int {
	mapOfFC := []int64{}
	indexM := []int{}
	for i := 0 ; i < len(favoriteCompanies);i++{
		indexM = append(indexM,i)
	}
	s := so{
		favoriteCompanies,
		indexM,
	}
	sort.Sort(s)
	index := 0
	mCover := make(map[int64]int)
	mString := make(map[string]int64)
	indexNum := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523, 541}
	for _,v := range favoriteCompanies{
		var mutil int64
		mutil = 1
		for _,company := range v {
			if mString[company] == 0 {
				mString[company] = int64(indexNum[index])
				index++
			}
			mutil = mutil* mString[company]
		}
		mCover[mutil]++
		mapOfFC = append(mapOfFC,mutil)
	}
	ret := []int{}
	de := false
	fmt.Println(mapOfFC,mCover,mString)
	for i,v := range mapOfFC{
		de = false
		//fmt.Println(len(mapOfFC),i)
		if mCover[v] >1 {
			continue
		}

		for j:=0;j<i;j++{
			if  mapOfFC[j]%v == 0 {
				de = true
				break
			}
		}
		if ! de {
			ret = append(ret,indexM[i])
		}
	}
	sort.Ints(ret)
	return ret
}	


