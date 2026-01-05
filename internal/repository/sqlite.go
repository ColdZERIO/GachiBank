package repository

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

const ( // Просто констатнты с названием таблиц в БД для будующего образения уже по ним
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	userListsTable  = "users_list"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

func NewSQLiteDB(driver, path string) (*sql.DB, error) {
	db, err := sql.Open(driver, path)
	if err != nil {
		log.Fatal("error opening DB")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error connectimg to DB")
		return nil, err
	}

	return db, nil
}
