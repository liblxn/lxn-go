package lxn

// Variable defines an interface for variable types which were
// passed to the replacements to replace the dynamic part
// in a message.
type Variable interface {
	String() string
}

// Context holds all the variables which are passed to the replacements.
type Context map[string]Variable // key => replacement variable
