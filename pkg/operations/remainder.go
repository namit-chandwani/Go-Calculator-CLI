package operations

// RemainderConfig stores the metadata for the remainder sub-command
var RemainderConfig = &CommandConfig{"remainder", "Remainder", "", Remainder}

// Remainder takes 2 integer values as arguments and returns their remainder result
func Remainder(x int, y int) int {
	return x % y
}
