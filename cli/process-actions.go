package actions

import (
	"fmt"
	. "github.com/davecremins/ToDo-Manager/config"
	"log"
	"os"
)

var actionMap map[string]Action

func init() {
	stranglePattern := os.Getenv("STRANGLE")
	if stranglePattern != "" {
		config := LoadConfig()
		actionMap = make(map[string]Action)
		actionMap["list"] = showTasks(config)
		actionMap["add"] = newTask(config)
	} else {
		config := LoadConfig()
		actionMap = make(map[string]Action)
		actionMap["newtodo"] = newTodoActionMakeup(config)
		actionMap["today"] = todaysTodosActionMakeup(config)
		actionMap["complete"] = completeTodoActionMakeup(config)
		actionMap["move"] = moveTodoActionMakeup(config)
		actionMap["merge"] = mergeTodoActionMakeup(config)
	}
}

func Process(args []string) {
	if len(args) < 2 {
		log.Println("No command provided, printing default usage instead")
		printDefaults()
	} else {

		action, ok := actionMap[args[1]]
		if !ok {
			msg := fmt.Sprintf("Command '%s' not supported, printing default usage instead", args[1])
			log.Println(msg)
			printDefaults()
		} else {
			action(args)
		}
	}
}
