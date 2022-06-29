package backtrack

import (
	"sort"
)

type pair struct {
	name    string
	visited bool
}

type pairs []*pair

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pairs) Less(i, j int) bool {
	return p[i].name < p[j].name
}

func findItinerary(tickets [][]string) []string {
	var result []string
	hash := make(map[string]pairs)
	for _, ticket := range tickets {
		if hash[ticket[0]] == nil {
			hash[ticket[0]] = make(pairs, 0)
		}
		hash[ticket[0]] = append(hash[ticket[0]], &pair{
			name:    ticket[1],
			visited: false,
		})
	}

	for _, p := range hash {
		sort.Sort(p)
	}
	result = append(result, "JFK")
	var backtrack func() bool
	backtrack = func() bool {

		if len(tickets)+1 == len(result) {
			return true
		}

		for _, p := range hash[result[len(result)-1]] {
			if !p.visited {
				result = append(result, p.name)
				p.visited = true
				if backtrack() {
					return true
				}
				p.visited = false
				result = result[:len(result)-1]
			}
		}
		return false
	}
	backtrack()
	return result
}
