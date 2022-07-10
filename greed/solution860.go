package greed

func lemonadeChange(bills []int) bool {

	diff := 0
	cnt := make(map[int]int)

	for i := 0; i < len(bills); i++ {
		diff = bills[i] - 5
		if diff == 5 {
			if cnt[diff] <= 0 {
				return false
			}
			cnt[diff]--
		}
		if diff == 15 {
			if cnt[10] > 0 && cnt[5] > 0 {
				cnt[10]--
				cnt[5]--
			} else if cnt[5] >= 3 {
				cnt[5] = cnt[5] - 3
			} else {
				return false
			}
		}
		cnt[bills[i]]++
	}
	return true
}
