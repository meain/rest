# rest

> CLI version of https://github.com/pashky/restclient.el

## Usage

### HEAD

You can specify an http get endpoint

``` rest
HEAD https://postman-echo.com/head
```

### GET

You can specify an http get endpoint

``` rest
GET https://postman-echo.com/get
```

### GET with comments

You can specify an http get endpoint

``` rest
# This is a sample requst
GET https://postman-echo.com/get
```

### GET with params

You can specify an http get endpoint

``` rest
GET https://postman-echo.com/get?foo1=bar1&foo2=bar2
```

### GET with headers

You can specify an http get endpoint

``` rest
GET https://postman-echo.com/get
Sample-Header: Hello-World
```

### POST with data

You can specify an http get endpoint

``` rest
POST https://postman-echo.com/post
Content-Type: application/json

{
  "key": "value"
}
```
