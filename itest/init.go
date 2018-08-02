package itest

import (
	"fmt"
	. "github.com/thomasxnguy/todofinder"
	. "github.com/thomasxnguy/todofinder/app"
	. "github.com/thomasxnguy/todofinder/error"
)

var configurationFile = "../conf/todofinder_test.yaml"
var errorFile = "../conf/errors.yaml"
var ch1 = make(chan struct{})

// init is the global initialization function for the itest package.
// This function will load the test config file and run an instance of the service.
// Configuration file is located to 'configurationFile' path.
func init() {
	config, err := LoadConfiguration(&configurationFile)
	if err != nil {
		fmt.Printf("[ERROR] Couldn't load configuration file for integration testing : %v", err)
	}
	InitTodoFinder(config)
}

// InitTodoFinder run an instance of TodoFinder on port 8089 for testing purpose.
func InitTodoFinder(config *Configuration) {
	if err := LoadMessages(errorFile); err != nil {
		fmt.Printf("[ERROR] Couldn't load error file for integration testing : %v", err)
	}
	logger, err := LoadAppLogger(config)
	if err != nil {
		fmt.Printf("[ERROR] Couldn't load app logger for integration testing : %v", err)
	}

	server := &Server{}
	if err = server.Init(config, logger); err != nil {
		fmt.Printf("[ERROR] Couldn't initialize todofinder: %v", err)
	}

	go func() {
		if rerr := server.Run(); rerr != nil {
			fmt.Printf("[ERROR] Couldn't run todofinder: %v", err)
		}
		//ch1 <- struct{}{}
	}()
}