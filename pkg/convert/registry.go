// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package convert

import (
	"reflect"
	"sync"
)

// Registry is a registry of [FromTo] functions.
type Registry struct {
	m  map[reflect.Type]map[reflect.Type]*wrapper // Source to destination.
	mx sync.RWMutex                               // Guards the above map.
}

// NewRegistry returns a new instance of [Registry].
//
// Use [RegisterConverter] and [LookupConverter] functions to operate on it.
func NewRegistry() *Registry { return &Registry{} }

// register registers the wrapped [FromTo] function. If a wrapper for the
// given [FromTo] already exists, it is returned and the new one replaces it.
func (reg *Registry) register(wrp *wrapper) *wrapper {
	reg.mx.Lock()
	defer reg.mx.Unlock()

	if reg.m == nil {
		reg.m = make(map[reflect.Type]map[reflect.Type]*wrapper, 20)
	}
	if reg.m[wrp.from] == nil {
		reg.m[wrp.from] = make(map[reflect.Type]*wrapper)
	}

	old := reg.m[wrp.from][wrp.to]
	reg.m[wrp.from][wrp.to] = wrp
	return old
}

// lookup returns a [wrapper] registered for the given source-destination type
// pair. Returns nil when a wrapper for a given pair doesn't exist.
func (reg *Registry) lookup(from, to reflect.Type) *wrapper {
	reg.mx.RLock()
	defer reg.mx.RUnlock()

	if reg.m == nil {
		return nil
	}
	if f := reg.m[from]; f != nil {
		return f[to]
	}
	return nil
}

// RegisterConverter registers the [FromTo] in the provided [Registry]. If a
// converter for the same source-destination type pair already exists, it is
// replaced, and the previous converter is returned; otherwise nil is returned.
func RegisterConverter[From, To any](reg *Registry, conv FromTo[From, To]) FromTo[From, To] {
	if conv == nil {
		return nil
	}
	wrp := reg.register(wrap(conv))
	if wrp == nil {
		return nil
	}
	if h, ok := wrp.cnv.(FromTo[From, To]); ok {
		return h
	}
	return nil
}

// LookupConverter returns the [FromTo] for the given source-destination
// type pair from the provided [Registry]. Returns nil if no converter was
// registered for the given source-destination type pair.
func LookupConverter[From, To any](reg *Registry) FromTo[From, To] {
	wrp := reg.lookup(reflect.TypeFor[From](), reflect.TypeFor[To]())
	if wrp == nil {
		return nil
	}
	if h, ok := wrp.cnv.(FromTo[From, To]); ok {
		return h
	}
	return nil
}
