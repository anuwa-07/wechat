package main
import (
	"fmt"
	"os"
	"context"

	// internal packages
	"github.com/anuwa-07/wechat/pkg/sql"
)

func main() {
	// make a context for the application.
	ctx := context.Background();

	//
	dbConfig := sql.DBConfig{
		Username: "root",
		Password: "12345678",
		Host: "localhost",
		Port: "3306",
		Database: "wechat",
		MaxOpenConnections: 10,
		MaxIdleConnections: 5,
	}
	dbConn, err := sql.ConnInit(dbConfig);
	if err != nil {
		fmt.Println("[SQL] Conncetion Initilization Failed!", err);
		os.Exit(1);
	}
	fmt.Println("[SQL] Connection Initilized Successfully!");

	// TODO: call on wechat server and run the application...
}

