// +build ignore

package main

import (
	"log"

	"github.com/shurcooL/vfsgen"

	"gitlab.com/bboehmke/sunny/cmd/monitor/assets"
)

func main() {
	err := vfsgen.Generate(assets.Assets, vfsgen.Options{
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "Assets",
		Filename:     "./assets/assets_vfsdata.go",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
