package rig

import (
	"flag"
	"strconv"

	"github.com/yazgazan/rig/validators"
)

type uint64Validators struct {
	*uint64Value
	validators []validators.Uint64
}

func (v uint64Validators) Set(s string) error {
	err := v.uint64Value.Set(s)
	if err != nil {
		return err
	}

	for _, validator := range v.validators {
		err = validator(uint64(*v.uint64Value))
		if err != nil {
			return err
		}
	}

	return nil
}

func (v uint64Validators) New(i interface{}) flag.Value {
	return uint64Validators{
		uint64Value: (*uint64Value)(i.(*uint64)),
		validators:  v.validators,
	}
}

func (v uint64Validators) IsNil() bool {
	return v.uint64Value == nil
}

type uint64Value uint64

func (i uint64Value) String() string {
	return strconv.FormatUint(uint64(i), 10)
}

func (i *uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uint64Value(v)
	return err
}

// Uint64 creates a flag for a uint64 variable.
func Uint64(v *uint64, flag, env, usage string, validators ...validators.Uint64) *Flag {
	return &Flag{
		Value: uint64Validators{
			uint64Value: (*uint64Value)(v),
			validators:  validators,
		},
		Name:     flag,
		Env:      env,
		Usage:    usage,
		TypeHint: "uint64",
	}
}

// Uint64Generator is the default uint64 generator, to be used with Repeatable for uint64 slices.
func Uint64Generator() Generator {
	return func() flag.Value {
		return new(uint64Value)
	}
}
