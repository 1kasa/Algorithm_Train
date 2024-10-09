package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, currentGas, startIndex := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i] - cost[i]
		currentGas += gas[i] - cost[i]

		if currentGas < 0 {
			startIndex = i + 1
			currentGas = 0
		}
	}

	if totalGas >= 0 {
		return startIndex
	} else {
		return -1
	}
}

func canCompleteCircuitOther(gas []int, cost []int) int {
	var ans, s, minS int

	for i, v := range gas {
		s += v - cost[i]
		if s < minS {
			minS = s
			ans = i + 1
		}
	}

	if s < 0 {
		return -1
	}

	return ans
}
