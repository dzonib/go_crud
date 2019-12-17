package main

import "todo/postgres"

import "github.com/go-pg/pg"

func main() {
	DB := postgres.New(&pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "todo_dev",
	})

	defer DB.Close()
}
