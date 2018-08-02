package search

import (
	"flag"
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"github.com/thomasxnguy/todofinder"
	. "github.com/thomasxnguy/todofinder/error"
	"runtime"
	"path"
)

// Todofinder search mode command tool properties.
const (
	CommandName  = "search"
	Description  = "run in command line"
	usageMessage = "%s [-package input_package] [-pattern input_pattern] \n"
)

// Options structure for search mode.
type option struct {
	packageName string
	pattern     string
	flagSet     *flag.FlagSet
}

// newOption initialise command options for search mode.
func newOption(w io.Writer, eh flag.ErrorHandling) (opt *option) {
	opt = &option{
		flagSet: flag.NewFlagSet(CommandName, eh),
	}
	// option settings
	opt.flagSet.SetOutput(w)
	opt.flagSet.StringVar(&opt.packageName, "package", "", "input package")
	opt.flagSet.StringVar(&opt.pattern, "pattern", "", "input pattern")

	return opt
}

// parse validate the flags for search mode.
func (opt *option) parse(args []string) (err error) {
	if err = opt.flagSet.Parse(args); err != nil {
		return
	}
	//Check extra flags
	if nonFlag := opt.flagSet.Args(); len(nonFlag) != 0 {
		return fmt.Errorf("invalid argument: %v", nonFlag)
	}
	//Check mandatory flags
	if opt.packageName == "" {
		return fmt.Errorf("missing flag: %v", "package")
	}

	if opt.pattern == "" {
		return fmt.Errorf("missing flag: %v", "pattern")
	}
	return
}

// OptionCheck receives a slice of args and returns an error if it was not successfully parsed.
func OptionCheck(args []string) (err error) {
	opt := newOption(ioutil.Discard, flag.ContinueOnError)
	if e := opt.parse(args); e != nil {
		return fmt.Errorf("%v, %v", CommandName, e)
	}
	return nil
}

// Usage provides usage message for search mode.
func Usage() {
	fmt.Printf(usageMessage, CommandName)
}

// PrintDefaults prints the default flag for search mode.
func PrintDefaults(eh flag.ErrorHandling) {
	opt := newOption(os.Stderr, eh)
	opt.flagSet.PrintDefaults()
}

// Run Execute todofinder search command.
// It receives arg from command line and validate them.
// It will then run the main command option.
func Run(args []string) error {
	opt := newOption(os.Stderr, flag.ExitOnError)
	if e := opt.parse(args); e != nil {
		Usage()
		PrintDefaults(flag.ExitOnError)
		return fmt.Errorf("%v, %v", CommandName, e)
	}
	return command(opt)
}

// Search main execution program.
func command(opt *option) error {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Errorf("cannot get file directory")
	}
	errorFilePath := path.Join(path.Dir(filename), "../../conf/errors.yaml")
	err := LoadMessages(errorFilePath)
	if err != nil {
		return err
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	if p, error := todofinder.ImportPkg(opt.packageName, dir); error == nil {
		rch := make(chan *todofinder.SearchResult, 10)
		go todofinder.ExtractPattern(p, opt.pattern, rch)
		for {
			searchResult := <-rch
			if searchResult == nil {
				return nil
			} else if searchResult.GetError() != nil {
				return fmt.Errorf("%v, %v", CommandName, searchResult.GetError().GetMessage())
			} else {
				fmt.Printf(searchResult.ToString())
			}
		}
	} else {
		return fmt.Errorf("%v, %v", CommandName, error.GetMessage())
	}
	return nil
}
