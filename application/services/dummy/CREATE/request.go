package dummy

import (
	"errors"
	domain "go-skeleton/application/domain/dummy"
	"go-skeleton/pkg/validator"
)

type Request struct {
	Dummy     domain.Dummy
	Err       error
	validator *validator.Validator
}

func NewRequest(dummy domain.Dummy, validator *validator.Validator) Request {
	return Request{
		Dummy:     dummy,
		validator: validator,
	}
}

func (r *Request) Validate() error {
	if err := r.dummyCreateRule(); err != nil {
		return err
	}
	errs := r.validator.ValidateStruct(r.Dummy)
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Request) dummyCreateRule() error {

	if r.Dummy.DummyName == "" {
		return errors.New("invalid_argument")
	}
	return nil
}
