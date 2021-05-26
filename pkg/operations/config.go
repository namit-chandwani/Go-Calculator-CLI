package operations

// CommandConfig represents the metadata for the subtract sub-command
type CommandConfig struct {
	Name             string
	ShortDescription string
	Examples         string
	OperationFunc    func(int, int) int
}
