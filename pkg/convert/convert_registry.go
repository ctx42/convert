// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package convert

import (
	"reflect"
	"sync"
)

// Registry is a registry of [SrcToDst] converter functions.
type Registry struct {
	m  map[reflect.Type]map[reflect.Type]*wrapper // Source to destination.
	mx sync.RWMutex                               // Guards the above map.
}

// NewRegistry returns a new instance of [Registry].
//
// Use [RegisterConverter] and [LookupConverter] functions to operate on it.
func NewRegistry() *Registry { return &Registry{} }

// register registers the wrapped [SrcToDst] function. If a wrapper for the
// same source-destination type pair already exists, it is returned and the
// new one replaces it.
// Does nothing and returns nil for a nil registry.
func (reg *Registry) register(wrp *wrapper) *wrapper {
	if reg == nil {
		return nil
	}
	reg.mx.Lock()
	defer reg.mx.Unlock()

	if reg.m == nil {
		reg.m = make(map[reflect.Type]map[reflect.Type]*wrapper, 20)
	}
	if reg.m[wrp.src] == nil {
		reg.m[wrp.src] = make(map[reflect.Type]*wrapper)
	}

	old := reg.m[wrp.src][wrp.dst]
	reg.m[wrp.src][wrp.dst] = wrp
	return old
}

// lookup returns a [wrapper] registered for the given source-destination type
// pair. Returns nil when a wrapper for a given pair doesn't exist or the
// registry is nil.
func (reg *Registry) lookup(from, to reflect.Type) *wrapper {
	if reg == nil {
		return nil
	}
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

// RegisterConverter registers the [SrcToDst] in the provided [Registry]. If a
// converter for the same source-destination type pair already exists, it is
// replaced, and the previous converter is returned; otherwise nil is returned.
// Registers nothing and returns nil when the registry is nil.
func RegisterConverter[Src, Dst any](reg *Registry, conv SrcToDst[Src, Dst]) SrcToDst[Src, Dst] {
	if conv == nil {
		return nil
	}
	wrp := reg.register(wrap(conv))
	if wrp == nil {
		return nil
	}
	if h, ok := wrp.cnv.(SrcToDst[Src, Dst]); ok {
		return h
	}
	return nil
}

// LookupConverter returns the [SrcToDst] for the given source-destination
// type pair from the provided [Registry]. Returns nil if no converter was
// registered for the given source-destination type pair, or the registry is
// nil.
func LookupConverter[From, To any](reg *Registry) SrcToDst[From, To] {
	wrp := reg.lookup(reflect.TypeFor[From](), reflect.TypeFor[To]())
	if wrp == nil {
		return nil
	}
	if h, ok := wrp.cnv.(SrcToDst[From, To]); ok {
		return h
	}
	return nil
}
