package main

import "sort"

func main() {

}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		c := format(v)
		if _, ok := m[c]; ok {
			m[c] = append(m[c], v)
		} else {
			m[c] = []string{v}
		}
	}
	ret := [][]string{}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

type SortBy []rune

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }
func format(a string) string {
	s := SortBy(a)
	sort.Sort(s)
	return string(s)
}
