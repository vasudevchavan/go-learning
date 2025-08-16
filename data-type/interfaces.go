package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type botEngish struct{}
type botSpanish struct{}

// Define a common interface for both bot types
type pBot interface {
	getGreeting() string
}

func (eb botEngish) getGreeting() string {
	return "Hello, friend!"
}

func (sb botSpanish) getGreeting() string {
	return "Hola, amigo!"
}

func (eb botEngish) printGreeting() {
	fmt.Println(eb.getGreeting())
}

func (sb botSpanish) printGreeting() {
	fmt.Println(sb.getGreeting())
}

// Function that accepts the common interface
func cPrintGreeting(b pBot) {
	fmt.Println(b.getGreeting())
}

type Response struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"body"`
	Status     string `json:"status"`
}

func main() {
	eb := botEngish{}
	sb := botSpanish{}

	eb.printGreeting()
	sb.printGreeting()

	// Using the common interface
	cPrintGreeting(eb)
	cPrintGreeting(sb)

	// New example of using the interface

	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}

	// check once.
	// bs := make([]byte, 999999)
	// resp.Body.Read(bs)
	// fmt.Println("Response Body:", string(bs))

	io.Copy(os.Stdout, resp.Body)

	fmt.Println("Response body ", resp.Body)

	// bodyBytes, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }

	// // fmt.Println("Response Body:", string(bodyBytes))

	// readRespondJson := Response{
	// 	StatusCode: resp.StatusCode,
	// 	Body:       string(bodyBytes),
	// 	Status:     resp.Status,
	// }

	// dataJson, err := json.Marshal(readRespondJson)
	// if err != nil {
	// 	fmt.Println("Error marshalling response:", err)
	// 	return
	// }
	// fmt.Println("Response JSON:", string(dataJson))

	defer resp.Body.Close()
}
