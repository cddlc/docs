/*
  File generated using `cddlc.exe gen`. DO NOT EDIT
*/

package

// (cddlc) Ident: name
type Name string

// Valid evaluates type constraints on name and returns nil if valid 
// else it returns a list of validation errors
func (name *Name) Valid() error {
	return nil
}

// (cddlc) Ident: age
type Age uint

// Valid evaluates type constraints on age and returns nil if valid 
// else it returns a list of validation errors
func (age *Age) Valid() error {
	return nil
}

// (cddlc) Ident: public-key
type PublicKey []byte

// Valid evaluates type constraints on public-key and returns nil if valid 
// else it returns a list of validation errors
func (public-key *PublicKey) Valid() error {
	return nil
}

// (cddlc) Ident: account-balance
type AccountBalance int

// Valid evaluates type constraints on account-balance and returns nil if valid 
// else it returns a list of validation errors
func (account-balance *AccountBalance) Valid() error {
	return nil
}
