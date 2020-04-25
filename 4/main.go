package main

import (
	"fmt"
	"math"
)

func main() {
	a := findMedianSortedArrays([]int{}, []int{2, 3})
	fmt.Println(a)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m == 0 {
		return middle(nums2)
	}
	if n == 0 {
		return middle(nums1)
	}
	if n < m {
		return findMedianSortedArrays(nums2, nums1)
	}
	l := 0
	r := m
	k := (m + n + 1) / 2
	for r > l {
		slice := l + (r-l)/2
		slice2 := k - slice
		if slice2 < 0 {
			break
		}
		if nums1[slice] < nums2[slice2-1] {
			l = slice + 1
		} else {
			r = slice
		}
	}
	slice := l
	slice2 := k - slice
	lmax := 0
	rmin := 0
	if slice == 0 {
		lmax = nums2[slice2-1]
	} else if slice2 == 0 {
		lmax = nums1[slice-1]
	} else {
		lmax = int(math.Max(float64(nums1[slice-1]), float64(nums2[slice2-1])))
	}
	if slice == m {
		rmin = nums2[slice2]
	} else if slice2 == n {
		rmin = nums1[slice]
	} else {
		rmin = int(math.Min(float64(nums1[slice]), float64(nums2[slice2])))
	}
	fmt.Println(lmax, rmin)
	if (m+n)%2 != 0 {
		return float64(lmax)
	}
	return (float64(lmax) + float64(rmin)) / 2
}

func middle(a []int) float64 {
	n := len(a)
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return float64(a[n/2-1]+a[n/2]) / 2.0
	}
	return float64(a[n/2])
}
