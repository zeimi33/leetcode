package main

import (
	"fmt"
	"sort"
)

func main(){
	a := []int{65,70,56,75,60,68}
	b := []int{100,150,90,190,95,110}
	c := bestSeqAtIndex(a,b)
	fmt.Println(c)
}

type  people struct{
	height int
	weight int
	num int
}
type peoples struct{
	pps []people
	by func(p1,p2 *people)bool
}
func (p peoples)Len()int{
	return len(p.pps)
}
func(p peoples)Swap(i,j int){
	p.pps[i],p.pps[j] = p.pps[j],p.pps[i]
}
func (p peoples)Less(i,j int)bool{
	return p.by(&p.pps[i],&p.pps[j])
}
func bestSeqAtIndex(height []int, weight []int) int {
	pps := []people{}
	for index,hei := range height{
		p := people{}
		p.height = hei
		p.weight = weight[index]
		p.num = 1
		pps = append(pps,p)
	}
	sort.Sort(peoples{pps,func(p1,p2 *people)bool{
			if p1.height > p2.height{
				return true
			}
			if p1.height < p2.height{
				return false
			}
			return  p1.weight >p2.weight
		}})

		ret := 0

	for index := range pps{
		max :=0
		for i:=0 ;i<index ;i++{
			if pps[i].weight > pps[index].weight&& pps[i].height>pps[index].height{
				if pps[i].num>max {
					max = pps[i].num
				}
			}
		}
		
		pps[index].num =max+1
		if ret < pps[index].num {
			ret = pps[index].num
		}
	}
	return ret
}





