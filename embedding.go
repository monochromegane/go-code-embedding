package embedding

import accessor "github.com/monochromegane/go-embedding-accessor"

func Generate(pkg, outputDir string) error {
	return accessor.Generate(
		pkg,
		outputDir,
		"code",
		generateFiles(),
	)
}

func generateFiles() string {
	return "*.go */**/*.go go.mod go.sum"
}
