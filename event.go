package atomicstream

// Event has all info about the new event emmited
type Event struct {
	EventPayload interface{}
	Aggregate    interface{}
	IsException  bool
}

// Apply generates a new event with all nesessary data
func Apply(aggregate, event interface{}) Event {
	return Event{
		Aggregate:    aggregate,
		EventPayload: event,
		IsException:  false,
	}
}
