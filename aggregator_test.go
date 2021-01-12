package atomicstream

import (
	"encoding/json"
	"github.com/VictorNicolasLO/atomic-stream-go/packageutils"
	"github.com/go-playground/validator/v10"
	"reflect"
	"testing"
)

// To use validations and json rename for json parsing
type UserProspect struct {
	Aggregate
	Name  string `validate:"required"`
	Email string
}

type StartOnboardingProccess struct {
	Command
	Name  string
	Email string
}

type OnboardingProccessStarted struct {
	Name  string
	Email string
}

func (userProspect UserProspect) StartOnboardingProccess(cmd StartOnboardingProccess) Event {
	userProspect.Name = cmd.Name
	userProspect.Email = cmd.Email
	return Apply(userProspect, OnboardingProccessStarted{Name: cmd.Name, Email: cmd.Email})
}

type CommandRoot struct {
	Name    string
	Payload interface{}
}

func TestContext_Aggregator(t *testing.T) {
	// preload functions

	t.Log("SOME")
	// Proccess message
	commandStr := []byte(`{"name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	var command UserProspect
	command = UserProspect{}
	err := json.Unmarshal(commandStr, &command)
	if err != nil {
		t.Log("error")
	}
	validate := validator.New()
	errStruct := validate.Struct(command)
	t.Log(errStruct)
	t.Log(command)

	// other tests
	str := []byte(`{"name":"Alice","Body":"Hello","Time":1294706395881547000,"email":"abc"}`)
	var jsonVar UserProspect
	jsonVar = UserProspect{}
	json.Unmarshal(str, &jsonVar)
	t.Log(jsonVar.Email)

	// Lets tryit
	commandString := []byte(`{"name":"start-onboarding-proccess","payload":{"name":"Victor","email":"victor@hotmail.com"}}`)
	var payload json.RawMessage
	cmd := CommandRoot{
		Payload: &payload,
	}
	json.Unmarshal(commandString, &cmd)
	t.Log(cmd)

	t.Log(cmd.Payload)

	// //commandPayload := commandHandlers[cmd.Name]()
	// commandPayload := commandHandlers[0]()
	// json.Unmarshal(payload, commandPayload)
	// t.Log(reflect.TypeOf(commandPayload))
	// reflected := reflect.ValueOf(commandPayload)
	// t.Log(reflected.MethodByName("Handle").Type().NumIn())
	// t.Log(reflected.MethodByName("Handle").Call([]reflect.Value{
	// 	reflect.ValueOf(StartOnboardingProccess{}),
	// }))

	// Setup aggregate

	// getting methods
	aggregateGenerator := func() interface{} { return &UserProspect{} }
	aggregateBase := aggregateGenerator()
	aggregateType := reflect.TypeOf(aggregateBase)
	reflectedValue := reflect.ValueOf(aggregateBase)
	commandTypeHandlers := map[string]func() interface{}{}
	methodHandlers := map[string]func(payload interface{}) interface{}{}

	for i := 0; i < aggregateType.NumMethod(); i++ {
		method := aggregateType.Method(i)
		methodValue := reflectedValue.MethodByName(method.Name)
		methodType := methodValue.Type()
		commandType := methodType.In(0)
		getCommandType := func() interface{} {
			return reflect.New(commandType).Elem()
		}
		methodHandler := func(payload interface{}) interface{} {
			return methodValue.Call([]reflect.Value{
				reflect.ValueOf(payload),
			})
		}

		kebabMethodName := packageutils.CamelCaseToKebabCase(method.Name)
		commandTypeHandlers[kebabMethodName] = getCommandType
		methodHandlers[kebabMethodName] = methodHandler
		t.Log(kebabMethodName)
	}

	commandHandler := func(commandStr string, currentValue interface{}) interface{} {
		var payload json.RawMessage
		rawCommand := CommandRoot{
			Payload: &payload,
		}
		json.Unmarshal(commandString, &rawCommand)
		commandPayload := commandTypeHandlers[rawCommand.Name]()
		json.Unmarshal(payload, commandPayload)

		t.Log(commandPayload)
	}

}

// https://stackoverflow.com/questions/47626052/decoding-arguments-using-reflection
