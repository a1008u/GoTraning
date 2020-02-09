package LongestCommonPrefix

import "strings"

//THE PURPOSE
// Approach: Horizontal scanning
//INTENTION
// horizontal scanningの理解
// sliceの一番最初の文字列を基準に、sliceにつづく文字列順番に同じ文字を調査している
func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0{
		return ""
	}

	preStr := strs[0]
	for i:=1; i<len(strs); i++{
		for strings.Index(strs[i], preStr) != 0 {
			preStr = preStr[0: len(preStr) - 1]
			if len(preStr) == 0 {
				return ""
			}
		}
	}
	return preStr
}

//THE PURPOSE
// Approach 2: Vertical scanning
//INTENTION
// horizontal scanningの理解
// sliceの各文字列を一文字づつ確認する
func LongestCommonPrefixV2(strs []string) string {

	if len(strs) == 0 || strs == nil{
		return ""
	}

	for i:=0; i<len(strs[0]); i++{
		c := string(strs[0][i])
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || string(strs[j][i]) != c{
				return string(strs[0][0: i])
			}
		}
	}
	return strs[0]
}
