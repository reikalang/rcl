# rcl-go

Go implementation of RCL

- [antlr/go-target](https://github.com/antlr/antlr4/blob/master/doc/go-target.md)

## Develop

- Go 1.9+
- dep

````bash
make dep-install
make test
make install
````

NOTE: the grammar file is in [<project-root>/spec/RCL.g4](../../spec/RCL.g4) and generated using `make gen-parser`