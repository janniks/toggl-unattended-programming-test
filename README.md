# Toggl Backend Unattended Programming Test

_March 6th 2020_

_This is a test as part of a job interview at Toggl.
The test consists of a simple API backend to be written in Go.
The test should be completed in three hours._

> _**Update** March 12th 2020_
> 
> I was invited to interview at Toggl.
> The interview invite already included a clause agreeing to a salary of USD 50,000.
> Toggl is a great company building great products and I like its transparency.
> However, I will be not pursuing employment at Toggl at this time.
> I still have a list of companies with products and ideologies that I love (some of them with [globally-competitive compensation](https://github.com/ymslavov/established-remote)).
>
> So, it looks like the programming test was good enough—although, I'd still love to get some feedback if anybody with Go experience ever finds this repository.

---

♣️♦️
The goal of the programming test is to create an API that can create decks of playing cards.
Those decks of cards can be requested and viewed.
Additionally, cards can be drawn from a deck—removing the card from it and leaving fewer remaining cards in the deck.
♠️♥️ — [`full specification`](./spec.pdf)

## Usage ⚡️

> The current API is live at [`toggl.janniks.com/ping`](https://toggl.janniks.com/ping)

The API structure is very simple:
```
GET  /ping                # pong
POST /deck                # create new deck (optional parameters: cards, shuffle)
GET  /deck/:deck_id       # fetch deck by id
GET  /deck/:deck_id/draw  # draw cards from deck (optional parameters: count)
```

_The exact usage can be inspected via the [`api.postman_collection.json`](./api.postman_collection.json) postman collection._

## Run 🛠

> All commands are to be run from the root project directory.
> This project uses Go version `1.14`.
> Current `go` executables should detect and install dependencies correctly.
> Set `GIN_MODE` to `release` for production deploys.

To manually install dependencies run:
```
go get -d -v ./...
go install -v ./...
```

To start the complete API run:
```
go run initial.go
```

To test all available test files run:
```
go test -v ./...
```

> Additionally, a [`Dockerfile`](./Dockerfile) is provided that will run anywhere Docker runs.

## Preparation 📚

This is the first [Go](https://golang.org/) project that I worked on. To prepare, I went through the following resources beforehand

- I read through [Learn Go in Y Minutes](https://learnxinyminutes.com/docs/go/) which gave me a super quick and dirty introduction to Go
- I followed the project structuring guidelines of [How to Write Go Code](https://golang.org/doc/code.html)
- I quickly read over the most import idioms and patterns of Go at [Effective Go](https://golang.org/doc/effective_go.html)
- I also chose [GoLand](https://www.jetbrains.com/go/) as an IDE and went through their onboarding tutorial

To choose a backend framework and ORM I simply launched a few Google searches.
Go offers many promising solutions, but I chose one of the most popular (although lacking documentation) frameworks [Gin](https://github.com/gin-gonic/gin) along with [GORM](https://github.com/jinzhu/gorm).

## Experience 🎡

The code itself is fairly self-documenting and was super fun to write.

Using Go, Gin, and GORM was fairly straight-forward and I could get going very quickly.
GoLand was also a great choice with many cool features and tricks (like [postfix completions](https://twitter.com/golandide/status/991301502449963009)). 

I chose SQLite as a file database (perfect for a proof-of-concept project) which can easily be replaced by any other SQL-type database without any logic changes.

The actual data being stored for cards in a deck is simply an `integer` array.
The `integer` values are converted to cards (with values and suits) during runtime.

The only notable problems I encountered are listed in the subsequent pitfalls section.
Due to the time constraint, I was not able to complete all planned refactoring. For example, the [`initial_test.go`](./initial_test.go) API test file is not very DRY. Additionally, there are some test cases and parameter edge cases that were missed and could cause unexpected behavior.

## Pitfalls ⚠️

- No generics in Go _([very interesting approach](https://blog.golang.org/why-generics), yet somewhat annoying due to other issues mentioned below)_
- Constraints to certain data-types due to lack of support in GORM ([gorm#1588](https://github.com/jinzhu/gorm/issues/1588)) lead to some lost time.
- Testing best practices are not very clear/defined in the Go community. Auto-generated code was often very different from the sample code of multiple projects.
- Now—knowing what I learned through this project—I would start writing tests and API validation first. Go/Gin was very new to me which lead me to change project structures a couple of times and made true TDD a bit difficult. 

## Roadmap 🚧

_Some stuff that I'd continue if there had been more time_

- DRY repeated code blocks (`db` fetching, http requests in tests, etc.)
- Shorten (aka modularize) some of the longer API controller methods
- Simplify conversion of card `id` to `code` and vice-versa
- Add more tests

## Feedback 💬

If you notice anything in this repository—bad code style, bad practices, bugs, very wet code, etc.—[please let me know!](https://twitter.com/messages/compose?recipient_id=82144365) I enjoyed this project in Go and would like to learn more 🙏

## Acknowledgments 🙌

The specification for this programming test came from [Toggl](https://toggl.com/)—an awesome company making great [time-tracking](https://toggl.com/features/) and [hiring](https://toggl.com/hire/) tools.
The test is likely inspired by the [Deck of Cards](https://deckofcardsapi.com/) project.

I have asked Toggl whether I may keep this programming test open-source and online.
So far, I have not been prohibited to share the test and the results.
I will readily remove the repository if anything changes.
