package main

import (
	"go-design-patterns/patterns/creational/singleton/database"
	myLogger "go-design-patterns/patterns/creational/singleton/my-logger"
	"log"
	"os"
	"time"
)

func main() {
	logger := myLogger.GetLoggerInstance()

	logger.SetLevel(1)
	logger.Log("Hello world")

	logger = myLogger.GetLoggerInstance()
	logger.Log("Testing logger again ")

	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_MAX_IDLE_CONNECTIONS", "10")

	// Testing thread safefety
	for i := 0; i < 10; i++ {
		log.Println("Iteration ", i)
		go func() {
			db, err := database.GetDatabaseInstance()
			if err != nil {
				log.Fatal(err)
			}
			db.GetInstance()
		}()
	}
	time.Sleep(time.Second * 1)
}
