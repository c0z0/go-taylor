package calculator

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func pow(x float64, p int) float64 {
	if p == 0 {
		return 1
	}

	return x * pow(x, p-1)
}

func normTrig(x float64) float64 {

	for x > Pi {
		x -= 2 * Pi
	}

	for x < -Pi {
		x += 2 * Pi
	}

	return x
}
