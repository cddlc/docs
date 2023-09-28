/*
  File generated using `cddlc.exe gen`. DO NOT EDIT
*/

package foo

// (cddlc) Ident: null-type
type NullType nil

// Valid evaluates type constraints on null-type and returns nil if valid 
// else it returns a list of validation errors
func (null-type *NullType) Valid() error {
	return nil
}

// (cddlc) Ident: true-alt
var TrueAlt = true

// Valid evaluates type constraints on true-alt and returns nil if valid 
// else it returns a list of validation errors
func (true-alt *TrueAlt) Valid() error {
	return nil
}

// (cddlc) Ident: false-val
var FalseVal = false

// Valid evaluates type constraints on false-val and returns nil if valid 
// else it returns a list of validation errors
func (false-val *FalseVal) Valid() error {
	return nil
}
