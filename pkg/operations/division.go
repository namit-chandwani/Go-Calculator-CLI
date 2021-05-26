package operations

// DivideConfig stores the metadata for the divide sub-command
var DivideConfig = &CommandConfig{"divide", "Division", "", Divide}

// Divide takes 2 integer values as arguments and returns their division result
func Divide(x int, y int) int {
	return x / y
}
