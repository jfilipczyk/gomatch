# Gomatch

<img align="right" width="147px" src="https://raw.github.com/jfilipczyk/gomatch/master/logo.png">

[![Build Status](https://travis-ci.com/jfilipczyk/gomatch.svg?branch=master)](https://travis-ci.com/jfilipczyk/gomatch)
[![codecov](https://codecov.io/gh/jfilipczyk/gomatch/branch/master/graph/badge.svg)](https://codecov.io/gh/jfilipczyk/gomatch)
[![GoDoc](https://godoc.org/github.com/jfilipczyk/gomatch?status.svg)](https://godoc.org/github.com/jfilipczyk/gomatch)
[![Go Report Card](https://goreportcard.com/badge/github.com/jfilipczyk/gomatch)](https://goreportcard.com/report/github.com/jfilipczyk/gomatch)

Library created for testing JSON against patterns. The goal was to be able to validate JSON focusing only on parts essential in given test case so tests are more expressive and less fragile. It can be used with both unit tests and functional tests.

When used with Gherkin driven BDD tests it makes scenarios more compact and readable. See [Gherkin example](#gherkin-example)

## Contests

  - [Installation](#installation)
  - [Basic usage](#basic-usage)
  - [Available patterns](#available-patterns)
  - [Gherkin example](#gherkin-example)
  - [License](#license)
  - [Credits](#credits)

## Installation

```shell
go get github.com/jfilipczyk/gomatch
```

## Basic usage

```go

actual := `
{
  "id": 351,
  "name": "John Smith",
  "address": {
    "city": "Boston"
  }
}
`
expected := `
{
  "id": "@number@",
  "name": "John Smith",
  "address": {
    "city": "@string@"
  }
}
`

m := gomatch.NewDefaultJSONMatcher()
ok, err := m.Match(expected, actual)
if ok {
  fmt.Printf("actual JSON matches expected JSON")
} else {
  fmt.Printf("actual JSON does not match expected JSON: %s", err.Error())
}

```

## Available patterns

* `@string@`
* `@number@`
* `@bool@`
* `@array@`
* `@uuid@`
* `@wildcard@`
* `@...@` - unbounded array or object

### Unbounded pattern

It can be used at the end of an array to allow any extra array elements:
```json
[
  "John Smith",
  "Joe Doe",
  "@...@"
]
```

It can be used at the end of an object to allow any extra keys:
```json
{
  "id": 351,
  "name": "John Smith",
  "@...@": ""
}
```

## Gherkin example

Gomatch was created to use it together with tools like [GODOG](https://github.com/DATA-DOG/godog).
The goal was to be able to validate JSON response focusing only on parts essential in given scenario.

```gherkin
Feature: User management API
  In order to provide GUI for user management 
  As a frontent developer
  I need to be able to create, retrive, update and delete users

  Scenario: Get list of users sorted by username ascending
    Given the database contains users:
    | Username   | Email                  |
    | john.smith | john.smith@example.com |
    | alvin34    | alvin34@example.com    |
    | mike1990   | mike.jones@example.com |
    When I send "GET" request to "/v1/users?sortBy=username&sortDir=asc"
    Then the response code should be 200
    And the response body should match json:
    """
    {
      "items": [
        {
          "username": "alvin34",
          "@...@": ""
        },
        {
          "username": "john.smith",
          "@...@": ""
        },
        {
          "username": "mike1990",
          "@...@": ""
        }
      ],
      "@...@": ""
    }
    """
```

## License

This library is distributed under the MIT license. Please see the LICENSE file.

## Credits

This library was inspired by [PHP Matcher](https://github.com/coduo/php-matcher)

### Logo
The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/).
Gomatch logo was based on a gopher created by Takuya Ueda (https://twitter.com/tenntenn). Licensed under the [Creative Commons 3.0 Attributions license](http://creativecommons.org/licenses/by/3.0/deed.en). Gopher eyes were changed.