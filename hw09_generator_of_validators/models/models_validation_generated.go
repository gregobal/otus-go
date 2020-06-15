// Code generated by go-validate. DO NOT EDIT.
package models

import (
	"fmt"
	"regexp"
)

type ValidationError struct {
	Field string
	Err   error
}

func (a App) Validate() ([]ValidationError, error) {
	vErrs := []ValidationError{}

	if len(a.Version) != 5 {
		vErrs = append(vErrs, ValidationError{
			Field: "Version",
			Err:   fmt.Errorf("wrong length"),
		})
	}

	return vErrs, nil
}

func (r Response) Validate() ([]ValidationError, error) {
	vErrs := []ValidationError{}

	arr := []int{200, 404, 500}
	isErr := true
	for _, i := range arr {
		if r.Code == i {
			isErr = false
			break
		}
	}
	if isErr {
		vErrs = append(vErrs, ValidationError{
			Field: "Code",
			Err:   fmt.Errorf("not allowed value"),
		})
	}

	return vErrs, nil
}

func (u User) Validate() ([]ValidationError, error) {
	vErrs := []ValidationError{}

	if len(u.ID) != 36 {
		vErrs = append(vErrs, ValidationError{
			Field: "ID",
			Err:   fmt.Errorf("wrong length"),
		})
	}

	if u.Age < 18 {
		vErrs = append(vErrs, ValidationError{
			Field: "Age",
			Err:   fmt.Errorf("less than min"),
		})
	}

	if u.Age > 50 {
		vErrs = append(vErrs, ValidationError{
			Field: "Age",
			Err:   fmt.Errorf("more than max"),
		})
	}

	matched, err := regexp.MatchString("^\\w+@\\w+\\.\\w+$", u.Email)
	if err != nil {
		return vErrs, fmt.Errorf("wrong regexp")
	}
	if !matched {
		vErrs = append(vErrs, ValidationError{
			Field: "Email",
			Err:   fmt.Errorf("does not match regexp"),
		})
	}

	for i, _ := range u.Phones {

		if len(u.Phones[i]) != 11 {
			vErrs = append(vErrs, ValidationError{
				Field: "Phones[i]",
				Err:   fmt.Errorf("wrong length"),
			})
		}

	}

	return vErrs, nil
}