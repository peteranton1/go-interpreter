// Package vm vm/frame.go
package vm

import (
	"monkey/code"
	"monkey/object"
)

// Frame struct
type Frame struct {
	fn *object.CompiledFunction
	ip int
}

// NewFrame func
func NewFrame(fn *object.CompiledFunction) *Frame {
	return &Frame{fn: fn, ip: -1}
}

// Instructions func
func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}
