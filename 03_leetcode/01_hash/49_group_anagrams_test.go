package main

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/smartystreets/goconvey/convey"
	"sort"
	"testing"
)

// 字母异位词分组 https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked
func TestGroupAnagrams(t *testing.T) {
	convey.Convey("字母异位词分组 ", t, func() {
		testCase := []struct {
			input    []string
			expected [][]string
		}{
			{
				[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{
					{"bat"},
					{"nat", "tan"},
					{"ate", "eat", "tea"},
				},
			},

			{
				[]string{}, [][]string{},
			},
		}

		for _, tst := range testCase {
			ret := groupAnagrams(tst.input)
			convey.ShouldBeTrue(cmp.Equal(ret, tst.expected, cmpopts.SortSlices(func(x, y int) bool { return x < y })), true)
		}
	})

}

// 先按字母排序，排序后的字母为hash key
func groupAnagrams(strs []string) [][]string {
	var keyMap = map[string][]string{}
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		sortedStr := string(s)
		keyMap[sortedStr] = append(keyMap[sortedStr], str)
	}
	ans := make([][]string, 0, len(keyMap))
	for _, v := range keyMap {
		ans = append(ans, v)
	}

	return ans
}
