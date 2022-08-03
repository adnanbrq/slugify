# Slugify

Simple server that connects to a PostgreSQL database to save known urls under a custom slug.\
To follow a saved link you have to visit following url *{URL}/link/:slug*

## Contents

- [Customize](#customize)
- [Installation](#installation)
- [Usage](#usage)
- [Dependencies](#dependencies)

## Customize

Clone this Repository and change it how you like.

```sh
$ git clone https://github.com/adnanbrq/slugify.git
$ code ./slugify
```

## Usage

Either run it with go itself or build a Docker Image and start a Container

```sh
$ git clone https://github.com/adnanbrq/slugify.git
$ docker build -t slugify .
$ docker run -ti -p 3000:3000 -e PORT=3000 -e AUTO_MIGRATE=false -e DSN='postgresql://user:pass@host.docker.internal:5432/db' slugify
```

## Dependencies

- [https://github.com/adnanbrq/validation - v1.1.2](https://github.com/adnanbrq/validation)
Input validation
- [https://github.com/gofiber/fiber - v2.36.0](https://github.com/gofiber/fiber)
Simple and easy web framework
- [https://gorm.io](https://gorm.io/)
ORM to access the database