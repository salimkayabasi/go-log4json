# GO - Log4JSON

It's a primitive library for logging in JSON format with your `go` project.


Sample:

```go
package main

import (
	"net/http"
	"fmt"
	"os"
	Log4json "github.com/salimkayabasi/go-log4json"
)

var logger *Log4json.Logger

func init() {
	logger = Log4json.NewLogger()
}

func main()  {
    logger.Debug("done")
}
```

Output:

```
{"ts":"2018-06-26T09:55:40.742009133Z","level":"debug","msg":"done"}
```