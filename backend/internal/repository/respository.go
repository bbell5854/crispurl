package repository

import "fmt"

func buildTableName(tablename, env string) string {
	return fmt.Sprintf("%s_%s", tablename, env)
}
