## Description

project is based on Apalon Backend Coding Challenge.
It implements 4 math operations: `"add", "substract", "divide", "multiply"`

Numeric values are expected: `1,3,4,5`
Also digits support was added: `1.43, 134.12`
All non-numeric input parameters will be rejected with error details

Project use math.big package to calculate large numbers

Example of usage:
`http://localhost/add?x=2&y=5
Output:
{“action”: “add”, “x”: 2, “y”: 5, “answer”, 7, “cached”: false}`

## Configured with

- Go Modules
- Gin
- Built-in **Custom Validators**
- Enviroment support
- And few other important utilties
- Internal cache is used, records are expiring after 1 min

### Installation

Project requires installed go 1.20 version

```
$ go get github.com/vlbaluk/apalon-test
```

```
$ cd $GOPATH/src/github.com/vlbaluk/apalon-test
```

```
$ go mod init
```

```
$ go install
```

## Building Your Application

```
$ go build -v
```
## Start application
- By Default application starts at the 80 port. Configurable in .env file
```
go run .
```