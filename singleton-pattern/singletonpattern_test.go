package singleton_pattern

import (
	"fmt"
	"strings"
	"testing"
)

func TestSingletonPattern(t *testing.T) {
	randomGenerator := newGenerator()
	randomGenerator.generateRandomInt()
}

func TestArr(t *testing.T) {

	fmt.Println(computeLPSArray("wwwabc"))
}

func computeLPSArray(pattern string) []int {
	length := 0
	lps := make([]int, len(pattern))
	lps[0] = 0

	i := 1
	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	start := 0
	end := len(s) - 1

	for i := 0; start < end; i++ {
		fmt.Println(start, end, s[start:start+1], int(s[start : start+1][0]))
		if !(int(s[start : start+1][0]) >= 97 && int(s[start : start+1][0]) <= 122) && !(int(s[start : start+1][0]) >= 48 && int(s[start : start+1][0]) <= 57) {
			start++
			continue
		}

		if !(int(s[end : end+1][0]) >= 97 && int(s[end : end+1][0]) <= 122) && !(int(s[end : end+1][0]) >= 48 && int(s[end : end+1][0]) <= 57) {
			end--
			continue
		}

		if s[start:start+1] != s[end:end+1] {
			return false
		}
		start++
		end--
	}
	return true
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	indexNums1 := 0
	indexNums2 := 0
	last := m

	for i := 0; i < len(nums1); i++ {
		fmt.Println(nums1)
		fmt.Println(indexNums1, indexNums2)
		if indexNums2 >= len(nums2) {
			break
		}

		if indexNums1 >= last {
			nums1[indexNums1] = nums2[indexNums2]
			indexNums1++
			indexNums2++
			continue
		}

		if nums1[indexNums1] <= nums2[indexNums2] {
			indexNums1++
			continue
		}

		nums1 = append(nums1[0:indexNums1], nums1[indexNums1-1:len(nums1)-1]...)
		nums1[indexNums1] = nums2[indexNums2]
		indexNums2++
		indexNums1++
		last++
	}
}

func longestCommonPrefix(strs []string) string {
	commonStr := ""

	for i := 0; ; i++ {
		if len(strs[0]) <= i {
			return commonStr
		}

		curStr := strs[0][i : i+1]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i {
				return commonStr
			}

			if curStr != strs[j][i:i+1] {
				return commonStr
			}
		}

		commonStr = commonStr + strs[0][i:i+1]
	}

	return commonStr
}
