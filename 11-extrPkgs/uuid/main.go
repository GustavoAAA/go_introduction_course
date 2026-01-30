package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	p := fmt.Println
	id1 := uuid.New()
	p(id1)

	id2 := uuid.NewString()
	p(id2)

	uid := uuid.New()
	p(uid.Version())
	p(uid.String())
}
