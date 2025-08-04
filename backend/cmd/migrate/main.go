package main

import (
	"fmt"
	"samokat/internal/setting"
)

func main() {
	app := setting.App{}
	err := app.Migrate()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database migration completed successfully.")
}
