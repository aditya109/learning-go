package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningOperation(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Operation completed successfully")
		return nil
	case <-ctx.Done():
		fmt.Println("Operation cancelled")
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := longRunningOperation(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Operation completed without errors")
	}
}
