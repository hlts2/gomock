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
path: /api/v1/users/*/          // match: /api/v1/users/little/, /api/v1/users/tiny/, etc

path: /api/v1/users/*/name/     // match: /api/v1/users/g444/name/, /api/v1/users/f5444/name/, etc

path: /api/v1/users/*/bools/*/  // match: /api/v1/users/111/books/f343/, /api/v1/users/4444/books/d343/, etc

path: /api/v1/users?id=*        // match: /api/v1/users?id=1111, /api/v1/users?id=2222, etc
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
$ curl -v localhost:1234/api/v1/todos/

> GET /api/v1/todos/ HTTP/1.1
> Host: localhost:1234
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
NAME:
   gomock run - start API mock server

USAGE:
   gomock run [command options] [arguments...]

OPTIONS:
   --set value, -s value  config file (default: "config.yml")
   --tls-path value       directory to the TLS server.crt/server.key file
```
