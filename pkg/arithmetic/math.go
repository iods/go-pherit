package arithmetic

/*
Add Returns the sum of two numbers for testing tests. */
func Add(a float64, b float64) float64 {
	return a + b
}

/*
Subtract Returns the difference of two numbers for testing tests. */
func Subtract(a float64, b float64) float64 {
	return a - b
}

/*
Multiply Returns the product (*) of two numbers. */
func Multiply(a float64, b float64) float64 {
	return a * b
}

/*
Divide Returns the quotient of two numbers. */
func Divide(a float64, b float64) float64 {
	return a / b
}

/*
Negate Returns the expression with it's sign changed.
 */
func Negate(a float64) float64 {
	return -(a)
}