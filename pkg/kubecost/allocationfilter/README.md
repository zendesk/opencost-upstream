# Allocation Filters

This package includes the full grammar for the Allocation filter language in
ANTLR4 format ([allocationfilter.g4](./allocationfilter.g4)); the generated Go
code for lexing and parsing a filter; and tests.
    
## Generating Go code

Set up ANTLR4 + Go target: https://github.com/antlr/antlr4/blob/master/doc/go-target.md

``` sh
antlr4 -Dlanguage=Go allocationfilter.g4
```

## Testing

There is a Go-native test in `_test.go`. Make sure to run the Go generation first.

For testing just the grammar, do the following:

``` sh
antlr4 allocationfilter.g4 && javac allocationfilter*.java
```

Then you can use `grun`:

``` sh
grun allocationfilter filter inputX.txt -tokens
grun allocationfilter filter inputX.txt -tree
grun allocationfilter filter inputX.txt -gui
```
