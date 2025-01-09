package main

import "fmt"

func main() {
	printHello()
}

func printHello() {
	fmt.Printf(Hello())
}

func Hello() string {
	return fmt.Sprintf("Hello, %s!\n", "http_server")
}
