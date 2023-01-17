package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
        "os"
        "strings"
)

func main() {
        access_token := os.Getenv("LINE_NOTIFY_ACCESS_TOKEN")

	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	formField, err := writer.CreateFormField("message")
	if err != nil {
		log.Fatal(err)
	}
        
        // メッセージを定義
        input := os.Args
        var message string

        // 引数がある場合
        if len(input) > 1 {
            message = strings.Join(input[1:], " ")
        // 引数がない場合
        } else {
            message = "From Go"
        }
	_, err = formField.Write([]byte(message))

	writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", form)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+access_token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", bodyText)
}
