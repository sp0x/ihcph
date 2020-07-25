# ihcph
[![sp0x](https://circleci.com/gh/sp0x/ihcph.svg?style=shield)](https://circleci.com/gh/sp0x/ihcph)
[![codecov](https://codecov.io/gh/sp0x/ihcph/branch/master/graph/badge.svg)](https://codecov.io/gh/sp0x/ihcph)
[![Go Report Card](https://goreportcard.com/badge/github.com/sp0x/ihcph)](https://goreportcard.com/report/github.com/sp0x/ihcph)

This project tracks the *IHCPH* website for available application dates/hours and notifies you on telegram whenever there's a new date/hour.  
It was created to help you find cancelled bookings.

### Installation
Just grab the latest image from [Docker hub](https://hub.docker.com/r/sp0x/ihcph).
You can use the existing docker-compose.yml file to run this project.  
Here's an example of a docker-compose file:
```yaml
version: "2.0"
services:
  ihcph:    
    image: sp0x/ihcph
    env_file: .env
```

### Configuration

Environment variables:
- `TELEGRAM_TOKEN` To use the telegram functionality you need to provide a token
- `INDEXER` The name of the indexer, by default that's ihcph.kk.dk

### Supported sites  
- ihcph.kk.dk - `International House Copenhagen`
  
You can add new sites in a few ways:
- Add a new YML file in the `./indexes/` directory
- Add a new YML file in the `~./.ihcph/indexes/` directory
- Add a new YML file in the sites directory of the project and rebuild it.
