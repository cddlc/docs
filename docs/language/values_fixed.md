# Fixed Values

The four supported fixed value types are `nil`/`null`, `true`, `false`

From the definition file `fixed.cddl` containing the following content, 

``` cddl title="fixed.cddl" linenums="1"
--8<-- "includes/examples/cddl/fixed.cddl"
```

go code can be generated by invoking

``` sh
--8<-- "includes/examples/build_docs.sh:lang_fixed_build"
```

This generates the declarations and validation methods. Since the types are directly enforced by the Go type checker, the Valid() methods return nil by default. All values are public by default with identifiers formatted as PascalCase.

``` Go title="fixed.go" linenums="1"
--8<-- "includes/examples/lib/fixed.go"
```