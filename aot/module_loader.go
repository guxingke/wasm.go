package aot

import (
	"errors"
	"plugin"

	"github.com/zxh0/wasm.go/instance"
)

type NewFn = func(instance.Map) (instance.Instance, error)

// load compiled module
func Load(filename string) (instance.Instance, error) {
	p, err := plugin.Open(filename)
	if err != nil {
		return nil, err
	}

	f, err := p.Lookup("Instantiate")
	if err != nil {
		return nil, err
	}

	newFn, ok := f.(NewFn)
	if !ok {
		msg := "'Instantiate' is not 'func(instance.Map) instance.Instance'"
		return nil, errors.New(msg)
	}

	return newFn(nil) // TODO
}
