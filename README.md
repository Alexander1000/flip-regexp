# flip-regexp

tests:
```shell
go test -v -cover .
```

usage:
```shell
package main

import (
    "flip-regexp"
    "fmt"
)

func main() {
    var str string
    fmt.Println("Hello world")

    flip := flip_regexp.NewBuilder("[0-9]{5,9}")

    str, err := flip.Render()

    if err != nil {
        return -1
    }

    fmt.Printf("String: %s", str)
    return 0
}

```
