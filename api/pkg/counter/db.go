package counter

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Counter struct {
	db *sql.DB
}

func NewCounter(connStr string) (*Counter, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Get err", err)
		return nil, err
	}
	// conn이 끊길 수도 있으니 defer는 주석
	// defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Database connected successfully")

	return &Counter{
		db: db,
	}, nil
}

func (c *Counter) IncreaseCounter() error {
	var count int
	err := c.db.QueryRow(`
		UPDATE ping_counter 
		SET count = count + 1 
		WHERE id = (SELECT id FROM ping_counter ORDER BY id LIMIT 1)
		RETURNING count;
	`).Scan(&count)

	if err != nil {
		log.Printf("Error updating counter: %v", err)
		return err
	}

	return nil
}
