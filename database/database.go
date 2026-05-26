package database

var connection string

func init () {
	connection = "MariaDb"
}

func Geterring () string {
	return connection
}