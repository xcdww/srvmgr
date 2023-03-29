package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	//"github.com/joho/godotenv"
)

var TOKEN string = "5803513827:AAE7jEV8JkUiDCFY-OXH-MuzmbcPF_dFi7Y"
var CHAT_ID string = "6244882472"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("start server!")

	/*
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	*/
	//log_setting
	logFile := openLogFile()
	defer logFile.Close()
	log.SetOutput(logFile)

	mux := http.NewServeMux()
	mux.HandleFunc("/srvmgr/alert", alertHandler)
	mux.HandleFunc("/srvmgr/stat", statHandler)

	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println(err)
	}
}

func alertHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("Could not read body: %s\n", err)
	}
	fmt.Printf("%s: got request", body)
	sendMessage(string(body))
	//io.WriteString(w, "POST\n")
	w.Write([]byte(body))
}

func statHandler(w http.ResponseWriter, req *http.Request) {
	cmd := exec.Command("go", "version")

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	check(err)

	//w.Write(outb.Bytes())
}

func getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", TOKEN)
}

func sendMessage(text string) (bool, error) {
	var err error
	var response *http.Response

	url := fmt.Sprintf("%s/sendMessage", getUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id": CHAT_ID,
		"text":    text,
	})
	response, err = http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	check(err)

	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	check(err)

	log.Printf("Sent Message: '%s'", text)
	log.Printf("Response JSON: '%s'", string(body))

	return true, nil
}

func openLogFile() *os.File {
	f, err := os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
