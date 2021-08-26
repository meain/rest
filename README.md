# rest

> CLI version of https://github.com/pashky/restclient.el

## Install

### macOS

```shell
brew install meain/homebrew-meain/gloc
```

### Manual

Download the binary from the [release page](https://github.com/meain/rest/releases).

## Usage

You can pass the thing to run via stdin or pass a filename as the first argument.

```shell
echo 'GET https://meain.io' | rest
```

```shell
rest file-with-content.http
```

## Options

#### GET

You can specify an http get endpoint

```rest
GET https://postman-echo.com/get
```

_You can do any other http operation the same way_

#### GET with comments

You can specify an http get endpoint

```rest
# This is a sample requst
GET https://postman-echo.com/get
```

#### GET with params

You can specify an http get endpoint

```rest
GET https://postman-echo.com/get?foo1=bar1&foo2=bar2
```

#### GET with headers

You can specify an http get endpoint

```rest
GET https://postman-echo.com/get
Sample-Header: Hello-World
```

#### POST with data

You can specify an http get endpoint

```rest
POST https://postman-echo.com/post
Content-Type: application/json

{
  "key": "value"
}
```

## Alternatives

- [restcli/restcli](https://github.com/restcli/restcli)
