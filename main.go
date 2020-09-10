package main

import (
	"fmt"
	"os"
	"path/filepath"
	"taskManager/cmd"
	"taskManager/db"
)

func main() {
	// home, _ := homedir.Dir()
	dbPath := filepath.Join("/Users/raghavtayal/go/src/taskManager", "tasks.db")
	errHandler(db.Init(dbPath))
	errHandler(cmd.RootCmd.Execute())
}

func errHandler(err error) {
	if err != nil {
		fmt.Println("Some error occured")
		os.Exit(1)
	}
}
