package piscine

func Abort(a, b, c, d, e int) int {
	nums := [5]int{a, b, c, d, e}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4-i; j++ {
			if nums[j] > nums[j-i] {
				nums[j], nums[j+i] = nums[j+i], nums[j]
			}
		}
	}
	return nums[2]
}
