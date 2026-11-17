package main

import(
	"os"
	"fmt"
)

func main() {
	if err:= run(); err != nil{
		fmt.Fprintln(os.Stderr,"erreur:" , err)
		os.Exit(1)
	}
}
