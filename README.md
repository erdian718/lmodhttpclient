# lmodmsgpack

`http.Client` bindings for [Lua](https://github.com/ofunc/lua).

## Usage

```go
package main

import (
	"github.com/ofunc/lmodhttpclient"
	"github.com/ofunc/lua/util"
)

func main() {
	l := util.NewState()
	l.Preload("http/client", lmodhttpclient.Open)
	util.Run(l, "main.lua")
}
```

```lua
local io = require 'io'
local client = require 'http/client'

local resp = client.get('https://github.com/ofunc/lmodhttpclient')
io.copy(resp, io.stdout)
resp:close()
```

## Dependencies

* [ofunc/lua](https://github.com/ofunc/lua)
* [publicsuffix](https://golang.org/x/net/publicsuffix)

## Documentation

### client.head(url)

Issues a `HEAD` to the specified `url`.

### client.get(url)

Issues a `GET` to the specified `url`.

### client.post(url, x[, y])

Issues a `POST` to the specified `url`.
`x` is a `Content-Type` and `y` is a `io.reader`.
Or `x` is a form data.

### client.fetch(options)

Sends an HTTP request and returns an HTTP response.
```
options.url: The path name to send the request to.
options.method: The HTTP method to use.
options.body: Body that you want to add to your request.
options.header: Header to append to the request before sending it.
```

### client.encode(data)

Encodes the `data` into `URL encoded` form.
