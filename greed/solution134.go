package greed

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	sum := 0
	cur := 0
	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		cur += gas[i] - cost[i]
		if cur < 0 {
			cur = 0
			start = i + 1
		}
	}
	if sum < 0 {
		return -1
	}
	return start
}
