// 参考: https://golangtokyo.github.io/codelab/find-gophers/?index=codelab#6
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
)

func main() {
	// ファイルごとのトークンの位置を記録するFileSetを作成
	fset := token.NewFileSet()

	// ファイル単位で構文解析を行う
	f, err := parser.ParseFile(fset, "_gopher.go", nil, 0)
	if err != nil {
		log.Fatal("Error:", err)
	}

	log.Println(f)

	ast.Inspect(f, func(n ast.Node) bool {
		fmt.Println("============")
		t := reflect.TypeOf(n)
		fmt.Printf("node: %s\n", n)
		fmt.Printf("nodeType: %s\n", t)
		switch node := n.(type) {
		case *ast.Ident:
			if node.Name != "Gopher" {
				return true
			}
			fmt.Println((fset.Position(node.Pos())))
		case *ast.AssignStmt:
			fmt.Println((fset.Position(node.Pos())))
		default:
			return true
		}

		fmt.Println("============")

		return true
	})
}
