package server

import (
	"flag"
	"fmt"
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder"
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/app"
	. "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder/error"
	"runtime"
	"path"
)

// Todofinder command tool properties for server mode.
const (
	CommandName  = "server"
	Description  = "run in server mode"
	usageMessage = "%s [-config path_to_config] \n"
)

// Options structure for server mode.
type option struct {
	configPath string
	flagSet    *flag.FlagSet
}

//n ewOption initialises the command options for server mode.
func newOption(eh flag.ErrorHandling) (opt *option) {
	opt = &option{
		flagSet: flag.NewFlagSet(CommandName, eh),
	}
	// option settings
	opt.flagSet.StringVar(&opt.configPath, "config", "", "configuration file path")

	return opt
}

// parse validates the flags for server mode.
func (opt *option) parse(args []string) (err error) {
	if err = opt.flagSet.Parse(args); err != nil {
		return
	}
	//Check extra flags
	if nonFlag := opt.flagSet.Args(); len(nonFlag) != 0 {
		return fmt.Errorf("invalid argument: %v", nonFlag)
	}

	if opt.configPath == "" {
		return fmt.Errorf("missing flag: %v", "config")
	}
	return
}

// Usage provides the usage message for server mode.
func Usage() {
	fmt.Printf(usageMessage, CommandName)
}

// PrintDefaults prints the default flags for server mode.
func PrintDefaults(eh flag.ErrorHandling) {
	opt := newOption(eh)
	opt.flagSet.PrintDefaults()
}

// Run executes todofinder in server mode.
// It receives arg from command line and validate them.
// It will then run the command.
func Run(args []string) error {
	opt := newOption(flag.ExitOnError)
	if e := opt.parse(args); e != nil {
		Usage()
		PrintDefaults(flag.ExitOnError)
		return fmt.Errorf("%v, %v", CommandName, e)
	}
	return command(opt)
}

// Execute todofinder in server mode.
func command(opt *option) error {
	config, err := LoadConfiguration(&opt.configPath)
	if err != nil {
		return err
	}
	//error file location can be putted into global configuration file
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Errorf("cannot get file directory")
	}
	errorFilePath := path.Join(path.Dir(filename), "../../conf/errors.yaml")
	err = LoadMessages(errorFilePath)
	if err != nil {
		return err
	}
	logger, err := LoadAppLogger(config)
	if err != nil {
		return err
	}

	server := &Server{}
	if rerr := server.Init(config, logger); rerr != nil {
		return fmt.Errorf("[ERROR] Couldn't run: %v", rerr)
	}
	server.Run()
	return nil
}
