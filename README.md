# KSortMap - a package for sorting the map by keys

## Usage

```go
package main

import (
	"fmt"
	"github.com/KoNekoD/KSortMap-go"
)

func main() {
	fmt.Println(KSortMap.Sort(map[string]int{"CCC": 1, "BBB": 2, "AAA": 3}))

	fmt.Println(KSortMap.Sort(map[int]string{3: "AAA", 1: "BBB", 2: "CCC"}))
}
```