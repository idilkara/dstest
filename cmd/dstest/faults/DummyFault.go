package faults

import (
	"fmt"
	"github.com/egeberkaygulcan/dstest/cmd/dstest/faults/behavior"
	"github.com/egeberkaygulcan/dstest/cmd/dstest/faults/trigger"
)

type DummyFault struct {
	BaseFault
}

type DummyFaultParams struct {
	name string
	age  int
}

var _ Fault = (*DummyFault)(nil)

func NewDummyFault(params map[string]interface{}) (*DummyFault, error) {
	fmt.Println("Creating a new DummyFault")

	if _, ok := params["name"]; !ok {
		return nil, fmt.Errorf("name parameter is required")
	}

	if _, ok := params["age"]; !ok {
		return nil, fmt.Errorf("age parameter is required")
	}

	parsedParams := &DummyFaultParams{
		name: params["name"].(string),
		age:  params["age"].(int),
	}

	fmt.Println("Creating a new DummyFault with params: ", parsedParams)

	return &DummyFault{
		BaseFault: BaseFault{
			FaultTrigger:  &trigger.DummyFaultTrigger{},
			FaultBehavior: &behavior.DummyFaultyBehavior{},
		},
	}, nil
}
