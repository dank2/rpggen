package data

import (
	"fmt"
	"testing"
)

const (
	n                = 100000
	acceptedVariance = 0.1
)

func TestRaceDistribution(t *testing.T) {
	distribution := make(map[string]int)
	for i := 0; i < n; i++ {
		race := GetRace()

		if v, ok := distribution[race]; ok {
			v++
			distribution[race] = v
		} else {
			distribution[race] = 1
		}

	}

	for _, race := range Races {
		median := n * race.Distribution
		variance := median * acceptedVariance
		lowerBound := int(median - variance)
		upperBound := int(median + variance)

		rDist := distribution[race.Race]
		// fmt.Printf("Race: %s\nMedian: %f, Upper: %d, Lower: %d, Actual: %d\n", race.Race, median, upperBound, lowerBound, rDist)
		if rDist < lowerBound || rDist > upperBound {
			fmt.Printf("Race %s, did not fall within the bounds of %d and %d with %d\n", race.Race, lowerBound, upperBound, rDist)
			t.Fail()
		}
	}
}
