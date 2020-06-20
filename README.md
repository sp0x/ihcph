# ihcph
[![sp0x](https://circleci.com/gh/sp0x/ihcph.svg?style=shield)](https://circleci.com/gh/sp0x/ihcph)
[![codecov](https://codecov.io/gh/sp0x/ihcph/branch/master/graph/badge.svg)](https://codecov.io/gh/sp0x/ihcph)
[![Go Report Card](https://goreportcard.com/badge/github.com/sp0x/ihcph)](https://goreportcard.com/report/github.com/sp0x/ihcph)

This project tracks the *IHCPH* website for available application dates/hours and notifies you on telegram whenever there's a new date/hour.
It was mainly created because people cancel their bookings sometimes and you can miss-out on that.

### Configuration

Environment variables:
- `TELEGRAM_TOKEN` To use the telegram functionality you need to provide a token
- `INDEXER` The names of the indexer, by default that's ihcph.kk.dk

### Supported sites  
- ihcph.kk.dk - `International House Copenhagen`
  
You can add new sites in a few ways:
- Add a new YML file in the `./indexes/` directory
- Add a new YML file in the `~./.ihcph/indexes/` directory
- Add a new YML file in the sites directory of the project and rebuild it.
