package user

import (
    "fmt"
    "sync"

    "github.com/go-playground/validator/v10"
)

// Validator is an interface that requires a Validate method.
type Validator interface {
    Validate() error
}

// structural validation of the User, SignUp and Login structs
var (
    // Use a single instance of Validate, it caches struct info.
    validate *validator.Validate
    once     sync.Once
)

func getValidator() *validator.Validate {
    once.Do(func() {
        validate = validator.New()
    })
    return validate
}

// Helper function to perform validation and print errors.
func printValidationErrors(err error) error {
    validationErrors, ok := err.(validator.ValidationErrors)
    if !ok {
        return err // Not validation errors
    }
    for _, err := range validationErrors {
        fmt.Printf("Validation Error: Field '%s' failed validation with tag '%s'\n", err.Field(), err.Tag())
    }
    return err
}

// Now, each structure simply calls the validate helper.
func (u *User) Validate() error {
    err := getValidator().Struct(u)
    if err != nil {
        return printValidationErrors(err)
    }
    return nil
}

func (s *CreateUserReq) Validate() error {
    err := getValidator().Struct(s)
    if err != nil {
        return printValidationErrors(err)
    }
    return nil
}

func (l *Login) Validate() error {
    err := getValidator().Struct(l)
    if err != nil {
        return printValidationErrors(err)
    }
    return nil
}