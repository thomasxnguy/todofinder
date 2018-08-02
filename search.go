package todofinder

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"strings"
	. "github.com/thomasxnguy/todofinder/error"
	"path/filepath"
)

// SearchResult represent one result when parsing file in research of a pattern.
// It collect essential data to display to end-users.
type SearchResult struct {
	FileName string `json:"file"`
	Position int    `json:"pos"`
	Comment  string `json:"com"`
	error    *Error
}

// ToString return a verbose message corresponding to the search result.
func (s *SearchResult) ToString() string {
	return fmt.Sprintf("%s:%v:\n%s\n", s.FileName, s.Position, s.Comment)
}

// GetError return error in result search.
func (s *SearchResult) GetError() *Error {
	return s.error
}

// ImportPkg returns details about the Go package named by the import path,
// interpreting local import paths relative to dir.
// If the path is a local import path naming a package that can be imported
// using a standard import path, the returned package will set p.ImportPath
// to that path.
func ImportPkg(path, dir string) (*build.Package, *Error) {
	//TODO Optimisations can be done here
	p, err := build.Import(path, dir, build.ImportComment)
	if err != nil {
		return nil, &Error{PACKAGE_NOT_FOUND, Params{"package": path}, err}
	}
	if p.IsCommand() {
		return nil, &Error{NO_SOURCE, Params{"source": p.Name}, err}
	}

	return p, nil
}

// ExtractPattern extracts all the comment in all the .go files from a specific package
// and search for a specific pattern. Each occurrence produce a result that is sent in the channel.
// Termination occurs when sending error or nil element in the channel.
func ExtractPattern(p *build.Package, pattern string, resultChannel chan *SearchResult) {
	for _, f := range p.GoFiles {
		fname := filepath.Join(p.Dir, f)
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, fname, nil, parser.ParseComments)
		if err != nil {
			resultChannel <- &SearchResult{"", 0, "", &Error{SOURCE_NOT_READABLE, Params{"source": fname}, err}}
			return
		}
		cmap := ast.NewCommentMap(fset, f, f.Comments)
		for n, cgs := range cmap {
			f := fset.File(n.Pos())
			for _, cg := range cgs {
				t := cg.Text()
				if strings.Contains(t, pattern) {
					resultChannel <- &SearchResult{fname, f.Position(cg.Pos()).Line, t, nil}
				}
			}
		}
	}
	//End of function, nil object will terminate the routine
	resultChannel <- nil
}
