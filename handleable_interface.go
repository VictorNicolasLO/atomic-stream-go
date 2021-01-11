package atomicstream

// Handleable is an interface for Events
type Handleable interface {
	Handle(val interface{})
}
