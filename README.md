# gomock

gomock is command line tool which makes simple API mock server

# Install

```
go get github.com/hlts2/gomock
```

## Example

Create a `config.yml` file

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

Create a response json file (ex `todos.json`)

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

Start json mock server

```
gomock run -s config.yml
```

Mocked GET /api/v1/todos:

```
$ curl -v GET localhost:8080/api/v1/todos

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

## CLI Usage

```
$ gomock run --help
Usage:
  gomock run [flags]

Flags:
  -h, --help         help for run
  -s, --set string   set config file (default "config.yml")
```
