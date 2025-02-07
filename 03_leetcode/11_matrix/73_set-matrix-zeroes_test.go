package _1_matrix

// 矩阵置零

// 如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0

// 方法一：使用标记数组
func setZeroes(matrix [][]int) {
	// 首先遍历该数组一次，如果某个元素为 0，那么就将该元素所在的行和列所对应标记数组的位置置为 true。
	// 最后我们再次遍历该数组，用标记数组更新原数组即可
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}
