package lcm

func FindLCM(nums []int) int {
    lcm := 1
    for i := 0; i < len(nums); i++ {
        lcm = LCM(lcm, nums[i])
    }

    return lcm
}

func LCM(a, b int) int {
    return a * b / GCD(a, b)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
