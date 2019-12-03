# httpreused

## Usage

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/skaji/go-httpreused"
)

func main() {
	c := httpreused.Wrap(http.DefaultClient)

	for i := 0; i < 2; i++ {
		res, err := c.Get("https://www.google.com")
		if err != nil {
			panic(err)
		}
		reused := res.Header.Get("X-Connection-Reused")
		ip := res.Header.Get("X-Connection-IP")
		fmt.Println(i, ip, reused)
	}
}
```
```
❯ go run main.go
0 172.217.26.36 false
1 172.217.26.36 true
```
