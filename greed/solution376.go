package greed

func wiggleMaxLength(nums []int) int {
	cnt := 1
	preDif := 0
	for i := 1; i < len(nums); i++ {
		curDif := nums[i] - nums[i-1]
		if (curDif > 0 && preDif <= 0) || (curDif < 0 && preDif >= 0) {
			cnt++
			preDif = curDif
		}
	}
	return cnt
}
