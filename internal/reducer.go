package internal

import (
	"fmt"
	"os"
)

func AddWarpPoint(c *Command) {
	// get current Path
	loc, err := os.Getwd()
	if err != nil {
		ErrExit("E_NO_CUR_DIR") // TODO
	}

	p := Portal{
		c.Trim(),
		loc,
	}

	// init db conn
	db, err := Conn()

	if err != nil {
		ErrExit(err.Error())
	}

	defer db.Close()

	// TODO check if we are overwriting and inform the user
	if err = db.Put(&p); err != nil {
		ErrExit(fmt.Sprintf("Failed to add warp-point %s for %s", p.Alias, p.Path))
	}

	OkExit(fmt.Sprintf("Added new warp-point %s for %s", p.Alias, p.Path))
}

func RemoveWarpPoint(c *Command) {

	p := Portal{
		Alias: c.Trim(), // TODO check for empty
		Path:  "",
	}

	// init db conn
	db, err := Conn()
	if err != nil {
		ErrExit(err.Error())
	}

	defer db.Close()

	if err = db.Delete(&p); err != nil {
		if err == ErrNotFound {
			OkExit(err.Error())
		}

		ErrExit(err.Error()) // TODO check for del error and not just the one we manually set
	}

	OkExit(fmt.Sprintf("Warp-point %s has been removed", p.Alias))
}

func GotoWarpPoint(c *Command) {
	p := Portal{
		Alias: c.Trim(), // TODO check for empty, regex alphanumeric only
	}

	// init db conn
	db, err := Conn()
	if err != nil {
		ErrExit(err.Error())
	}

	defer db.Close()

	if err = db.Get(&p); err != nil {
		if err == ErrNotFound {
			OkExit(err.Error())
		}

		ErrExit(err.Error()) // TODO check for del error and not just the one we manually set
	}

	// does the Path we retrieved still exist?
	if !Exists(p.Path) {
		// if we were able to `Get` the key, we can presumably `Delete` it without too much concern
		db.Delete(&p) // TODO prompt user
		ErrExit(fmt.Sprintf("Warp-point %s for %s points to an invalid Path. Removed from storage", p.Alias, p.Path))
	}

	// print path to stdout, return code 3 to indicate to wrapper which dir to cd into
	fmt.Print(p.Path) // TODO hide output
	os.Exit(3)
}
