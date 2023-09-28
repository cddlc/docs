# Introduction

This section covers how to use `cddlc` to actually generate code from a definition file. It also explains some of the choices made in interpretation of the language. 

Throughout this section single file builds will be invoked by running `generate` command in the form

``` sh
cddlc generate --source foo.cddl --out lib/foo_gen.go --package foo
```

where `foo.cddl` is the input file, `lib/foo_gen.go` is the generated Go file and `foo` is the name of the package.


!!! tip
    For a full reference of cddl concepts and their level of support in `cddlc` skip to the [references](./reference.md) section.