# MortAPI
An API that calculates mortgages based on varying rates and locations(not sure what else we're adding in this)

## Installation(Windows Machine)

We will be using gorilla mux and sqlite3:

```sh
go get -u github.com/gorilla/mux
go get -u github.com/mattn/go-sqlite3
```

Make sure you have [msys2 GNU](https://www.msys2.org/) installed.

Make sure its in your systems PATH. 

This is for Go's use of the C-compiler "gcc".

To start the application and build the users.db table:

```sh
go run main.go
```

Or install vscode's golang debugger extension (`f5`)

## sqlite3

For the database, use [sqlite3](https://www.sqlite.org/download.html).

Sample commands once you're in the sqlite3 terminal:

```sh
sqlite3 users.db (to start the sql terminal)
.table (to check the tables that exist)
.exit (to exit sqlite3 and back to the main terminal)
```

Basic sql command to insert into the users.db:
```sh
INSERT INTO users (name, email, password) VALUES ('Thiel', 'thiel@gmail.com', 'password');
```

When using `.exit` and leaving sql folder back to the root directory:
```sh
cd ..
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
```
=======
``` 
>>>>>>> 53267546ddcf5993bdd0c4b921df39a025873e68
=======
```
>>>>>>> dbc0f0a (cleanup-getUser (#2))
=======
``` 
>>>>>>> d5a3455 (Update README.md)
