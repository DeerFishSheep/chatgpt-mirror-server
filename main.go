package main

import (
	_ "chatgpt-mirror-server/internal/packed"

	_ "github.com/cool-team-official/cool-admin-go/contrib/drivers/sqlite"

	_ "chatgpt-mirror-server/modules"

	"github.com/gogf/gf/v2/os/gctx"

	"chatgpt-mirror-server/internal/cmd"
)

func main() {
	// gres.Dump()
	cmd.Main.Run(gctx.New())
}