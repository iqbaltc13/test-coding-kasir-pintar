package main

// import the 2 modules we need
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile("params_payment.txt")
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}

	url := "https://dev.nicepay.co.id/nicepay/direct/v2/payment?"

	//request method
	method := "POST"

	var jsonData = []byte(data)
	url += string(jsonData)
	//fmt.Println(url)
	client := &http.Client{}
	req, errRequest := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	// set HTTP request header Content-Type
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if errRequest != nil {
		fmt.Println(errRequest)
		return
	}
	res, errResponse := client.Do(req)
	if errResponse != nil {
		fmt.Println(errResponse)
		return
	}
	defer res.Body.Close()

	body, errReadBody := ioutil.ReadAll(res.Body)
	if errReadBody != nil {
		fmt.Println(errReadBody)
		return
	}
	//fmt.Println(string(body))

	err = ioutil.WriteFile("response_payment.json", body, 0777)
	// handle this error
	if err != nil {
		// print it out
		fmt.Println(err)
	}

	// if it was successful in reading the file then
	// print out the contents as a string

}
