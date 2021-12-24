package main

import (
	"fmt"
	"reflect"
)

func sherlockAndAnagrams(s string) int32 {
	var numPairs int32 = 0

	size := len(s)
	if size <= 1 {
		return numPairs
	}
	foundPairItems := make([]string, 0)

	//isAnagram(s[:size-1], s[1:])
	substrs := getAllSubStrings(s, len(s)-1, make(map[int][]string, 0))
	for _, substrList := range substrs {
		for idxA, substrA := range substrList {
			for idxB, substrB := range substrList {
				isAnagram := isAnagram(substrA, substrB)
				//listHasPair := listContains(substrB, foundPairItems)
				if idxB != idxA && isAnagram {
					numPairs++
					foundPairItems = append(foundPairItems, substrA)

					if substrA != substrB {
						foundPairItems = append(foundPairItems, substrB)
					}
				}
			}
		}
	}

	//fmt.Println(numPairs)
	return numPairs
}

func fasterSherlockAndAnagrams(s string) int32 {
	var numPairs int32 = 0

	size := len(s)
	if size <= 1 {
		return numPairs
	}
	//foundPairItems := make([]string, 0)

	//isAnagram(s[:size-1], s[1:])
	substrs := getAllSubStrings(s, len(s)-1, make(map[int][]string, 0))
	for _, substrList := range substrs {
		for idxA, substrA := range substrList {
			for idxB, substrB := range substrList {
				isAnagram := isAnagram(substrA, substrB)
				//listHasPair := listContains(substrB, foundPairItems)
				if idxB != idxA && isAnagram {
					numPairs++
				}
			}
		}
	}

	fmt.Println("ran")
	return numPairs
}

//append(slice[:s], slice[s+1:]...)
//}

func listContains(needle string, haystack []string) bool {
	for _, val := range haystack {
		if val == needle {
			return true
		}
	}
	return false
}

//length equal, frequency of each character is same
func isAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	charsA := getCharMap(a)
	charsB := getCharMap(b)

	//fmt.Println(charsA, charsB, reflect.DeepEqual(charsA, charsB))
	if reflect.DeepEqual(charsA, charsB) {
		//fmt.Println(a, b)
		return true
	}

	return false
}

func getAllSubStrings(s string, length int, substrings map[int][]string) map[int][]string {
	if length > 0 && length < len(s) {
		substrList := make([]string, 0)
		for i := 0; i+length < len(s)+1; i++ {
			substrList = append(substrList, s[i:i+length])
		}
		substrings[length] = substrList
		return getAllSubStrings(s, length-1, substrings)
	}
	return substrings
}

func getCharMap(s string) map[int32]int {
	m := make(map[int32]int)

	for _, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = 0
		}
		m[c]++
	}

	return m
}

func main() {
	/*
		res := sherlockAndAnagrams("kkkk")
		fmt.Println(res) //should be 10

		res = sherlockAndAnagrams("abba")
		fmt.Println(res) //should be 4

		res = sherlockAndAnagrams("abcd")
		fmt.Println(res) //should be 0

		res = sherlockAndAnagrams("ifailuhkqq")
		fmt.Println(res) //should be 3
	*/

	res := fasterSherlockAndAnagrams("dbcfibibcheigfccacfegicigcefieeeeegcghggdheichgafhdigffgifidfbeaccadabecbdcgieaffbigffcecahafcafhcdg")
	fmt.Println(res) //should be 1464

	res = fasterSherlockAndAnagrams("dfcaabeaeeabfffcdbbfaffadcacdeeabcadabfdefcfcbbacadaeafcfceeedacbafdebbffcecdbfebdbfdbdecbfbadddbcec")
	fmt.Println(res) //should be 2452

	res = fasterSherlockAndAnagrams("gjjkaaakklheghidclhaaeggllagkmblhdlmihmgkjhkkfcjaekckaafgabfclmgahmdebjekaedhaiikdjmfbmfbdlcafamjbfe")
	fmt.Println(res) //should be 873

	res = fasterSherlockAndAnagrams("fdbdidhaiqbggqkhdmqhmemgljaphocpaacdohnokfqmlpmiioedpnjhphmjjnjlpihmpodgkmookedkehfaceklbljcjglncfal")
	fmt.Println(res) //should be 585

	res = fasterSherlockAndAnagrams("bcgdehhbcefeeadchgaheddegbiihehcbbdffiiiifgibhfbchffcaiabbbcceabehhiffggghbafabbaaebgediafabafdicdhg")
	fmt.Println(res) //should be 1305

	res = fasterSherlockAndAnagrams("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	fmt.Println(res) //should be 166650

	res = fasterSherlockAndAnagrams("mhmgmbbccbbaffhbncgndbffkjbhmkfncmihhdhcebmchnfacdigflhhbekhfejblegakjjiejeenibemfmkfjbkkmlichlkbnhc")
	fmt.Println(res) //should be 840

	res = fasterSherlockAndAnagrams("fdacbaeacbdbaaacafdfbbdcefadgfcagdfcgbgeafbfbggdedfbdefdbgbefcgdababafgffedbefdecbaabdaafggceffbacgb")
	fmt.Println(res) //should be 2134

	res = fasterSherlockAndAnagrams("bahdcafcdadbdgagdddcidaaicggcfdbfeeeghiibbdhabdhffddhffcdccfdddhgiceciffhgdibfdacbidgagdadhdceibbbcc")
	fmt.Println(res) //should be 1571

	res = fasterSherlockAndAnagrams("dichcagakdajjhhdhegiifiiggjebejejciaabbifkcbdeigajhgfcfdgekfajbcdifikafkgjjjfefkdbeicgiccgkjheeiefje")
	fmt.Println(res) //should be 1042
	/*
			dbcfibibcheigfccacfegicigcefieeeeegcghggdheichgafhdigffgifidfbeaccadabecbdcgieaffbigffcecahafcafhcdg
			dfcaabeaeeabfffcdbbfaffadcacdeeabcadabfdefcfcbbacadaeafcfceeedacbafdebbffcecdbfebdbfdbdecbfbadddbcec
			gjjkaaakklheghidclhaaeggllagkmblhdlmihmgkjhkkfcjaekckaafgabfclmgahmdebjekaedhaiikdjmfbmfbdlcafamjbfe
			fdbdidhaiqbggqkhdmqhmemgljaphocpaacdohnokfqmlpmiioedpnjhphmjjnjlpihmpodgkmookedkehfaceklbljcjglncfal
			bcgdehhbcefeeadchgaheddegbiihehcbbdffiiiifgibhfbchffcaiabbbcceabehhiffggghbafabbaaebgediafabafdicdhg
			aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
			mhmgmbbccbbaffhbncgndbffkjbhmkfncmihhdhcebmchnfacdigflhhbekhfejblegakjjiejeenibemfmkfjbkkmlichlkbnhc
			fdacbaeacbdbaaacafdfbbdcefadgfcagdfcgbgeafbfbggdedfbdefdbgbefcgdababafgffedbefdecbaabdaafggceffbacgb
			bahdcafcdadbdgagdddcidaaicggcfdbfeeeghiibbdhabdhffddhffcdccfdddhgiceciffhgdibfdacbidgagdadhdceibbbcc
			dichcagakdajjhhdhegiifiiggjebejejciaabbifkcbdeigajhgfcfdgekfajbcdifikafkgjjjfefkdbeicgiccgkjheeiefje

		expected values:
		1464
		2452
		873
		585
		1305
		166650
		840
		2134
		1571
		1042
	*/

}
