package echovalidate

import (
	"log"
	"reflect"
)

type (
	Validation struct{}
	Rules      [][]interface{}
	Validator  struct {
		Rules Rules
	}
)

func New() *Validation {
	return &Validation{}
}

func (v *Validation) Validate(i interface{}) error {

	switch i.(type) {
	case Validator:
		vv := i.(Validator)
		for _, rule := range vv.Rules {
			validationFunc := rule[0]

			if reflect.TypeOf(validationFunc).Kind() != reflect.Func {
				log.Panicf("echovalidate: First Member of validator.Rules must be a function, I got %v", reflect.TypeOf(validationFunc).Kind())
			}
			reflectedValidationFunc := reflect.ValueOf(validationFunc)

			validationArgs := rule[1:]
			reflectedValidationArgs := make([]reflect.Value, len(validationArgs))
			for i, a := range validationArgs {
				reflectedValidationArgs[i] = reflect.ValueOf(a)
			}

			err := reflectedValidationFunc.Call(reflectedValidationArgs)[0].Interface()
			if err != nil {
				return err.(error)
			}
		}
	default:
		log.Panic("echovalidate: argument passed to Validate() must be of type echovalidate.Validator")
	}
	return nil
}
