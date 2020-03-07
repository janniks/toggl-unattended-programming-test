# ⏰ Toggl Backend Unattended Programming Test

_March 6th 2020_

This an initial project as part of a job interview at toggl.

The test description describes a simple API backend to be written in [Go](https://golang.org/). The original specification can be found at [`/spec.pdf`](./spec.pdf).

## Run 🚀

> All commands are to be run from the root project directory.
> Set `GIN_MODE` to `release` for production deploys.

To start the complete API run:
```
go run initial.go
```

To test all available test files run:
```
go test -v ./...
```

## Preparation 📚

This is the first project in Go that I worked on. To prepare, I went through the following resources beforehand

- I read through [Learn Go in Y Minutes](https://learnxinyminutes.com/docs/go/) which gave me a super quick and dirty introduction to Go
- I followed the project structuring guidelines of [How to Write Go Code](https://golang.org/doc/code.html)
- I quickly read over the most import idioms and patterns of Go at [Effective Go](https://golang.org/doc/effective_go.html)
- I also chose [GoLand](https://www.jetbrains.com/go/) as an IDE and went through their onboarding tutorial

To choose a backend framework and ORM I simply launched a few Google searches. Go offers many promising solutions, but I chose one of the most popular (although lacking documentation) frameworks [Gin](https://github.com/gin-gonic/gin) along with [GORM](https://github.com/jinzhu/gorm).

## Experience 🎡

The code itself is fairly self-documenting and was super fun to write.

Using Go, Gin, and GORM was fairly straight-forward and I could get going very quickly. GoLand was also a great choice with many cool features and tricks (like [postfix completions](https://twitter.com/golandide/status/991301502449963009)). 

I chose SQLite as a file database (perfect for a proof-of-concept project) which can easily be replaced by any other SQL-type database without any logic changes.

The API structure is very simple:
```
POST /deck                # create new deck (parameters: cards, shuffle)
GET  /deck/:deck_id       # fetch deck by id
GET  /deck/:deck_id/draw  # draw cards from deck (parameters: count)
```

_The exact usage can be inspected via the [`/api.postman_collection.json`](./api.postman_collection.json) postman collection._

The only notable problems I encountered are listed in the subsequent pitfalls section.

## Pitfalls ⚠️

- No generics in Go 😱 _(somewhat annoying due to other issues mentioned below)_
- Constraints to certain data-types due to lack of support in GORM ([gorm#1588](https://github.com/jinzhu/gorm/issues/1588)) lead to some lost time.
- Testing best practices are not very clear/defined in the Go community. Auto-generated code was often very different from sample code of multiple projects.
- Now knowing what I learned through this project, I would start writing tests and API validation first. Go/Gin was very new to me which lead me to change project structures a couple of times and made true TDD a bit difficult. 
