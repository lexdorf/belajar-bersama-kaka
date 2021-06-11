package database

var connection string

func init(){
	fmt.Println("Int dipanggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}