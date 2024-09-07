package sqlite

import (
	"embed"
	"fmt"

	"github.com/midir99/sqload"
)

//go:embed migrations
var fsys embed.FS

var Q = sqload.MustLoadFromFS[struct {
	CreateUserTable   string `query:"CreateUserTable"`
	CreateUserProfile string `query:"CreateUserProfile"`
}](fsys)

func CreateTables() {
	fmt.Println("Queries are: ", Q.CreateUserProfile)
}
