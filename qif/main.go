package main

import (
	"fmt"
	"os"

	"github.com/pacificbrian/qif"
)

func main() {

	var file, err = os.Open("./data/export.qif")

	if err != nil {
		panic("Failed to read file " + err.Error())
	}
	qifFile := qif.NewReader(file)

	fmt.Println(qifFile.Read())
}
