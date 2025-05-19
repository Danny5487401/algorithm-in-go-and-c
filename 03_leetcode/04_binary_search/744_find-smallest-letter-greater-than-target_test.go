package _4_binary_search

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 寻找比目标字母大的最小字母 https://leetcode.cn/problems/find-smallest-letter-greater-than-target/description/
func TestNextGreatestLetter(t *testing.T) {
	convey.Convey("给你一个字符数组 letters，该数组按非递减顺序排序.如果不存在这样的字符，则返回 letters 的第一个字符。", t, func() {
		testCase := []struct {
			input    []byte
			target   byte
			expected byte
		}{
			{
				[]byte{'c', 'f', 'j'}, 'a', 'c',
			},
			{
				[]byte{'c', 'f', 'j'}, 'c', 'f',
			},
			{
				[]byte{'x', 'x', 'y', 'y'}, 'z', 'x',
			},
		}

		for _, tst := range testCase {
			rsp := nextGreatestLetter(tst.input, tst.target)
			convey.So(rsp, convey.ShouldResemble, tst.expected)
		}
	})

}

func nextGreatestLetter(letters []byte, target byte) byte {
	// 求 >x -->  >= x+1
	index := lowerCharBound(letters, target+1)
	if index == len(letters) { // 如果不存在这样的字符，则返回 letters 的第一个字符。
		return letters[0]
	}
	return letters[index]

}

func lowerCharBound(letters []byte, target byte) int {
	left := 0
	right := len(letters) - 1
	for left <= right {
		mid := left + (right-left)/2
		if letters[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
