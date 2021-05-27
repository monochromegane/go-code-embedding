# go-code-embedding [![Actions Status](https://github.com/monochromegane/go-code-embedding/workflows/Go/badge.svg)](https://github.com/monochromegane/go-code-embedding/actions)

A tool to embed Go source code into binary using go:embed.
This is a typical use case for [monochromegane/go-embedding-accessor](https://github.com/monochromegane/go-embedding-accessor).

## Usage

1. Install go-code-embedding.
2. Execute `go-code-embedding` command at your repository.

```sh
$ go-code-embedding --pkg PACKAGE_NAME
```

You can find `code_embedding.go` in your repository.
So, your Go program get {show,list,restore}-code options.

### List code

`list-code` list your code.

```sh
$ my-app --list-code
app.go
my-app/main.go
```

### Show code

`show-code` show your code.

```sh
$ my-app --show-code my-app/main.go
package main

func main() {
    my-app.Do()
}
```

### Restore code

`restore-code` restore your code.

```sh
$ my-app --restore-code
```

You also specify output path by `--restore-path` option.

## Generated code dependency

Generated code (`code_embedding.go`) doesn't depend on your code, only provides {show,list,restore}-code options at init() function.


## Installation

```sh
$ go install github.com/monochromegane/go-code-embedding/cmd/go-code-embedding@latest
```

## Contribution

1. Fork it
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

## License

[MIT](https://github.com/monochromegane/go-code-embedding/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)
