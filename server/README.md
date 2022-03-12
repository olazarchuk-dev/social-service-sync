# social-service server

### tech stack
+ **Back End**
  - Golang 1.13+
  - Fiber
    - is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go
    - ( https://github.com/gofiber/fiber )
  - Postgre SQL
  - Websocket
    - [The WebSocket Protocol](https://www.rfc-editor.org/rfc/rfc6455.txt)
    - ( https://github.com/gofiber/websocket ) Based on Fasthttp WebSocket for Fiber with available *fiber.Ctx methods like Locals, Params, Query and Cookies
    - ( https://github.com/fasthttp/websocket ) Gorilla WebSocket is a Go implementation of the WebSocket protocol

```shell script
> go build main.go
```
