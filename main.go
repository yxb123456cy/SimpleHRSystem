package main

import (
	_ "SimpleHRSystem/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"SimpleHRSystem/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
