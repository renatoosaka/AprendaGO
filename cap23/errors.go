package main

import (
	"os"
)

func main() {
	// f, err := os.Create("log.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer f.Close()

	// log.SetOutput(f)

	_, errr := os.Open("file.txt")
	if errr != nil {
		// fmt.Println(err)
		// log.Println(errr)
		// log.Fatalln(errr)
		// log.Panicln(errr)
		// panic(errr)
	}
}
