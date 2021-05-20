package internal

import (
	"strings"
)

/* Types */
type Portal struct {
	Alias string
	Path  string
}

type Command struct {
	Directive string
	Alias     string
}

/* Interfaces */
type Input interface {
	Trim() string
}

func (c *Command) Trim() string {
	return strings.TrimSpace(c.Alias)
}

func Conn() (*Database, error) {
	// f := os.Getenv("APPARATE_CONF")

	// if f == "" {
	// 	return nil, ErrNoConfigVar
	// } else if !Exists(f) {
	// 	return nil, ErrFileNotFound(f)
	// }

	f := "/home/goldmund/bolt"

	db, err := Open(f)
	if err != nil {
		return nil, err
	}

	return db, nil
}
