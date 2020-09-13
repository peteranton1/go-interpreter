// Package vm vm/frame.go
package vm

import (
	"monkey/code"
	"monkey/object"
)

// Frame struct
type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int
}

// NewFrame func
func NewFrame(cl *object.Closure, basePointer int) *Frame {
	f := &Frame{
		cl:          cl,
		ip:          -1,
		basePointer: basePointer,
	}
	return f
}

// Instructions func
func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
