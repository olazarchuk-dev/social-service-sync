# social-service client

* `Рруководство чтобы начать работу с Next.js`: https://www.freecodecamp.org/news/nextjs-tutorial
* `Использование 'State Hook'`: https://reactjs.org/docs/hooks-state.html
  * `Хуки` — это новое дополнение в React 16.8. Они позволяют вам использовать состояние и другие функции React без написания класса.
* `Введение в Next.js и React Framework | Создание приложение Next.js`: https://nextjs.org/learn/basics/create-nextjs-app
* `Как добавить слайдер в Next.js`: https://www.geeksforgeeks.org/how-to-add-slider-in-next-js
* `Создание 'pages/_app.js' в Next.js`: https://nextjs.org/docs/messages/css-global
* [Bootstrap — Документация / Примеры](https://getbootstrap.com/docs/4.0/components/buttons)
* [Bootstrap — Справочник с часто используемыми элементами](https://www.rotamaxima.com/en/bootstrap-4-reference-guide-of-the-most-used-elements)
  * `Bootstrap` — это веб-фреймворк для разработки интерфейсов и внешних компонентов для веб-приложений с использованием HTML, CSS и JavaScript, основанный на шаблонах дизайна для типографики, улучшающий взаимодействие с пользователем.
  * Чтобы начать использовать **Bootstrap** все что нужно сделать, это импортировать файлы CSS и JS в веб-приложение:
  ```haml
  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">
  <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
  <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
  ```
  * Важно отметить что в **Bootstrap** есть система разрешения экрана: 
    * `sm` = 576 пикселей
    * `md` = 768 пикселей
    * `lg` = 992 пикселей
    * `xl` = 1200 пикселей 


### tech stack
+ **Front End**
  - Node.js 12.0+
  - React.Js
  - Next.Js
  - Bootstrap 4
  - Tailwind Css

```shell script
> npm install
> npm run-script build
> yarn run dev
```

---

* [Next.js — Подробное руководство 1](https://habr.com/ru/company/timeweb/blog/588498)
* [Next.js — Подробное руководство 2](https://habr.com/ru/company/timeweb/blog/590157)
* [Next.js — Краткое руководство 1](https://pxstudio.pw/blog/chto-takoe-next-js-i-dlya-chego-on-nuzhen)
* [Next.js — Краткое руководство 2](https://pxstudio.pw/blog/poluchenie-dannyh-v-next-js)
* [Next.js — Краткое руководство 3.1](https://nextjs.org/docs)
* [Next.js — Краткое руководство 3.2](https://nextjs.org/docs/getting-started)
* [Next.js — Краткое руководство 3.3](https://nextjs.org/learn/basics/create-nextjs-app/setup)
* [Next.js — Краткое руководство 3.4](https://nextjs.org/docs/api-reference/cli)
* [Next.js — Краткое руководство 3.5](https://nextjs.org/learn/basics/create-nextjs-app)
  * `Next.js` — это фреймворк, основанный на **React**, который позволяет создавать веб-приложения с улучшенной производительностью с помощью дополнительных функций предварительного рендеринга, таких как полноценный рендеринг на стороне сервера (SSR) и статическая генерация страниц (SSG).
  * `Next.js` — лучше всего подходит, когда необходимо создать оптимизированный лендинг или домашнюю страницу, а также любые другие страницы, которые полагаются на органический поисковый трафик.
* Запуск приложения в разных режимах:
  * `build` — собирает приложение Next.js для Production;
  * `start` — запускает Next.js в режиме Production;
  * `dev` — запускает Next.js в режиме разработки;
* Требования и окружающая среде:
  * (помимо самого фреймворка Next.js) потребуются Node.js, `npm` version 12-16 и `npx` (обычно `npm` и `npx` одновременно устанавливаются при установке `Node.js`)
  * [Update node to v12 on ubuntu](https://stackoverflow.com/questions/60679889/update-node-to-v12-on-ubuntu)

* [Разница между протоколами TCP и UDP](http://pyatilistnik.org/chem-otlichaetsya-protokol-tcp-ot-udp)
  * `TCP` — гарантирует доставку пакетов данных в неизменных виде, последовательности и без потерь; TCP требует заранее установленного соединения;
  * `UDP` — обеспечивает более высокую скорость передачи данных и ничего не гарантирует; UDP соединения не требует; UPD не содержит функций восстановления данных;

**TCP — надежнее и осуществляет контроль над процессом обмена данными.**

**UDP — обеспечивает более высокую скорость передачи данных.**
