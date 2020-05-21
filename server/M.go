package main

import "fmt"

type Tempora struct {
	Hola  int
	Hola2 []string
}

func main() {
	// s, err := goscraper.Scrape("www.truora.com/", 5)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("Icon : %s\n", s.Preview.Icon)
	// fmt.Printf("Name : %s\n", s.Preview.Name)
	// fmt.Printf("Title : %s\n", s.Preview.Title)
	// fmt.Printf("Description : %s\n", s.Preview.Description)
	// fmt.Printf("Image: %s\n", s.Preview.Images[0])
	// fmt.Printf("Url : %s\n", s.Preview.Link)

	var tem Tempora
	tem = Tempora{Hola: 2}
	fmt.Println(tem.Hola2)
}
