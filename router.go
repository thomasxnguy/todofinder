package todofinder

import (
	"github.com/valyala/fasthttp"
	"strings"
	"encoding/json"
	"os"
	. "github.com/thomasxnguy/todofinder/error"
	. "github.com/thomasxnguy/todofinder/http"
	. "github.com/thomasxnguy/todofinder/app"
)

// SearchResult HTTP Result.
type SearchResultHttpResponse struct {
	Result []SearchResult `json:"result"`
}


// Router route the request to a specific handler.
// Because the application is very simple here, we don't need to use external libraries.
func Router(ctx *fasthttp.RequestCtx) *Error {
	InitRequest(ctx)
	path := strings.Split(string(ctx.Path()), "/")
	switch len(path) {
	case 2:
		switch path[1] {
		case "search":
			switch {
			case ctx.IsGet():
				{
					if err := validateSearchHandler(ctx); err != nil {
						return err
					}
					return searchHandler(ctx)
				}
			default:
				{
					return &Error{METHOD_NOT_ALLOWED, Params{"method": string(ctx.Method())}, nil}
				}
			}
		}
	}
	return &Error{NOT_FOUND, Params{"resource": string(ctx.Path())}, nil}
}

// validateSearchHandler is the request validator for /search endpoint
func validateSearchHandler(ctx *fasthttp.RequestCtx) *Error {
	if !ctx.QueryArgs().Has("package") {
		return &Error{BAD_PARAMETER, Params{"parameter": "package"}, nil}
	} else if !ctx.QueryArgs().Has("pattern") {
		return &Error{BAD_PARAMETER, Params{"parameter": "pattern"}, nil}
	}
	return nil
}

// validateSearchHandler is the request handler for /search endpoint
func searchHandler(ctx *fasthttp.RequestCtx) *Error {
	searchResponse := SearchResultHttpResponse{}
	packageName := string(ctx.QueryArgs().Peek("package"))
	pattern := string(ctx.QueryArgs().Peek("pattern"))
	dir, err := os.Getwd()
	if err != nil {
		return &Error{INTERNAL_SERVER_ERROR, nil, err}
	}
	if p, error := ImportPkg(packageName, dir); error == nil {
		rch := make(chan *SearchResult, 10)
		go ExtractPattern(p, pattern, rch)
		for {
			searchResult := <-rch
			if searchResult == nil {
				break
			} else if searchResult.GetError() != nil {
				return searchResult.GetError()
			} else {
				searchResponse.Result = append(searchResponse.Result, *searchResult)
			}
		}
	} else {
		return error
	}
	return JsonMarshal(ctx, searchResponse)
}

// JsonMarshal set the response body content with the interface v converted into JSON.
func JsonMarshal(ctx *fasthttp.RequestCtx, v interface{}) *Error {
	ctx.SetContentType(CONTENT_TYPE_JSON)
	err := json.NewEncoder(ctx).Encode(v)
	if err != nil {
		return &Error{INTERNAL_SERVER_ERROR, nil, err}
	}
	return nil
}
