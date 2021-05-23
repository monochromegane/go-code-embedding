package main

import (
	"flag"
	"fmt"
	"os"

	embedding "github.com/monochromegane/go-code-embedding"
)

var pkg string
var outputDir string

func init() {
	flag.StringVar(&pkg, "pkg", "main", "Package name to use in the generated code.")
	flag.StringVar(&outputDir, "output-dir", ".", "Output directory")
	flag.Parse()
}

func main() {
	err := embedding.Generate(pkg, outputDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
