#! /bin/bash

set -euxo pipefail

# Install using go install
# --8<-- [start:go_cddlc_install]
go install github.com/HannesKimara/cddlc/cmd/cddlc@latest
# --8<-- [end:go_cddlc_install]

# Test installation and log output
{
# --8<-- [start:test_cddlc_install]
cddlc --help
# --8<-- [end:test_cddlc_install]
} | tee ./extra/cddlc_out.txt

# Build language docs
{
# --8<-- [start:lang_primitives_build]
cddlc generate --source primitives.cddl --out lib/primitives.go --package foo
# --8<-- [end:lang_primitives_build]

# --8<-- [start:lang_fixed_build]
cddlc generate --source fixed.cddl --out lib/fixed.go --package foo
# --8<-- [end:lang_fixed_build]

# --8<-- [start:lang_socket_build]
cddlc generate --source plug_socket.cddl --out lib/plug_socket.go --package foo
# --8<-- [end:lang_socket_build]
}


