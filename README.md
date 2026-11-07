# go-sonyflake

Sonyflake ID generator with injectable machine ID and epoch.

## Install

```bash
go get github.com/andywang2011/go-sonyflake@v0.1.0
```

## Quick start

```go
import (
	"fmt"

	"github.com/andywang2011/go-sonyflake"
)

cfg := sonyflake.DefaultConfig()
cfg.MachineID = 42

gen, err := sonyflake.New(cfg)
if err != nil {
	fmt.Println("init error")
	return
}

id, err := gen.NextID()
if err != nil {
	fmt.Println("next error")
	return
}
fmt.Println(id > 0)
```

## Examples

```bash
cd go-sonyflake/examples && go run .
MACHINE_ID=99 cd go-sonyflake/examples && go run .
```

## License

MIT — see [LICENSE](LICENSE).
