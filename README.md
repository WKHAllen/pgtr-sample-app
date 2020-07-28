# PGTR Sample App

This is a sample app that uses a PostgreSQL-Go-TypeScript-React stack.

It can serve as a template for projects designed using the PGTR stack. Simply clone the repository, and you are good to go.

## Clone

```console
$ git clone https://github.com/WKHAllen/pgtr-sample-app.git
```

## Deployment

This sample app is [deployed here](https://pgtr.herokuapp.com/) on [Heroku](https://heroku.com/).

## Stack

For the frontend, [React](https://reactjs.org/) was the obvious choice, as it is one of the most powerful frontend frameworks out there. The [TypeScript](https://www.typescriptlang.org/) variant of React is used because it is very easy to make mistakes when using JavaScript.

Most modern languages could work for the backend. [Go](https://golang.org/) was chosen because of its simplicity and because of one of its most powerful web frameworks, [Gin](https://github.com/gin-gonic/gin). 

[PostgreSQL](https://www.postgresql.org/) is used for the database because of its popularity and advanced features. In addition, Heroku has a helpful addon for interacting with Postgres databases.
