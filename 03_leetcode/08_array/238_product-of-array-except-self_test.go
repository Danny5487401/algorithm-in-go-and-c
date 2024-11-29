package _8_array

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 除自身以外数组的乘积 https://leetcode.cn/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-100-liked
// 不要使用除法

func TestProductExceptSelf(t *testing.T) {
	convey.Convey("除自身以外数组的乘积", t, func() {
		testCase := []struct {
			input  []int
			target []int
		}{
			{

				[]int{1, 2, 3, 4}, []int{24, 12, 8, 6},
			},

			{

				[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0},
			},
		}

		for _, tst := range testCase {
			rsp := productExceptSelf(tst.input)
			convey.So(rsp, convey.ShouldResemble, tst.target)
		}
	})

}

func productExceptSelf(nums []int) []int {
	/*
		    定义 pre[i] 表示从 nums[0] 到 nums[i−1] 的乘积。
		    定义 suf[i] 表示从 nums[i+1] 到 nums[n−1] 的乘积

			pre[i]=pre[i−1]⋅nums[i−1]
			suf[i]=suf[i+1]⋅nums[i+1]
			answer[i]=pre[i]⋅suf[i]
	*/
	var pre = make([]int, len(nums))
	pre[0] = 1
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}

	var suf = make([]int, len(nums))
	suf[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		suf[i] = suf[i+1] * nums[i+1]
	}
	var ans = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = pre[i] * suf[i]
	}
	return ans

}
