package embedding

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestListCode(t *testing.T) {
	buf := &bytes.Buffer{}
	embed := &embedding{writer: buf}
	err := embed.list()
	if err != nil {
		t.Errorf("embedding.list should not return error, but %v", err)
	}

	actual := buf.String()
	expects := []string{
		"cmd/go-code-embedding/main.go",
		"code_embedding_test.go",
		"embedding.go",
		"embedding_test.go",
		"go.mod",
		"go.sum",
	}

	for _, e := range expects {
		if !strings.Contains(actual, e) {
			t.Errorf("code list should contain %s, but not contains", e)
		}
	}
}

func TestShowCode(t *testing.T) {
	buf := &bytes.Buffer{}
	embed := &embedding{writer: buf}
	embed.list()

	list := strings.Split(buf.String(), "\n")

	for _, c := range list {
		if c == "" {
			continue
		}
		expect := &bytes.Buffer{}
		code := &embedding{writer: expect}
		code.show(c)

		actual, err := ioutil.ReadFile(c)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		if expect.String() != string(actual)+"\n" {
			t.Errorf("code should equal %s but not equal", c)
		}
	}
}
