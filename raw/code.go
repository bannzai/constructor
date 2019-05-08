package raw

import "go/ast"

// Code is presentation of .go file content.
type Code struct {
	FilePath Path
	ASTFile  ast.File
}
