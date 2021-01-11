package atomicstream

import (
	"fmt"
)

// Aggregator is an object in charge of proccess commands and update state objects
type Aggregator struct {
	brokers []string
}

// Start method runs the aggregator
func (aggregator Aggregator) Start() {

}

// NewAggregator creates an aggregator instance with the most common options
func NewAggregator(model, brokers []string) {
	fmt.Print("Creating Aggregator")

}
