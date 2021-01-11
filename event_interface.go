package atomicstream

// Event is an interface for Events
type Event interface {
	Handle(aggregate interface{}) interface{}
	cloneInterface() interface{}
}
