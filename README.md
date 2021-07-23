# go_restapi
This is a Rest API for handling data on games. It's written in Go and all data is saved in a sqlite3 database.

Things to fix:
- Game ID handling (user should not need to specify a game ID when creating or updating a game - that should be handled automatically)
- Error handling (is there a lot of redundant code? are there places where error handling is missing? is panic ok to use or are there better alternatives?)
- getGames function is super messy
- The database is opened at the start of every function - that's repetitive and there must be a more efficient way open the db
- This can definitely be cleaned up by splitting it into different files