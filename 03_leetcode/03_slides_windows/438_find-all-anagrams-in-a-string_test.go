package _3_slides_windows

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 找到字符串中所有字母异位词 https://leetcode.cn/problems/find-all-anagrams-in-a-string/?envType=study-plan-v2&envId=top-100-liked
func TestFindAnagrams(t *testing.T) {
	convey.Convey("找到字符串中所有字母异位词:字母异位词是通过重新排列不同单词或短语的字母而形成的单词或短语，并使用所有原字母一次", t, func() {
		testCase := []struct {
			s        string
			p        string
			expected []int
		}{
			{
				"cbaebabacd", "abc", []int{0, 6},
			},
			{
				"abab", "ab", []int{0, 1, 2},
			},
		}

		for _, tst := range testCase {
			rsp := findAnagrams(tst.s, tst.p)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

// 方法一: 不定长滑窗
func findAnagrams(s string, p string) []int {
	countMap := map[byte]int{}
	for index, _ := range p {
		countMap[p[index]]++
	}

	var ans []int
	var left int
	for right, subStr := range s {
		countMap[byte(subStr)]--
		for countMap[byte(subStr)] < 0 { // 判断是否不存在,或则说字母太多了
			// 说明不存在

			// 恢复数据
			countMap[s[left]]++
			// 左端点右移动
			left += 1
		}
		if right-left+1 == len(p) { // s 和 p 的每种字母的出现次数都相同
			ans = append(ans, left)
		}

	}

	return ans
}
