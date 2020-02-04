package actions

import (
	"flag"
	. "github.com/davecremins/ToDo-Manager/config"
	"log"
)

var actionMap map[string]ConfigFunc

func Process(args []string, defaultConfig *ToDoConfig) {
	if len(args) < 2 {
		log.Fatal("expected subcommands to perform an action")
	}

	actionMap = make(map[string]ConfigFunc)
	actionMap["init"] = initActionMakeup(args, defaultConfig)
	actionMap["newday"] = newDayActionMakeup(args, defaultConfig)
	actionMap["newtodo"] = newTodoActionMakeup(args, defaultConfig)
	actionMap["today"] = todaysTodosActionMakeup(args, defaultConfig)

	action, ok := actionMap[args[1]]
	if !ok {
		log.Fatal(args[1] + " subcommand is not supported right now :(")
		flag.PrintDefaults()
	}
	action(defaultConfig)
}
