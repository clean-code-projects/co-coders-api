# Co-Coders API

[![Build Status](https://travis-ci.org/clean-code-projects/co-coders-api.svg?branch=master)](https://travis-ci.org/clean-code-projects/co-coders-api)
[![Codecov][codecov-image]][codecov-url]

## contribute

```sh
# test
go test -v ./...

# coverage
go test -coverprofile cover.out ./...
```

## development environment

We will primarily be using VS Code.

here are some extensions you can use to make the sharing the coding a bit easier:

- [Live Share](https://marketplace.visualstudio.com/items?itemName=MS-vsliveshare.vsliveshare)
- [Live Share Chat](https://marketplace.visualstudio.com/items?itemName=karigari.chat)
- [Go](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)
- [Coverage gutters](https://github.com/ryanluker/vscode-coverage-gutters)
- [Run on Save](https://github.com/emeraldwalk/vscode-runonsave.git)

### Workspace settings for code coverage

```json
{
  "folders": [
    {
      "path": "."
    }
  ],
  "settings": {
    "coverage-gutters.coverageReportFileName": "coverage.xml",
    "go.coverOnSave": true,
    "go.coverOnSingleTest": true,
    "go.coverOnSingleTestFile": true,
    "go.testOnSave": true,
    "emeraldwalk.runonsave": {
      "commands": [
        {
          "match": "\\.go$",
          "isAsync": true,
          "cmd": "gocov convert cover.out | gocov-xml > coverage.xml"
        }
      ]
    }
  }
}
```

## godoc Documentation

```sh
godoc -http=localhost:6060
```

Point your browser  to `http://localhost:6060/pkg/github.com/clean-code-projects/co-coders-api/`

## other resources

- [project-layout](https://github.com/golang-standards/project-layout)

[codecov-image]: https://codecov.io/gh/clean-code-projects/co-coders-api/branch/master/graph/badge.svg
[codecov-url]: https://codecov.io/gh/clean-code-projects/co-coders-api
