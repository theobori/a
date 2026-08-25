# My opinionated tool to manage notes

[![build badge](https://github.com/theobori/fleur/actions/workflows/build.yml/badge.svg)](https://github.com/theobori/fleur/actions/workflows/build.yml)

[![built with nix](https://builtwithnix.org/badge.svg)](https://builtwithnix.org)

The name of the project is "a", pronounced \\eɪ\\. Yes.

This GitHub repository is a [KISS](https://en.wikipedia.org/wiki/KISS_principle) project that contains a CLI wrote in the Go language. This CLI allows me to manage my note databases as files, each with a name that follows a naming convention. The tool supports nested folders, but it was originally designed to operate in folders with a depth of one, like a flat file system.

It can create new notes, edit existing ones, and filter by metadata in the file names.

## Getting started

Before continuing to read the documentation, make sure you have at least [Go](https://go.dev/dl/) version 1.26.5.

### Installation

If you want to use the CLI, you can build and install the binary using the following command.

```bash
go install github.com/theobori/a/cmd/a
```

### Environment variable

You can set a value for the `EDITOR` environment variable, which will then be used for managing the editor program. If unset, its value will be `vi`.

## File name scheme

The project is entirely based on a specific file name scheme.

### Grammar rules

The file name scheme is described by the grammar rules in BNF format below.

```bnf
<filename> ::= <filename-no-ext> | <filename-no-ext> ".file-extension"
<filename-no-ext> ::= (<title> | <title> "__" <tags>)

<title> ::= <any-word>
<tag> ::= <any-word>
<tags> ::= <tag> | <tag> <tags-separator> <tags>

<parts-separator> ::= "__"
<tags-separator> ::= "_"

<any-word> ::= "any-chars-sequence"
```

<!-- comment>
TODO: Create these parameters for the query command
<-->
Note that `<parts-separator>` and `<tags-separator>` can be customized by changing the values of certain command-line parameters.

## Dedicated query language

The project uses a dedicated query language for searching notes. It is capable of compiling expressions into bytecode and then passing them to a stack-based bytecode virtual machine, which evaluates the bytecode and returns whether the expression matches some of the file metadatas.

### Grammar Rules

The dedicated query language is quite flexible in its syntax. It is described by the grammar rules in BNF format below.

```bnf
<expression> ::= <filter> | <literal>
<mutiple-expressions> ::= <expression> | <expression> <mutiple-expressions>

<filter> ::= <begin-expr> (<unary> | <binary> | <function-call>) <end-expr>
<literal> ::= <string>
<string> ::= '"' "any-chars-sequence" '"'

<binary> ::= <and> | <or>
<and> ::= "and" <mutiple-expressions>
<or> ::= "or" <mutiple-expressions>

<unary> ::= <not>
<not> ::= "not" <expression>

<function-call> ::= <title> | <tag> | <ext> | <from-date> | <to-date> | <on-date>
<title> ::= "title" <expression>
<tag> ::= "tag" <expression>
<ext> ::= "ext" <expression>
<from-date> ::= "from-date" <expression>
<to-date> ::= "to-date" <expression>
<on-date> ::= "on-date" <expression>

<begin-expr> ::= "("
<end-expr> ::= ")"
```

### Specifications

This section contains the language specifications.

#### Native Functions

The table below summarizes the features of the native functions.

| Name | Description | Arguments count | Arguments type |
| - | - | - | - |
| **title** | Match the title of a note. | 1 | string |
| **tag** | Match the tag of a note. | 1 | string |
| **ext** | Match the file extension of a note. | 1 | string |
| **from-date** | Returns true if the file has been modified on a date or after. | 1 | date string |
| **to-date** | Returns true if the file has been modified on a date or before. | 1 | date string |
| **on-date** | Returns true if the file has been modified on a date. | 1 | date tring |

#### Operators

The table below summarizes the features of the operator.

| Name | Description | Arguments count | Arguments type |
| - | - | - | - |
| **and** | Returns true if there are 0 arguments. Otherwise, returns true if all arguments are true. | 0 to unlimited | string, boolean |
| **or** | Returns false if there are 0 arguments. Otherwise, returns true if any of the arguments is true. | 0 to unlimited | string, boolean |
| **not** | Returns true is the operand is [falsey](https://developer.mozilla.org/en-US/docs/Glossary/Falsy). Otherwise false. | 0 to unlimited | string, boolean |

### Example

Below is a simple expression.

```elisp
(tag "medium")
```

Below is another simple expression.

```elisp
(or
  (tag "medium")
  (tag "hard"))
```

Below is a complex expression.

```elisp
(and
  (tag "leetcode")
  (or
    (tag "medium")
    (tag "hard"))
  (ext ".md"))
```

Below is another complex expression.

```elisp
(and
  (or 
    (title "title1")
    (title "title2))
  (tag "tag1")
  (from-date "2020-01-01")
  (to-date "2020-02-01")
  (not (title "title1"))
  (and
    (tag "tag2")
    (tag "tag3")))
```

## CLI

The tool's commands are described in this section.

### `help`

This command displays the help message.

### `create`

This command is used to create a new note file in a specified directory. The creation process is interactive; the user is prompted to provide the various pieces of information needed to construct the file name. Below is its help message.

```text
usage: a create [<flags>] [<directory>]

Create a new note file.


Flags:
  --[no-]help          Show context-sensitive help (also try --help-long and --help-man).
  --[no-]disable-open  Disable editor file opening after creating a new file.

Args:
  [<directory>]  Base directory path.
```

### `edit`

This command is used to edit an existing note file by specifying a path. The `EDITOR` environment variable contains the name or path to a text editor that must be able to accept the path of the file to be edited as a positional argument. The tool will create a subprocess with this information. Below is its help message.

```text
usage: a edit <path>

Edit a note file.


Flags:
  --[no-]help  Show context-sensitive help (also try --help-long and --help-man).

Args:
  <path>  File path.
```

### `rename`

This command is used to rename an existing note file by specifying a path. The renaming process is interactive, just like when creating a file. Below is its help message.

```text
usage: a rename [<flags>] <path>

Rename a note file.


Flags:
      --[no-]help      Show context-sensitive help (also try --help-long and --help-man).
  -i, --[no-]insecure  Skip the source filename verification.
medium
Args:
  <path>  File path.
```

### `query`

This command is used to filter note files by specifying an expression in the dedicated query language that will match the metadata in the names of those files. The matched file are passed to [go-fuzzyfinder](github.com/ktr0731/go-fuzzyfinder) that exposed them via an intuitive interactive TUI. Below is its help message.

```text
usage: a query [<flags>] [<directory>]

Filter files in the base directory using a custom query language.


Flags:
      --[no-]help        Show context-sensitive help (also try --help-long and --help-man).
  -q, --query=""  The query to filter file metadatas.
  -r, --[no-]recursive   Recursively walk into the base directory.

Args:
  [<directory>]  Base directory path.
```

## Examples

You can view the example of a note files directory with the [example](./example) folder.

## Contribute

If you'd like to contribute to the project, please follow the instructions provided in the [CONTRIBUTING.md](./CONTRIBUTING.md) file.