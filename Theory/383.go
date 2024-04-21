package main

import "fmt"

//	func canConstruct(ransomNote string, magazine string) bool {
//		runes := make(map[rune]int)
//		for _, byt := range magazine {
//			runes[byt]++
//		}
//		for _, byt := range ransomNote {
//			if _, ok := runes[byt]; ok == true {
//				if runes[byt] != 0 {
//					runes[byt]--
//				} else {
//					return false
//				}
//
//			} else {
//				return false
//			}
//		}
//		return true
//	}
func canConstruct(ransomNote string, magazine string) bool {
	ints := make([]int, 26)
	for _, i := range magazine {
		ints[i-'a']++
	}
	for _, i := range ransomNote {
		if ints[i-'a'] == 0 {
			return false
		} else {
			ints[i-'a']--
		}

	}
	return true
}
func main() {
	fmt.Println(canConstruct("aa", "aab"))
}
