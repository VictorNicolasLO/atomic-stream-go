package atomicstream

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"
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
	Event
	Name  string
	Email string
}

func (useProspect UserProspect) startOnboardingProccess(cmd StartOnboardingProccess) {

}

func (command StartOnboardingProccess) Handle(userProspect UserProspect) OnboardingProccessStarted {
	return OnboardingProccessStarted{Name: command.Name}
}

func (event OnboardingProccessStarted) Handle(userProspect UserProspect) UserProspect {
	userProspect.Name = event.Name
	return userProspect
}

var commandHandlers = [...]func() interface{}{
	func() interface{} { return &StartOnboardingProccess{} },
}

type CommandRoot struct {
	Name    string
	Payload interface{}
}

func TestContext_Aggregator(t *testing.T) {
	// preload functions
	fooType := reflect.TypeOf(UserProspect{})

	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		t.Log(method.Name)
	}
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
	var payload json.RawMessage
	cmd := CommandRoot{
		Payload: &payload,
	}
	commandString := []byte(`{"name":"start-onboarding-proccess","payload":{"name":"Victor","email":"victor@hotmail.com"}}`)
	json.Unmarshal(commandString, &cmd)
	t.Log(cmd)

	t.Log(cmd.Payload)

	//commandPayload := commandHandlers[cmd.Name]()
	commandPayload := commandHandlers[0]()
	json.Unmarshal(payload, commandPayload)
	t.Log(reflect.TypeOf(commandPayload))
	reflected := reflect.ValueOf(commandPayload)
	t.Log(reflected.MethodByName("Handle").Type().NumIn())
	t.Log(reflected.MethodByName("Handle").Call([]reflect.Value{
		reflect.ValueOf(StartOnboardingProccess{}),
	}))
}
