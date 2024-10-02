package main

import "fmt"

func main() {

	//#1 map
	player := map[string]string{
		"username": "ligichi",
		"job":      "katana",
	}

	fmt.Println(player)
	fmt.Println(player["username"])
	fmt.Println(player["job"])

	//#2 func map make, delete

	programmingLang := make(map[string]string)
	programmingLang["diff"] = "mid/high"
	programmingLang["name"] = "GoLang"
	programmingLang["fun"] = "no"

	delete(programmingLang, "fun")

	fmt.Println(programmingLang)

}
