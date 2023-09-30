package main

import (
	"fmt"
	"log"

	"github.com/tealeg/xlsx"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	f, err := xlsx.OpenFile(`modules\resources\data.xlsx`)
	check(err)
	file, err := f.ToSlice()
	fmt.Println(file[0][1][0])
}
