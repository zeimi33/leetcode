package main

import "fmt"

func main(){
	b := MountainArray{}
	c := findInMountainArray(0,&b)
	fmt.Println(c)
}


  type MountainArray struct {
  }
 var a = []int{3,5,3,2,0}
func (this *MountainArray) get(index int) int {return a[index]}
func (this *MountainArray) length() int {return 5}


func findInMountainArray(target int, mountainArr *MountainArray) int {
	lengh := mountainArr.length()
	l := 0
	r := lengh-1
	for l<r{
		fmt.Println(l,r)
		midv := mountainArr.get((l+r)/2)
			
			if (l+r)/2 >0{
				lv := mountainArr.get((l+r)/2-1)
			if lv >midv{
				r = (l+r)/2-1
			}else{
				l = (l+r)/2+1
			}
			continue
		}
		if (l+r)/2 < lengh-1 {
			lv := mountainArr.get((l+r)/2+1)
						if lv <midv{
				r = (l+r)/2-1
			}else{
				l = (l+r)/2+1
			}
			continue
		}
	}
	fmt.Println(l)
	peak := l
	l = 0
	r =peak
	for l < r {
									midv := mountainArr.get((l+r)/2)
		fmt.Println(midv,target)
		if midv > target {
			r = (l+r)/2-1
			continue
		}
		if midv < target{
			l = (l+r)/2+1
			continue
		}
		if midv == target {
			return (l+r)/2
		}
	}
	if mountainArr.get(l) == target {
		return l
	}
	l = peak
	r = lengh-1
	for l < r {
		midv := mountainArr.get((l+r)/2)
		if midv < target {
			r = (l+r)/2-1
			continue
		}
		if midv > target{
			l = (l+r)/2+1
			continue
		}
		if midv == target {
			return (l+r)/2
		}
	}
		if mountainArr.get(l) == target {
		return l
	}
	return -1
}