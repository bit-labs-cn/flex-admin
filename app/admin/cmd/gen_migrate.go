package cmd

import (
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"text/template"
)

var tpl = `
// generate by auto_migrate Do not edit it
package model

import "gorm.io/gorm"

func Migrate(db *gorm.DB) {
	_ = db.Migrator().AutoMigrate(
		{{range .}}
			&{{.}},
		{{end}}
	)
}
`
var GenMigrate = &cobra.Command{
	Use: "gen:migrate",
	Run: func(cmd *cobra.Command, args []string) {
		var models []string
		dir := "./app/admin/repository/model" // 指定要扫描的目录
		files, err := filepath.Glob(filepath.Join(dir, "*.go"))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, file := range files {
			fset := token.NewFileSet()
			node, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
			if err != nil {
				fmt.Printf("Failed to parse %s: %v\n", file, err)
				continue
			}

			// 遍历 AST 节点以找到所有的结构体定义
			ast.Inspect(node, func(n ast.Node) bool {
				typeDecl, ok := n.(*ast.GenDecl)
				if !ok || typeDecl.Tok != token.TYPE {
					return true
				}
				for _, spec := range typeDecl.Specs {
					typeSpec := spec.(*ast.TypeSpec)
					models = append(models, typeSpec.Name.Name+"{}")
					fmt.Println("Found struct:", typeSpec.Name.Name)
					// 处理每个结构体
					//handleStruct(structType, pkg.Name, typeSpec.Name.Name)
				}
				return true
			})
		}
		parse, err := template.New("model").Parse(tpl)
		if err != nil {
			return
		}
		var b bytes.Buffer
		err = parse.Execute(&b, models)
		if err != nil {
			return
		}
	},
}
