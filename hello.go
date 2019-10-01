package main

import (
	"fmt"
	"sort"
	"strings"
)

type RuneSlice []rune
func (r RuneSlice) Len() int { 
	return len(r)
}
func (r RuneSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r RuneSlice) Less(i, j int) bool {
	return r[i] < r[j]
}

func getKey(s string) string {
	r := (RuneSlice)(strings.ToLower(s))
	sort.Sort(r)
	return string(r)
}

type multimap map[string] []string
func (data multimap)multimap_insert(val string) {
	key := getKey(val)
	lst, ok := data[key]
	if ok {
		for _, el := range lst {
			if el == val {
				return
			}
		}
		data[key] = append(lst, val)
	} else {
		data[key] = append(make([]string, 0), val)
	}
}
func (data multimap)multimap_get(val string) (ret []string) {
	ret, _ = data[getKey(val)]
	return
}
func (data multimap)multimap_retrieve() (ret []string) {
	for _, v := range data {
		ret = append(ret, v...)
	}
	return
}

type registry struct{
	indices map[string] []int
	dict []string
}
func (data registry)registry_insert(val string) {
	key := getKey(val)
	indices, ok := data.indices[key]
	if ok {
		for _, index := range indices {
			if data.dict[index] == val {
				return
			}
		}
		data.indices[key] = append(indices, len(data.dict))
		data.dict = append(data.dict, val) 
	} else {
		data.indices[key] = append(make([]int, 0), len(data.dict))
		data.dict = append(data.dict, val) 
	}
}
func (data registry)registry_get(val string) (ret []string) {
	indices, _ := data.indices[getKey(val)]
	for _, index := range indices {
		ret = append(ret, data.dict[index])
	} 
	return
}
func (data registry)registry_retrieve() (ret []string) {
	return data.dict
}

var dict[]string = []string{"foobar", "barfoo", "boofar", "живу", "вижу", "Abba", "BaBa", "abba", "bba"}

func main() {
	data := make(multimap) 
	for _, val := range dict {
		data.multimap_insert(val)
	}
	for _, val := range dict {
		fmt.Println(val, data.multimap_get(val))
	}
}