package main

import (
	"counterapp/internal/config"
	"counterapp/internal/dao"
	"counterapp/server/api"
	"fmt"
)

func main() {
	cfg := config.Load()

	db, err := dao.Connect()
	if err != nil {
		fmt.Printf("error occured while creating db: %s", err)
		return
	}
	fmt.Println("connected to database")

	err = dao.Migrate(db)
	if err != nil {
		fmt.Printf("error while migrating schemas: %s", err)
		return
	}
	fmt.Println("added the required tables to counter_db")

	router := api.SetupRouter(db)
	
	addr := fmt.Sprintf(":%s", cfg.Port)
	fmt.Printf("Starting server on %s\n", addr)
	
	if err = router.Run(addr); err != nil {
		fmt.Printf("error starting server: %s", err)
		return 
	}

}