package manifold

import (
	"github.com/asaskevich/govalidator"
	"regexp"

	"github.com/manifoldco/go-manifold/errors"
)

var (
	labelRegex = regexp.MustCompile("^[a-z0-9][a-z0-9-_]{1,128}$")
	nameRegex  = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9 -]{2,128}$")
)

var (
	errInvalidLabel = NewError(errors.BadRequestError, "Invalid label")
	errInvalidName  = NewError(errors.BadRequestError, "Invalid name")
	errInvalidEmail = NewError(errors.BadRequestError, "Invalid Email Provided")
)

// Label represents any object's label field
type Label string

// Validate ensures the label value is valid
func (lbl Label) Validate(_ interface{}) error {
	if labelRegex.Match([]byte(lbl)) {
		return nil
	}

	return errInvalidLabel
}

// Name represents any object's name field
type Name string

// Validate ensures the name value is valid
func (n Name) Validate(_ interface{}) error {
	if nameRegex.Match([]byte(n)) {
		return nil
	}

	return errInvalidName
}

// Email represents any email field
type Email string

// Validate ensures that the email is valid
func (e Email) Validate(_ interface{}) error {
	if govalidator.IsEmail(string(e)) {
		return nil
	}

	return errInvalidEmail
}