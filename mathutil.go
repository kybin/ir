package main

// return smaller number.
func minval(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// return bigger number.
func maxval(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// amount will clamped into 0 to 1.
// if amount is 0 return value is a.
// if amount is 1 return value is b.
func mix(a, b, amount float64) float64 {
	if amount <= 0 {
		return a
	}
	if amount >= 1 {
		return b
	}
	return a * (1-amount) + b * amount
}

// amount will clamped into 0 to 1.
// if amount is 0 return value is a.
// if amount is 1 return value is b.
func mixVector3(a, b vector3, amount float64) vector3 {
	if amount <= 0 {
		return a
	}
	if amount >= 1 {
		return b
	}
	return a.Mult(1-amount).Add(b.Mult(amount))
}
