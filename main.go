package main

import (
	"fmt"
	"os"

	"github.com/MatthewZito/apparate/internal"
)

/* Constants */
const (
	GOTO        = "goto"
	REMOVE      = "remove"
	ADD         = "add"
	HELP        = "help"
	CONFIG_PATH = "apparate.conf.json"
)

/* MAIN */
func main() {
	args := os.Args[1:]

	if !internal.Itob(len(args)) {
		internal.ErrExit("E_NO_CMD")
	} else if !(len(args) == 2) {
		internal.ErrExit("Insufficient arguments were provided")
	}

	cmd := internal.Command{
		Directive: args[0],
		Alias:     args[1],
	}

	switch cmd.Directive {
	case ADD:
		internal.AddWarpPoint(&cmd)
	case REMOVE:
		internal.RemoveWarpPoint(&cmd)
	case GOTO:
		internal.GotoWarpPoint(&cmd)
	case HELP:
	default:
		fmt.Println(args)
		os.Exit(1)
	}

}
