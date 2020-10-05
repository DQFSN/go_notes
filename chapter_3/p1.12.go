package main

import "fmt"

func Solution(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	dict := make(map[rune]int)

	for _, v := range s1 {
		if _, ok := dict[v]; ok {
			dict[v]++
		}else {
			dict[v] = 1
		}
	}

	for _, v := range s2 {
		if _, ok := dict[v]; ok {
			dict[v]--
		}else {
			return false
		}
	}

	for _, v := range dict {
		if v != 0 {
			return false
		}
	}

	return true
}


func main() {
	s1 := ""
	s2 := ""

	fmt.Println(Solution(s1, s2))
}
