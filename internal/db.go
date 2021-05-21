package internal

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

type Database struct {
	db *bolt.DB
}

var (
	// ErrNotFound indicates the key provided for Get, Delete ops
	// is not extant in the db
	ErrNotFound = errors.New("warp-point does not exist") // TODO add alias

	// ErrNoConfigVar indicates the environment variable for the apparate config file was not found
	ErrNoConfigVar = errors.New("configuration file not found; did you set APPARATE_CONF to the desired Path?")

	bucketName = []byte("apparate")
)

// ErrFileNotFound pattern-matches against os.Stat file-not-found errors
func ErrFileNotFound(name string) error {
	return fmt.Errorf("could not resolve Path for configuration file %s", name)
}

// Open initializes and returns a pointer to a new database connection
func Open(Path string) (*Database, error) {
	opts := &bolt.Options{
		Timeout: 50 * time.Millisecond,
	}

	if db, err := bolt.Open(Path, 0640, opts); err != nil {
		return nil, err
	} else {
		err := db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists(bucketName)
			return err
		})
		if err != nil {
			return nil, err
		} else {
			return &Database{db: db}, nil
		}
	}
}

// Get retrieves a database record, if extant
func (conn *Database) Get(p *Portal) error {
	return conn.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucketName).Cursor()

		if k, v := c.Seek([]byte(p.Alias)); k == nil || string(k) != p.Alias {
			return ErrNotFound
		} else {
			d := gob.NewDecoder(bytes.NewReader(v))
			return d.Decode(&p.Path)
		}
	})
}

// Put updates a database record if extant, else creates it
func (conn *Database) Put(p *Portal) error {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(p.Path); err != nil {
		return err
	}

	return conn.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucketName).Put([]byte(p.Alias), buf.Bytes())
	})
}

// Delete updates a database record by removing it, if extant
func (conn *Database) Delete(p *Portal) error {
	return conn.db.Update(func(tx *bolt.Tx) error {
		c := tx.Bucket(bucketName).Cursor()

		if k, _ := c.Seek([]byte(p.Alias)); k == nil || string(k) != p.Alias {
			return ErrNotFound
		} else {
			return c.Delete()
		}
	})
}

// Close closes the database connection
func (conn *Database) Close() error {
	return conn.db.Close()
}
