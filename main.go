package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album

	urlExample := "postgres://postgres:secret@db:5432/postgres"
	os.Setenv("DATABASE_URL", urlExample)
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
    // Loop through rows, using Scan to assign column data to struct fields.

	rows, errRows :=  conn.Query(context.Background(), "SELECT * FROM album WHERE artist=$1", name)
	if errRows != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if errAlbum := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); errAlbum != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, errAlbum)
		}
		albums = append(albums, alb)
	}

    if errLast := rows.Err(); errLast != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, errLast)
    }
 
    return albums, nil
}

func main() {
	albums, err := albumsByArtist("John Coltrane")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)
}