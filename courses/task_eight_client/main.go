package main

import (
	"bytes"
    "bufio"
	"fmt"
	"encoding/json"
	"log"
	"os"
	"net/http"
	"github.com/Jagrmi-C/gostarted/courses/lib"
	"strings"
)


func main() {
	url := "http://127.0.0.1:3001/increase"

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')

		if strings.Contains(text, "exit") {
			log.Fatal("You stopped application!")
		}

		lib.OnErrFail(err)

		jsonRequest := lib.Test_struct{
			ID:     text,
		}

		// use json body for training goals (simple -> text/plain)
		buf := new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(jsonRequest)

		lib.OnErrFail(err)

		req, _ := http.NewRequest("POST", url, buf)
		client := &http.Client{}
		httpData, err := client.Do(req)

		lib.OnErrFail(err)

		message, err := bufio.NewReader(httpData.Body).ReadString('\n')
		
		lib.OnErrFail(err)

		fmt.Print("Message from server: "+message, "\n")
	}
}
