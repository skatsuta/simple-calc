# simple-calc

Let's implement a dead simple command line calculator from scratch!

## Introduction

Your goal is to implement a very simple command line calculator from scratch in Go.
Write an arithmetic expression parser on your own, without using any existing parser libraries.

## Restrictions

* Do not use `go/*` packages such as `go/parser` or `go/scanner` in the standard library.
* Do not use `text/scanner` package in the standard library.
* Do not use any third-party libraries.
* Do not invoke any external processes, for example using `os/exec` package.

## Steps to implement

The recommended way is TDD style:

1. Run `go test` command and make sure all tests are passed (actually no tests are run).
1. Remove `t.Skip()` line within `TestInt` function in `main_test.go`.
1. Run `go test` command and make sure `TestInt` fails.
1. Implement `calc` function in `main.go` so that all the tests that failed above are passed.
1. If all tests have been passed, then remove `t.Skip()` line within `TestAddSub` function in `main_test.go`.
1. Run `go test` command and make sure `TestAddSub` fails.
1. Update `calc` function in `main.go` so that all the tests that failed above are passed.
1. If all tests have been passed, then remove `t.Skip()` line within `TestMulDiv` function in `main_test.go`.
1. In the same way, update `calc` function until all tests in `main_test.go` have been passed after removing all `t.Skip()` lines.

Eventually, your calculator should be able to correctly calculate complex arithmetic expressions like this:

```shell
$ go build
$ ./simple-calc '(8 - 5) * (-2) + (-9 + 8) * (5 - 7) - 12 / (22 - 16) * (-3)'
2
```
