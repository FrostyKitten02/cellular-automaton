package game

import (
	"strings"
)

func parseStringRule(rules string) ConwayRule {
	strs := strings.Split(rules, "/")
	bornStr := strings.Replace(strs[0], "B", "", -1)
	surviveStr := strings.Replace(strs[1], "S", "", -1)
	return ConwayRule{
		born:    parseStrNums(bornStr),
		survive: parseStrNums(surviveStr),
	}
}

func ruleApplies(counts NeighbourCounts, vals []int) bool {
	for _, val := range vals {
		if counts.alive == val {
			return true
		}
	}

	return false
}

func parseStrNums(str string) []int {
	nums := make([]int, len(str))

	for i, r := range str {
		nums[i] = int(r - '0')
	}

	return nums
}
