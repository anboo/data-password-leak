package main

import (
	"context"
	"dataPasswordLeak/cmd"
)

func main() {
	ctx := context.Background()

	app := cmd.NewApp(ctx)
	app.RegisterCommand(cmd.ImportCmd{})
	app.RegisterCommand(cmd.NewCmdFunc("start", cmd.GrpcCmd))
	app.Execute()
}
