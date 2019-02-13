# ganki
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fdulev%2Fganki.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fdulev%2Fganki?ref=badge_shield)


A go implementation of flashcard spaced repetition software. 
A server and a client are provided that can work together as
a local application or as a remote flashcard service.

The server provides a REST API which can be used by any UI implementation.
Initially a CLI will be developed and, if there is time, an Electron GUI will be implemented.

#### Features
- Custom deck and card creation
- Pictures and audio in cards
- Several input formats for cards (including markdown)
- Saving decks in git repos (this allows public card suggestions)
- Single/multi user support

#### Installation

```
$ go get github.com/dulev/ganki
```

#### Run

To run as a standalone application.
```
$ ganki
```

To connect to a specific server. This can be done from the UI as well.
```
$ ganki --server flashcards-hub.com
```


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fdulev%2Fganki.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fdulev%2Fganki?ref=badge_large)