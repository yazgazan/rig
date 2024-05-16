package rig

import (
	"flag"
	"strconv"

	"github.com/yazgazan/rig/validators"
)

type uintValidators struct {
	*uintValue
	validators []validators.Uint
}

func (v uintValidators) Set(s string) error {
	err := v.uintValue.Set(s)
	if err != nil {
		return err
	}

	for _, validator := range v.validators {
		err = validator(uint(*v.uintValue))
		if err != nil {
			return err
		}
	}

	return nil
}

func (v uintValidators) New(i interface{}) flag.Value {
	return uintValidators{
		uintValue:  (*uintValue)(i.(*uint)),
		validators: v.validators,
	}
}

func (v uintValidators) IsNil() bool {
	return v.uintValue == nil
}

type uintValue uint

func (i uintValue) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i *uintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, strconv.IntSize)
	*i = uintValue(v)
	return err
}

// Uint creates a flag for a uint variable.
func Uint(v *uint, flag, env, usage string, validators ...validators.Uint) *Flag {
	return &Flag{
		Value: uintValidators{
			uintValue:  (*uintValue)(v),
			validators: validators,
		},
		Name:     flag,
		Env:      env,
		Usage:    usage,
		TypeHint: "uint",
	}
}

// UintGenerator is the default uint generator, to be used with Repeatable for uint slices.
func UintGenerator() Generator {
	return func() flag.Value {
		return new(uintValue)
	}
}
