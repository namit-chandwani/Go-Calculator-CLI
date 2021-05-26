package operations

// SubtractConfig stores the metadata for the subtract sub-command
var SubtractConfig = &CommandConfig{"subtract", "Subtraction", "", Subtract}

// Subtract takes 2 integer values as arguments and returns their subtraction result
func Subtract(x int, y int) int {
	return x - y
}
