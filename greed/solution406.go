package greed

import (
	"container/list"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	l := list.New()
	for _, person := range people {
		cnt := person[1]
		mark := l.PushBack(person)
		cur := l.Front()
		for cnt > 0 {
			cur = cur.Next()
			cnt--
		}
		l.MoveBefore(mark, cur)
	}
	res := make([][]int, 0)

	for ele := l.Front(); ele != nil; ele = ele.Next() {
		res = append(res, ele.Value.([]int))
	}

	return res
}
