### amulet.go

> "An amulet is a kind of poem that depends on language, code, and luck." - Robin Sloan
>
> https://www.robinsloan.com/special/amulet/definition/


```bash
go get "github.com/shmup/amulet.go"
```

```go
package main

import (
    "fmt"
    amulet "github.com/shmup/amulet.go"
)

func main() {
    isAmulet, rarity = amulet.IsAmulet("DON'T WORRY.")
    fmt.Println(isAmulet, rarity) // true, 8888
}
```
