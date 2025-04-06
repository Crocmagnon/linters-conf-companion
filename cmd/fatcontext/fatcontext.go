package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"linters-conf-companion/fatcontext"
)

func main() {
	singlechecker.Main(fatcontext.Analyzer)
}
