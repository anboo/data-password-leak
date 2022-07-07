package cmd

import (
	"context"
	"fmt"
)

type ImportCmd struct {
}

func (c ImportCmd) Run(ctx context.Context) error {
	fmt.Println("Hello from import command")
	return nil
}

func (c ImportCmd) Name() string {
	return "import"
}
