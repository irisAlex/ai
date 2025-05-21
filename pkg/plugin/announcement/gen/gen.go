package main

import (
	"path/filepath" //go:generate go mod tidy

	"gorm.io/gen"
	//go:generate go mod download
	//go:generate go run gen.go
	"github.com/irisAlex/ai/pkg/plugin/announcement/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "announcement", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(
		new(model.Info),
	)
	g.Execute()
}
