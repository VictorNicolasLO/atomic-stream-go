package atomicstream

// Command is an interface for Events
type Command struct {
}

// Handle do something
func (cmd *Command) Handle(param interface{}) interface{} {
	return new(interface{})
}
