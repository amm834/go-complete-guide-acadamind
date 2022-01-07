package main

import "fmt"

func main() {
	websites := map[string]string{
		"google": "https://google.com",
		"github": "https://github.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google"])

	websites["fb"] = "https://fb.com"
	fmt.Println(websites)

	delete(websites, "fb")
	fmt.Println("Deleted Map", websites)
}
