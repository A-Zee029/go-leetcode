package greed

func candy(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}

	list := make([]int, 0)
	sum := 0
	for i := 0; i < len(ratings); i++ {
		if (i == 0 && ratings[i] <= ratings[i+1]) || (i == len(ratings)-1 && ratings[i] <= ratings[i-1]) || (i < len(ratings)-1 && i > 0 && ratings[i] <= ratings[i-1] && ratings[i] <= ratings[i+1]) {
			list = append(list, i)
		}
	}
	left := 0
	right := len(ratings) - 1
	for i := 0; i < len(ratings); i++ {
		if len(list) > 0 {
			right = list[0]

			if right == i {
				sum += 1
				left = i
				list = list[1:]
			} else {
				if i == 0 {
					sum += right - i + 1
				} else if i == len(ratings)-1 {
					sum += i - left + 1
				} else {
					if (ratings[i-1] > ratings[i] && ratings[i] >= ratings[i+1]) || (ratings[i-1] >= ratings[i] && ratings[i] > ratings[i+1]) {
						sum += right - i + 1
					} else if (ratings[i-1] < ratings[i] && ratings[i] <= ratings[i+1]) || (ratings[i-1] <= ratings[i] && ratings[i] < ratings[i+1]) {
						sum += i - left + 1
					} else if ratings[i-1] < ratings[i] && ratings[i] > ratings[i+1] {
						if right-i > i-left {
							sum += right - i + 1
						} else {
							sum += i - left + 1
						}
					}
				}

			}
		} else {
			sum += i - left + 1
		}

	}
	return sum
}

func main() {
	candy([]int{1, 0, 2, 3, 4, 1})
	//b := byte(9)
	//fmt.Printf("%v %T", b, b)
}
