package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println("Новый проект")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	fmt.Println(*city, *format)

	r := strings.NewReader("Привет,я поток!")
	b := make([]byte, 4)

	for {
		_, err := r.Read(b)
		fmt.Printf("%q\n", b)
		if err == io.EOF {
			break
		}

	}
}
