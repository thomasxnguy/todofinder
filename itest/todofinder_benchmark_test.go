package itest

import (
	"testing"
	"gopkg.in/h2non/baloo.v3"
)

var tfBalooServerBenchmark = baloo.New("http://localhost:8089")

func Benchmark_WrongEndpoint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tfBalooServerBenchmark.
			Get("/test_wrong")
	}
}

func Benchmark_BadParameters(b *testing.B) {
	queryParam := map[string]string{
		"pattern": "TODO",
	}
	for i := 0; i < b.N; i++ {
		tfBalooServerBenchmark.
			Get("/test_wrong").
			SetQueryParams(queryParam)
	}
}

func Benchmark_BadMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tfBalooServerBenchmark.
			Post("/search")
	}
}

func Benchmark_SearchEndpoint(b *testing.B) {
	queryParam := map[string]string{
		"pattern": "TODO",
		"package": "github.com/m-rec/14d4017ddb43a7c0cb3ab4be9ea18cbc74ee15ab/todofinder",
	}
	for i := 0; i < b.N; i++ {
		tfBalooServerBenchmark.
			Get("/search").
			SetQueryParams(queryParam)
	}
}
