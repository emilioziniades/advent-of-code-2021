package day1

// CountInc counts the number of times a measurement
// increases from the previous measurement
func CountInc(in []int) (out int) {
	for i := 0; i < len(in); i++ {
		switch {
		case i == 0:
			continue
		case in[i] <= in[i-1]:
			continue
		case in[i] > in[i-1]:
			out++
		}
	}
	return
}

// CountIncThree counts the number of times the sum of the three
// most recent measurement increases from the previous set of three
func CountIncThree(in []int) (out int) {
	for i := 0; i < len(in); i++ {
		if i < 3 {
			continue
		}
		curr, prev := sum(in[i], in[i-1], in[i-2]), sum(in[i-1], in[i-2], in[i-3])
		switch {
		case curr <= prev:
			continue
		case curr > prev:
			out++
		}
	}
	return
}

func sum(nums ...int) (tot int) {
	for _, n := range nums {
		tot += n
	}
	return
}
