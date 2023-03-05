package validator

import (
	"log"
	"sync"

	validatorv10 "github.com/go-playground/validator/v10"
)

var (
	validatorMutex = &sync.Mutex{}
	validator      *validatorv10.Validate
)

func instance() *validatorv10.Validate {
	if validator != nil {
		return validator
	}

	validatorMutex.Lock()
	defer validatorMutex.Unlock()
	if validator != nil {
		return validator
	}

	validator = validatorv10.New()
	return validator
}

func Validate(in interface{}) bool {
	err := instance().Struct(in)
	if err == nil {
		return true
	}
	validationErrors, ok := err.(validatorv10.ValidationErrors)
	if ok {
		log.Printf("invalid validator err: %v, in: %v\n", validationErrors, in)
	}
	return false
}
