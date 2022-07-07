package cmd

import (
	"context"
	"fmt"
)

func GrpcCmd(ctx context.Context) error {
	fmt.Println("Hello from GRPC command")
	return nil
}
