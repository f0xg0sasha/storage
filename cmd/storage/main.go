package main

import (
	"fmt"

	"github.com/f0xg0sasha/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("work it", st)
}
