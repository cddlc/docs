<!-- ---
title: Home
template: home.html
--- -->

# Overview

CDDL (Concise Data Definition Language) is a domain-specific language for describing the structure of CBOR (Concise Binary Object Representation) and JSON data. It provides a human-readable way to specify data schemas with precise type definitions, constraints, and validation rules.

`cddlc` is an experimental golang code generator for CDDL[^1] based data schemas. It generates Go code for (de)serialization and validation of CBOR data structures.

## Key Features

- **Code Generation**: Automatically generates Go structs and methods from CDDL schemas
- **CBOR Support**: Native support for CBOR serialization and deserialization
- **Type Safety**: Generates strongly-typed Go code with compile-time guarantees
- **Validation**: Built-in validation logic based on CDDL constraints
- **Extensible**: Supports custom types and validation rules

## Use Cases

`cddlc` is particularly useful for:

- **Protocol Implementation**: Building CBOR-based network protocols
- **Data Exchange**: Creating reliable data interchange formats
- **Configuration Files**: Validating structured configuration data
- **IoT Applications**: Efficient data representation for resource-constrained devices
- **API Development**: Type-safe data contracts between services

## How It Works

1. **Define Schema**: Write your data structure definition in CDDL
2. **Generate Code**: Run `cddlc` to generate Go types and methods
3. **Integrate**: Use the generated code in your Go applications
4. **Validate**: Automatically validate data against your schema

## Quick Example

```cddl
person = {
  name: text,
  age: uint,
  email: text,
  ? address: text
}
```

This CDDL definition generates Go structs with methods for marshaling, unmarshaling, and validation.

## Getting Started

Ready to start using `cddlc`? Check out our [Installation Guide](installation.md) to get up and running, then explore the [Language Reference](language/index.md) to learn CDDL syntax and features.

[^1]:
    Concise Data Definition Language (CDDL - [RFC 8610](https://www.rfc-editor.org/rfc/rfc8610.txt)) is a notational convention to express Concise Binary Object Representation (CBOR) and JSON Data Structures.
