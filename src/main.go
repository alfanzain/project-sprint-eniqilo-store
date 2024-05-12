package main

import (
	"github.com/alfanzain/project-sprint-eniqilo-store/src/databases"
	"github.com/alfanzain/project-sprint-eniqilo-store/src/http"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	databases.ConnectPostgre()

	h := http.New(&http.Http{})
	h.Launch()
}
