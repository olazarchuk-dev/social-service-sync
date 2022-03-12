# social-service-sync

+ Language: `Golang`
+ Пакеты:
  - `gorilla / mux` — URL-маршрутизатор и диспетчер, будем использовать этот пакет для сопоставления путей URL с их обработчиками.
  - `jinzhu / gorm` — ORM-библиотека для Golang, будем используем этот пакет, чтобы взаимодействовать с базой данных.
  - `dgrijalva / jwt-go` — используется для подписи и проверки JWT-токенов.
  - `joho / godotenv` — используется для загрузки .env-файлов в проект.

```shell script
CREATE TABLE "accounts" ("id" serial,"created_at" timestamp with time zone,"updated_at" timestamp with time zone,"deleted_at" timestamp with time zone,"email" text,"password" text,"token" text , PRIMARY KEY ("id"))  
CREATE INDEX idx_accounts_deleted_at ON "accounts"(deleted_at)   
CREATE TABLE "contacts" ("id" serial,"created_at" timestamp with time zone,"updated_at" timestamp with time zone,"deleted_at" timestamp with time zone,"name" text,"phone" text,"user_id" integer , PRIMARY KEY ("id"))  
CREATE INDEX idx_contacts_deleted_at ON "contacts"(deleted_at)   
```

```shell script
host=172.31.75.175 user=trademaster dbname=postgres sslmode=disable password=RVHmBTK6In1nvGryKxEy
8000
```

* [Socket.IO — Краткое руководство](https://coderlessons.com/tutorials/kompiuternoe-programmirovanie/uznaite-socket-io/socket-io-kratkoe-rukovodstvo)
  * `Socket.IO` — это библиотека JavaScript для веб-приложений реального времени (RTA). Он обеспечивает двустороннюю связь в реальном времени между веб-клиентами и серверами.
* [Протокол WebSocket — Краткое руководство](https://nodejsdev.ru/doc/websocket)
  * `WebSocket` - это протокол передачи данных, основанный на протоколе TCP обеспечивающий обмен сообщениями между клиентом и сервером в режиме реального времени.
* [Работа с WebSocket и Socket.IO в Go — Учебник](https://tutorialedge.net/golang/golang-websockets-tutorial)
* [Как использовать WebSocket в Golang — Пошаговое руководство](https://yalantis.com/blog/how-to-build-websockets-in-go)

---

+ [Пример приложения чата](https://github.com/kyopark2014/webchat-golang-socketio)
  - В котором socket.io используется для соединения сервера и клиента;
  - В этом проекте была развернута структура PUBSUB для поддержки Say, Join and Leave;
  - Оно поддерживает 1-to-1 и групповой чат с использованием канала связи;
+ [Пример приложения чата при поддержке Redis](https://levelup.gitconnected.com/building-chat-service-in-golang-and-websockets-backed-by-redis-b42a8784636c)
  - socket.io используется для соединения сервера и клиента;
  - развернута структура PUBSUB для поддержки Say, Join and Leave;
  - поддержка Redis и Redis-Persistence;
+ [Примеры Node.js Websocket с Socket.io](https://rukovodstvo.net/posts/id_1357)
