package operations

// MultiplyConfig stores the metadata for the multiply sub-command
var MultiplyConfig = &CommandConfig{"multiply", "Multiplication", "", Multiply}

// Multiply takes 2 integer values as arguments and returns their multiplication result
func Multiply(x int, y int) int {
	return x * y
}
