# SpaceCMS API

This repository holds a Golang API service for a CMS (Content Management System) personal project that I'm building open source!

## Tools used within this repo:
  - Logrus (Used for logging)
  - Mockery (Used for mock testing, usaully for database interface mocking)
  - Goa (Used to generate the controllers from a DSL, and generates Swagger Documentation)
  - Argon2id (Used for hashing of passwords)

## Controllers
  - All controllers require unit tests, and may require mocks used within the tests.

## Database 
  - This service uses a SQLite3 database and the Go SQL driver. 
  - I plan to upgrade this to use a package called `sqlc` to generate type safe SQL queries into Go structs.
