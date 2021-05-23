package embedding

import "path/filepath"

func Generate(pkg, outputDir string) error {
	return writeToFile(
		filepath.Join(outputDir, "code_embedding.go"),
		codeTemplates.toString(),
		data{Package: pkg},
	)
}
