package embedding

import (
	"bufio"
	"bytes"
	"io"
	"os"
	tmpl "text/template"

	"golang.org/x/tools/imports"
)

func writeToFile(file, template string, data data) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	b, err := writeWithFormat(file, template, data)
	if err != nil {
		return err
	}

	w.Write(b)
	return nil
}

func writeWithFormat(file, template string, data data) ([]byte, error) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	write(w, template, data)
	w.Flush()

	formatted, err := imports.Process(file, b.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	return formatted, nil
}

func write(w io.Writer, tplText string, data data) error {
	tpl := tmpl.Must(tmpl.New("t").Parse(tplText))
	if err := tpl.Execute(w, data); err != nil {
		return err
	}

	return nil
}
