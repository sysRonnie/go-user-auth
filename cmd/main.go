package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/sysronnie/go-user-auth/cmd/api"
	"github.com/sysronnie/go-user-auth/config"
	"github.com/sysronnie/go-user-auth/db"
	"github.com/sysronnie/go-user-auth/view/template"
)

func main() {
	app := echo.New()

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	templateHandler := template.TemplateHandler{DB: db}

	app.GET("/", templateHandler.ShowLandingPage) 
	apiServer := api.NewAPIServer(db)

	// Register API routes to the Echo instance
	if err := apiServer.APIService(app); err != nil {
		log.Fatal("Failed to register API service:", err)
	}

	if err := app.Start(":4200"); err != nil {
		log.Fatal(err)
	}

}



func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully Connected!")
}