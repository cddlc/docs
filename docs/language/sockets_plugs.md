# Plugs/Socket

Names that start with a single `$` are `type sockets`, starting out as an empty type, and intended to be extended via `/=`.  Names that start with a double `$$` are `group sockets`, starting out as an empty group choice, and intended to be extended via `//=`.  In either case, it is not an error if there is no definition for a socket at all; this then means there is no way to satisfy the rule (i.e., the choice is empty). [^1]


Below is an example of a group socket - `$$Location`

``` title="plug_socket.cddl" linenums="1"
--8<-- "includes/examples/cddl/plug_socket.cddl"
```

`cddlc` interpretes undefined plugs/sockets as Golang types that are defined in the same package as the generated file. In this case plugs/sockets are not differentiated and the generated output expects the plug/socket to be defined in the same package.

Generating code from this using

``` sh
--8<-- "includes/examples/build_docs.sh:lang_socket_build"
```

exports the Go declarations in "lib/plug_socket.go". To satisfy the Location type in Go a file with this declaration must be added. 


=== "lib/plug_socket.go"
    ``` Go 
    --8<-- "includes/examples/gen/plug_socket.go"
    ```

=== "lib/prelude.go"
    ``` Go
    --8<-- "includes/examples/gen/prelude.go"
    ```

[^1]: 
    https://www.rfc-editor.org/rfc/inline-errata/rfc8610.html