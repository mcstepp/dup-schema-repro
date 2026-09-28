package main

import (
	"fmt"
	"os"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/bundler"
	"github.com/pb33f/libopenapi/datamodel"
)

func main() {
	spec, _ := os.ReadFile("openapi.yaml")
	config := datamodel.NewDocumentConfiguration()
	config.BasePath = "."
	config.ExtractRefsSequentially = true
	doc, err := libopenapi.NewDocumentWithConfiguration(spec, config)
	if err != nil {
		panic(err)
	}
	model, err := doc.BuildV3Model()
	if err != nil {
		panic(err)
	}
	out, err := bundler.BundleDocumentComposed(&model.Model, nil)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(out))
}