# RCL Specification

The original version is written in ANTLR g4 grammar [RCL.g4](RCL.g4).
However, the hand written go version is the most up to date implementation.

In short RCL is json with comment and less restriction.

- underscore is allowed in number e.g. `100_000`
- extra comma is allowed in array and object e.g. `[1, 2, 3, 4,]` `{"a": 1, "b": 2,}`
- quote can be skipped for string and object key

```rcl
{
    parameters: {
        port: 100_100,
        host: localhost,
    },
    containers: [
        {
            image: mysql,
            name: gugumysql, // just name it after gugu .w.
        }
    ]
}
```