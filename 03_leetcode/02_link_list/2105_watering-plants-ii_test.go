package _2_link_list

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

// 给植物浇水 II https://leetcode.cn/problems/watering-plants-ii/description/
func TestMinimumRefill(t *testing.T) {
	convey.Convey("找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水", t, func() {
		testCase := []struct {
			input     []int
			capacityA int
			capacityB int
			expected  int
		}{
			{
				/*
					- 最初，Alice 和 Bob 的水罐中各有 5 单元水。
					- Alice 给植物 0 浇水，Bob 给植物 3 浇水。
					- Alice 和 Bob 现在分别剩下 3 单元和 2 单元水。
					- Alice 有足够的水给植物 1 ，所以她直接浇水。Bob 的水不够给植物 2 ，所以他先重新装满水，再浇水。
					所以，两人浇灌所有植物过程中重新灌满水罐的次数 = 0 + 0 + 1 + 0 = 1 。
				*/
				[]int{2, 2, 3, 3}, 5, 5, 1,
			},
			{
				/*
					- 最初，Alice 的水罐中有 3 单元水，Bob 的水罐中有 4 单元水。
					- Alice 给植物 0 浇水，Bob 给植物 3 浇水。
					- Alice 和 Bob 现在都只有 1 单元水，并分别需要给植物 1 和植物 2 浇水。
					- 由于他们的水量均不足以浇水，所以他们重新灌满水罐再进行浇水。
					所以，两人浇灌所有植物过程中重新灌满水罐的次数 = 0 + 1 + 1 + 0 = 2
				*/
				[]int{2, 2, 3, 3}, 3, 4, 2,
			},
			{
				/*
					- 只有一株植物
						- Alice 的水罐有 10 单元水，Bob 的水罐有 8 单元水。因此 Alice 的水罐中水更多，她会给这株植物浇水。
						所以，两人浇灌所有植物过程中重新灌满水罐的次数 = 0
				*/
				[]int{5}, 10, 8, 0,
			},
			{
				[]int{1, 2, 4, 4, 5}, 6, 5, 2,
			},
			{
				[]int{274, 179, 789, 417, 293, 336, 133, 334, 569, 355, 813, 217, 80, 933, 961, 271, 294, 933, 49, 980, 685, 470, 186, 11, 157, 889, 299, 493, 215, 807, 588, 464, 218, 248, 391, 817, 32, 606, 740, 941, 505, 533, 289, 306, 490}, 996, 1172, 23,
			},
		}

		for _, tst := range testCase {
			rsp := minimumRefill(tst.input, tst.capacityA, tst.capacityB)
			convey.So(rsp, convey.ShouldEqual, tst.expected)
		}
	})

}

func minimumRefill(plants []int, capacityA int, capacityB int) int {

	var ans = 0
	var capacityALeft, capacityBLeft = capacityA, capacityB
	left, right := 0, len(plants)-1
	for left < right {

		// 如果 Alice/Bob 水罐中的水足以 完全 灌溉植物，他们 必须 给植物浇水。否则，他们 首先（立即）重新装满罐子，然后给植物浇水。
		if capacityALeft < plants[left] {
			ans++
			capacityALeft = capacityA
		}
		capacityALeft = capacityALeft - plants[left]
		left++
		if capacityBLeft < plants[right] {
			ans++
			capacityBLeft = capacityB
		}
		capacityBLeft = capacityBLeft - plants[right]
		right--
	}
	// 如果 Alice 和 Bob 到达同一株植物，那么当前水罐中水 更多 的人会给这株植物浇水。
	if left == right && max(capacityALeft, capacityBLeft) < plants[left] {
		ans++
	}
	return ans
}
