# Rest API
A Rest API that handles data on video games and saves all data in a sqlite3 database.

## Get a list of all available games

`GET /api/games`

## Get a specific game

`GET /api/games/id`

## Create a new game

`POST /api/games`

## Update an existing game

`PUT /api/games/id`

## Delete an existing game

`DELETE /api/games/id`

## Issues
- When running transport_test.go with `Database` (the real DB) assigned to the DB variable, `unable to open database file: no such file or directory` occurs.
- The methods in database_mock.go contains no logic and only return dummy values or nil values.
- database_test.go contains no tests
- API doesn't return specific errors
