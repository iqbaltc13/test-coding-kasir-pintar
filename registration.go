package main

// import the 2 modules we need
import (
	"fmt"
	"io/ioutil"
)

func main() {
	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile("params_registration.json")
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("response_registration.json", data, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}

	// if it was successful in reading the file then
	// print out the contents as a string

}
