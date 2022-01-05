package day3

import (
	"log"
	"strconv"
)

func GammaDeltaProd(bits []string) int {
	var gamma, delta string
	for i := 0; i < len(bits[0]); i++ {
		var one, zero int
		for _, bit := range bits {
			switch string(bit[i]) {
			case "0":
				zero++
			case "1":
				one++
			}
		}
		switch {
		case one > zero:
			gamma += "1"
			delta += "0"
		case zero > one:
			gamma += "0"
			delta += "1"
		}
	}
	gammaInt, err := strconv.ParseInt(gamma, 2, 0)
	if err != nil {
		log.Fatalf("Gamma: %s\n", err)
	}
	deltaInt, err := strconv.ParseInt(delta, 2, 0)
	if err != nil {
		log.Fatalf("Delta: %s\n", err)
	}
	return int(gammaInt * deltaInt)
}

func OxygenCarbonDioxideRating(bits []string) int {
	oxygen := findMatches(bits, true, "1", 0)
	oxygenInt, err := strconv.ParseInt(oxygen, 2, 0)
	if err != nil {
		log.Fatalf("OxygenRating: converting to int: %s", err)
	}

	carbonDioxide := findMatches(bits, false, "0", 0)
	carbonDioxideInt, err := strconv.ParseInt(carbonDioxide, 2, 0)
	if err != nil {
		log.Fatalf("CarbonDioxideRating: converting to int: %s", err)
	}

	return int(oxygenInt * carbonDioxideInt)
}

func findMatches(bits []string, mostCommon bool, tiebreak string, position int) string {
	if len(bits) == 1 {
		return bits[0]
	}

	zeros := make([]string, 0)
	ones := make([]string, 0)

	for _, bit := range bits {
		switch s := string(bit[position]); s {
		case "0":
			zeros = append(zeros, bit)
		case "1":
			ones = append(ones, bit)
		}
	}
	match := compareOneZero(ones, zeros, mostCommon, tiebreak)
	return findMatches(match, mostCommon, tiebreak, position+1)

}

func compareOneZero(ones, zeros []string, mostCommon bool, tiebreak string) []string {
	switch lz, lo := len(zeros), len(ones); {
	case lz > lo && mostCommon:
		return zeros
	case lz > lo && !mostCommon:
		return ones
	case lo > lz && mostCommon:
		return ones
	case lo > lz && !mostCommon:
		return zeros
	case lo == lz && tiebreak == "0":
		return zeros
	case lo == lz && tiebreak == "1":
		return ones
	default:
		return nil
	}
}
