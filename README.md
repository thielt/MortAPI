# MortAPI

On a Window's Machine:
Ensure that you have the msys2 GNU and that your systems path has the bin file from the msys2 directory for golangs use of the C compiler "gcc".

Build go module:
$env:CGO_ENABLED=1; go build

To start the application:
go run main.go

--Ensure you have also downloaded the executable file of sqlite3 from their website. 
I will be building a local database so these files will now be readable on your ide.
sqlite3 can be accessible on your local terminal using the sqlite3 command ONLY in the directory's sql folder

Commands to access and manipulate a sample table:
<dir> - go run main.go to build the table
sqlite3 users.db 
> .table (to check the table exists)
> INSERT INTO users (name, email, password) VALUES ('Thiel', 'thiel@gmail.com', 'password');
> SELECT * FROM users; 
> .exit to exit sqlite3 and back to the main terminal 

cd .. back to root dir

Install Go debugger extension if not wanting to use the terminal.