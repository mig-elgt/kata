package grind75

func maxProduct(nums []int) int {
	products := [][]int{}
	for i := 0; i < len(nums); i++ {
		row := make([]int, len(nums))
		products = append(products, row)
	}
	max := nums[0]
	for i, num := range nums {
		products[0][i] = num
		if max < num {
			max = num
		}
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			products[i][j] = products[i-1][j-1] * nums[j]
			if max < products[i][j] {
				max = products[i][j]
			}
		}
	}
	return max
}
