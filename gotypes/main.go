package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

func main() {
	fset := token.NewFileSet()
	userGo, _ := parser.ParseFile(fset, "user/user.go", nil, 0)

	fmt.Println("==== Inspect user.go ====")
	printIdentPos(fset, userGo)

	// type check user/user.go
	conf := types.Config{Importer: importer.ForCompiler(fset, "source", nil)}
	userInfo := types.Info{
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
	}
	_, err := conf.Check("user", fset, []*ast.File{userGo}, &userInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("==== User identifier ====")
	userIdent := findIdentNode(userGo, 194)
	userObj := userInfo.ObjectOf(userIdent)
	fmt.Printf("defined at: %+v\n", fset.Position(userObj.Pos()))

	fmt.Println("==== Send identifier ====")
	sendIdent := findIdentNode(userGo, 219)
	sendObj := userInfo.ObjectOf(sendIdent)
	fmt.Printf("defined at: %+v\n", fset.Position(sendObj.Pos()))
}

func printIdentPos(fset *token.FileSet, f *ast.File) {
	ast.Inspect(f, func(n ast.Node) bool {
		switch n := n.(type) {
		case *ast.Ident:
			fmt.Printf("%+v | %+v: %+v\n", fset.Position(n.Pos()), n, n.Pos())
		}
		return true
	})
}

func findIdentNode(f *ast.File, pos token.Pos) *ast.Ident {
	var result *ast.Ident
	ast.Inspect(f, func(n ast.Node) bool {
		if result != nil {
			return false
		}
		switch n := n.(type) {
		case *ast.Ident:
			if n.Pos() <= pos && pos <= n.End() {
				result = n
			}
		}
		return true
	})
	return result
}
