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


---

* `Convert to time in Golang from milliseconds`: https://stackoverflow.com/questions/31744970/convert-to-time-in-golang-from-milliseconds
* `Пакет bson для работы с MongoDB в Golang`: https://golang-blog.blogspot.com/2020/06/bson-mongodb-golang.html
* `How to convert mongodb go driver's primitive.Timestamp type back to Golang time.Time type`: https://stackoverflow.com/questions/64418512/how-to-convert-mongodb-go-drivers-primitive-timestamp-type-back-to-golang-time
```text
BSON Timestamps содержат два значения:
- 'T' for the seconds since Unix epoch
- 'I' for an incrementing ordinal for operations within a given second

Поэтому, если вы хотите преобразовать временную метку bson в time.Time, вы можете просто использовать:
time.Unix(timestamp.T, 0)

Точно так же, чтобы преобразовать текущее time.Time в primitive.Timestamp, мы можем использовать:
primitive.Timestamp{T: uint32(time.Now().Unix()), I: 0}
```

#### Import cycle not allowed

* https://stackoverflow.com/questions/28256923/import-cycle-not-allowed

Here is an illustration of your first import cycle problem:
```text
                  project/controllers/account
                     ^                    \    
                    /                      \
                   /                        \ 
                 \/                         \/
         project/components/mux <--- project/controllers/base
```
