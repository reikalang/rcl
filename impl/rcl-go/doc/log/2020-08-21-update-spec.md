# 2020-08-21 Update Spec

## Background

While working on dyweb/pm, I want to use RCL instead of JSON to write document.
Though in the end I still chose JSON to boostrap the project, there are some notes I feel should be reflected in the spec.

- Full compatibility with JSON, a valid JSON is always a valid RCL, the opposite is not true
- Less quote, similar to YAML, we can skip quote if a string (object key or value) meets the following:
  - one line
  - no space
  - no quote
- Only double quote
- Decode into a struct with position information, user only need to write normal go struct and use `// rcl:genpos`

```go
// rcl:genpos
type Server struct {
    Port int
    Addr string
    Debug bool
}

// generated
type ServerDecoded struct {
    Port rcl.Int
    Addr rcl.String
    Debug rcl.Bool
}

// defined in rcl, this can even be Ast in RCL ....
type Int {
    Val int
    Pos rcl.Position
}

type Position {
    File string
    Start Axis
    End Axis
}

type Axis struct {
    Line int
    Col int
}
```