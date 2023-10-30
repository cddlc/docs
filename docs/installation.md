# Installation

## Executable
### go install (≥1.18)

`cddlc` is available as a standalone golang binary through go install.

<!-- ``` sh
go install go.hanneskimara.com/cddlc/cmd/cddlc
``` -->

``` sh 
--8<-- "includes/examples/build_docs.sh:go_cddlc_install"
```

The installation can be tested by running, 

``` sh
--8<-- "includes/examples/build_docs.sh:test_cddlc_install"
```

which will print the help to the terminal.

``` 
--8<-- "includes/examples/extra/cddlc_out.txt"
```

## Go Library

[![GoDoc](https://pkg.go.dev/badge/github.com/HannesKimara/cddlc?utm_source=godoc)](https://pkg.go.dev/github.com/HannesKimara/cddlc)

The `cddlc` module exports go packages used for tokenizing and parsing `CDDL` documents. To get, run:

``` sh
go get -u github.com/HannesKimara/cddlc
```

To parse a sample `CDDL` document to its AST representation 

=== "main.go"

    ``` Go
    package main

    import (
        "fmt"
        "io/ioutil"
        "log"

        "github.com/HannesKimara/cddlc/lexer"
        "github.com/HannesKimara/cddlc/parser"
    )

    func main() {
        src, err := ioutil.ReadFile("foo.cddl")
        if err != nil {
            log.Fatal(err)
        }
        lex := lexer.NewLexer(string(src))
        p := parser.NewParser(lex)
        cddl, err := p.ParseFile()

        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("Found %d rules\n", cddl.Rules)
    }

    ```

=== "foo.cddl"

    ``` cddl
    public-key = [24*24 byte]

    person = {
        name: tstr .size 3
        age: (0..120)
        address: $Address ; Address as a type plug
        public-key: public-key
    }
    ```


