package main

import (
	"github.com/tereus-project/tereus-transpiler-std/core"
	"github.com/tereus-project/tereus-transpiler-template/transpiler"
)

func main() {
	core.InitTranspiler(&core.TranspilerContextConfig{
		SourceLanguage:              "la",
		SourceLanguageFileExtension: ".la",
		TargetLanguage:              "lb",
		TargetLanguageFileExtension: ".lb",
		TranspileFunction:           transpiler.Transpile,
	})
}
