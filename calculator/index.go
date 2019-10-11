package calculator

func expN(x float64, n int) float64 {
	return pow(x, n) / (float64)(fact(n))
}

func Exp(x float64) float64 {
	var e float64
	for i := 0; i <= 20; i++ {
		e += expN(x, i)
	}

	return e
}

func Norm(x float64) float64 {
	return Exp(-pow(x, 2))
}

func sinN(x float64, n int) float64 {
	return pow(-1, n/2) * pow(x, n) / (float64)(fact(n))
}

const Pi = 3.141592653589793

func Sin(x float64) float64 {
	x = normTrig(x)

	var s float64
	for i := 1; i <= 20; i += 2 {
		s += sinN(x, i)
	}

	return s
}

func lnN(x float64, n int) float64 {
	return pow(-1, n-1) * pow(x-1, n) / (float64)(n)
}

func Ln(x float64) float64 {
	var l float64

	for i := 1; i < 200; i++ {
		l += lnN(x, i)
	}

	return l
}

func Cos(x float64) float64 {
	x = normTrig(x)

	var s float64
	for i := 0; i <= 20; i += 2 {
		s += sinN(x, i)
	}

	return s
}
