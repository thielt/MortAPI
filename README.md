# MortAPI

On a Window's Machine:
Ensure that you have the msys2 GNU and that your systems path has the bin file from the msys2 directory for golangs use of the C compiler "gcc".

Build go module:
$env:CGO_ENABLED=1; go build

To start the application:
go run main.go