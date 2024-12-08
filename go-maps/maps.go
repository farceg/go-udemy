package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://www.google.com",
		"AWS":    "https://aws.amazon.com"}

	fmt.Println(websites)
	fmt.Println(websites["AWS"])

	websites["LinkedIn"] = "https://www.linkedin.com/feed/"

	delete(websites, "AWS")

	fmt.Println(websites["AWS"])
	fmt.Println(websites)
}
