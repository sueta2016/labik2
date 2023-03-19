# Continuous integration and testing automation

This is lab 2 for software architecture subject made by pogromisty

## Download

```bash
git clone https://github.com/sueta2016/labik2
```
## Run

```bash
go run cmd/example/main.go
```
Specify input by flags:

- -e "+ 5 5"
- -f file_path.txt

Specify output by:

- -o file_path.txt

## Test
```bash
go test
```

## GitHub Actions

[Every commit and pull request is checked using GitHub Actions](https://github.com/sueta2016/labik2/actions)

Examples:

- [Check is passed](https://github.com/sueta2016/labik2/commit/04c0906d2d438111cff4c4134e98290624eadbe0)
- [Check is not passed](https://github.com/sueta2016/labik2/commit/83684e4d3ad439f31918376ba8b86bfbe01c2ad6)

[Pull Requests](https://github.com/sueta2016/labik2/pulls?q=is%3Apr+is%3Aclosed)
