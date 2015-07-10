package main

// return smaller number.
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// return bigger number.
func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// alpha will clamped into 0 to 1.
// if alpha is 0 return value is a.
// if alpha is 1 return value is b.
func mix(a, b, alpha float64) float64 {
	if alpha > 1 {
		alpha = 1
	} else if alpha < 0 {
		alpha = 0
	}
	return a * (1-alpha) + b * alpha
}
