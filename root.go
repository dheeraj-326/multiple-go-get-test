package get_test

import (
	"github.com/google/uuid"
)

func RootHello() string {
	return "Hello, Root!" + uuid.New().String()
}
