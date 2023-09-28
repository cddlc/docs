/*
  File generated using `cddlc.exe gen`. DO NOT EDIT
*/

package foo

// (cddlc) Ident: Person
type Person struct {
	Name     string
	Age      uint
	Location Location
}

// Valid evaluates type constraints on Person and returns nil if valid 
// else it returns a list of validation errors
func (person *Person) Valid() error {
	return nil
}
