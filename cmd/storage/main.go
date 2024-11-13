package main

import (
	"fmt"
	"log"

	"github.com/f0xg0sasha/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Update("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("uploaded", file)
}
