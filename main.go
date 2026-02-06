package main

import (
	"context"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n: args: %v", err, os.Args)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, getenv func(string) string) error {
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

	if typeName == "" {
		return fmt.Errorf("--type must be specified")
	}
	fileName := getenv("GOFILE")
	if fileName == "" {
		return fmt.Errorf("GOFILE is blank")
	}

	fmt.Println("type: ", typeName)
	fmt.Println("file: ", fileName)

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, fileName, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse dir: %w", err)
	}

	ast.Print(fset, file)

	ast.Inspect(file, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if !ok || ts.Name.Name != typeName {
			return true
		}

		fmt.Printf("type spec: %+v\n", ts)
		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return true
		}

		fmt.Printf("struct type: %+v\n", st)

		for _, field := range st.Fields.List {
			for _, name := range field.Names {
				fmt.Printf("method: %s\n", name.Name)
			}
		}
		return false
	})

	return nil
}
