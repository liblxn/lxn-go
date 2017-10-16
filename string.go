package lxn

// String is a string variable which can be passed to message replacements.
type String string

// String implements the Variable interface.
func (s String) String() string {
	return string(s)
}
