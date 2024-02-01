# lxn-go
`lxn-go` is an [lxn](https://github.com/liblxn/lxn) client library for the
Go programming language.

## Translating Text
To translate text, a dictionary has to be loaded:
```golang
func ReadDictionary(r io.Reader) (*Dictionary, error)
```
This function reads a binary dictionary from `r` and returns the dictionary
with all messages and the locale information needed to format these messages.
Once a dictionary is obtained, it can be used to translate messages.

A message can contain variable parts which can be replaced during runtime.
In order to render a message correctly all of the variables need to be
passed into the translation method of the dictionary. The following variable
types are currently supported:

* [`Int`](https://godoc.org/github.com/liblxn/lxn-go#Int) for signed integer values
* [`Uint`](https://godoc.org/github.com/liblxn/lxn-go#Uint) for unsigned integer values
* [`Float`](https://godoc.org/github.com/liblxn/lxn-go#Float) for floating-point numbers
* [`String`](https://godoc.org/github.com/liblxn/lxn-go#String) for strings

## Example
```golang
package main

import (
	"fmt"
	"log"

	lxn "github.com/liblxn/lxn-go"
)

func main() {
	dic := loadDictionary("en")
	msg := dic.Translate("the-section", "the-message-key", lxn.Context{
		"key1": lxn.Float(3.1415),
		"key2": lxn.Uint(7),
	})

	fmt.Println("message for 'hello.lxn':", msg)
}

func loadDictionary(lang string) *lxn.Dictionary {
	dic, err := lxn.FromFile(lxn.ReadDictionary, fmt.Sprintf("translations/%s.lxnc", lang))
	if err != nil {
		log.Fatal(err)
	}
	return dic
}
```
