# Hanaboso GO log

**Download**
```
go mod download github.com/hanaboso/go-log
```

**Usage**
```
import (
    "errors"    
    "github.com/hanaboso/go-log/pkg/zap"
    logLevel "github.com/hanaboso/go-log/pkg"
)

logger := zap.NewLogger()
logger.SetLevel(logLevel.ERROR)

logger.WithFields(map[string]interface{}{
		"string": "value",
		"int": 123,
	}).Info("Message with data of %s", "fields")

logger.Fatal(errors.New("this will panic"))
```
