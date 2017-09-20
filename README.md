# gomock

gomock is command line tool which makes simple API mock server

# Install

```
go get github.com/hlts2/gomock
```

## Example

Create a `config.yml` file

```
endpoints:
    - path: /api/v1/todos
      method: GET
      response_file: todos.json
```

Create a your json file (`todos.json`)

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

Start JSON server

```
gomock run --s config.yml
```

## CLI Usage

```
$ gomock --help
Usage:
  gomock run [flags]

Flags:
  -h, --help         help for run
      --s string     set config file (default "config.yml")
      --set string   set config file (default "config.yml")
```
