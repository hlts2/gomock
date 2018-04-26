# gomock

gomock is command line tool which makes simple API mock server. No more waiting on backend teams to deliver services.

# Install

```
go get github.com/hlts2/gomock
```

## Example

### Create a config file
`config.yml` will help you get started mocking your API's.

```
port: 1234
endpoints:
    - request:
        path: /api/v1/todos/
        method: GET
      response:
        code: 200
        body: todos.json # or `{"todos": [{"id": 1, "title": "hoge"}, {"id": 2, "title": "foo"}]}`
        headers:
            content-type: application/json
```

#### Request path
Paths can have variables. They are defined using the format {:id} or {:name}. 

```
path: /api/v1/user/id={:id}/

path: /api/v1/user/id={:id}/name={:name}

path: /api/v1/user?id={:id}
```

### Create JSON response file

```
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

## Start API mock server

```
gomock run -s config.yml
```

## Send request

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
