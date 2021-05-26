package operations

// AddConfig stores the metadata for the add sub-command
var AddConfig = &CommandConfig{"add", "Addition", "", Add}

// Add takes 2 integer values as arguments and returns their addition result
func Add(x int, y int) int {
	return x + y
}
