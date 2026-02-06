package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n: args: %v", err, os.Args)
		os.Exit(1)
	}
}

func run(ctx context.Context, w io.Writer, args []string) error {
	_, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	fmt.Println(args)
	var typeName string

	flagset := flag.NewFlagSet("run", flag.ExitOnError)

	flagset.StringVar(&typeName, "type", "", "struct type name gobuildergen generating builder for")

	err := flagset.Parse(args[1:])
	if err != nil {
		return fmt.Errorf("flagset failed to parse args: %w", err)
	}

	fmt.Println("type: ", typeName)

	return nil
}
