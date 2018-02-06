# lxn-go
`lxn-go` is an [lxn](https://github.com/liblxn/lxn) client library for the Go programming language.

## Translating Text
To translate text, a catalog has to be loaded:
```golang
func ReadCatalog(r io.Reader) (Translator, error)
```
This function reads a binary catalog from `r` and returns its translator function. Once a translator function of type
```golang
type Translator func(key string, ctx Context) string
```
is obtained, it can be used to convert a key and a context into a message. The key specifies the message key within the catalog preceded by its section, i.e. `section.message-key` (or simply `message-key` if the message has no section). The context contains all the necessary variables to render this message correctly. For each variable in the catalog there has to be a value of the corresponding type. The following variable types are currently supported:
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
    "os"

    lxn "github.com/liblxn/lxn-go"
)

func main() {
    tr := translator("en")
    msg := tr("hello.lxn", lxn.Context{
        "key1": lxn.Float(3.1415),
        "key2": lxn.Uint(7),
    })
    
    fmt.Println("message for 'hello.lxn':", msg)
}

func translator(lang string) lxn.Translator {
    f, err := os.Open(fmt.Sprintf("catalog-%s.lxnc", lang))
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    tr, err := lxn.ReadCatalog(f)
    if err != nil {
        log.Fatal(err)
    }
    return tr
}
```
