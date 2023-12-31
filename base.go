package main

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var conn *pgx.Conn

// connection
func ConnectPostgres() error {
	var err error

	if err = godotenv.Load()
	err != nil {
		fmt.Println((err.Error()))
		return err
	}
	dbUrl, _ := os.LookupEnv("DATABASE_URL")

	conn, err = pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		return err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}

func Create(notes *NotesType) error {
	_, e := conn.Exec(context.Background(), `INSERT INTO Notes (title, text ) VALUES ($1, $2, $3)`, notes.Title, notes.Text)
	if e != nil {
		fmt.Println(e.Error())
	}
	return e
}


func Update(notes *NotesType) error {
	_, e := conn.Exec(context.Background(), `UPDATE Notes SET Title = $1, Text = $2 WHERE id = $3`, notes.Title, notes.Text, &notes.Id)
	if e != nil {
		fmt.Println(e.Error())
	}
	return e

}

func GetNotes() ([]NotesType, error) {

	row, e := conn.Query(context.Background(), `SELECT * FROM notes`)

	if e != nil {
		fmt.Println(e.Error())
		return nil, e
	}

	defer row.Close()

	var note NotesType
	var notesArr = make([]NotesType, 0)

	for row.Next() {
		e = row.Scan(&note.Id, &note.Title, &note.Text)
		if e != nil {
			return nil, e
		}

		notesArr = append(notesArr, note)
	}

	e = row.Err()

	if e != nil {
		return nil, e
	}

	return notesArr, nil
}

func Delete(notes *NotesType) error {
	_, e := conn.Exec(context.Background(), `DELETE FROM notes WHERE id = $1`, notes.Id)
	if e != nil {
		fmt.Println(e.Error())
	}
	return e

}
