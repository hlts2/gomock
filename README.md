# gomock [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/hlts2/gomock)](https://goreportcard.com/report/github.com/hlts2/gomock) [![Join the chat at https://gitter.im/hlts2/gomock](https://badges.gitter.im/hlts2/gomock.svg)](https://gitter.im/hlts2/gomock?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

gomock is command line tool which makes simple API mock server. No more waiting on backend teams to deliver services.

## Install

```shell
go get github.com/hlts2/gomock
```

## Example

### Create a config file
`config.yml` will help you get started mocking your API's.

```yaml
port: 1234
endpoints:
    - request:
        path: /api/v1/todos/
        method: GET //uppercase
      response:
        code: 200
        body: todos.json
        headers:
            content-type: application/json
```

### Request path
Paths can have variables. They are defined using the format {:id} or {:name}.

```yaml
path: /api/v1/users/{:id}/

path: /api/v1/users/{:id}/name/

path: /api/v1/users/{:id}/bools/{:book_id}/

path: /api/v1/users?id={:id}
```

### Create JSON response file

```json
{
  "todos": [
    {
      "id": 1,
      "title": "hoge"
    },
    {
      "id": 2,
      "title": "foo"
    }
  ]
}

```

### Start API mock server

```
gomock run -s config.yml
```

### Send request

```
$ curl -v GET localhost:8080/api/v1/todos/

> GET /api/v1/todos HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.51.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Wed, 20 Sep 2017 23:56:27 GMT
< Content-Length: 119
<
{
  "todos": [
    {
      "id": 1,
      "title": "hoge"
    },
    {
      "id": 2,
      "title": "foo"
    }
  ]
}

```

## Usage

```
$ gomock run --help
Usage:
  gomock run [flags]

Flags:
  -h, --help         help for run
  -s, --set string   set config file (default "config.yml")
```
