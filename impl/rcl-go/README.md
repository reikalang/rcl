# rcl-go

Go implementation of RCL.

## TODO

Error

- [ ] keep track of position in error
- [ ] use standard error struct without `errors.Wrap`
  - learn from the failure crate and the `impl From<otherError> for MyError`

Parser

- [ ] check if escape is working properly, as long as it has the same behavior as encoding/json I should be good
- [ ] decode int
- [ ] decode float
- [ ] support comment
- [ ] preserve comment

Decoder

- [ ] visitor pattern
- [ ] decode into go struct

Encoder

- [ ] encode go struct to rcl

Done

- [x] switch to hand written version (which is still WIP)

## Develop

- Go 1.14

````bash
make test
make install
````

NOTE: the grammar file is in [<project-root>/spec/RCL.g4](../../spec/RCL.g4) and generated using `make gen-parser`

## History

- First version is using ANTLR [antlr/go-target](https://github.com/antlr/antlr4/blob/master/doc/go-target.md), but I switched to hand written parser in [2020-01-14](https://github.com/reikalang/rcl/commit/5ff778993d4201cccbc6457e49c58fe63cec1c42)
