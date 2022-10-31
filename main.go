package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	path := "D:\\Programming\\Go-Project\\1.png"

	resp, err := http.Get("https://www.google.com/logos/fnbx/westminster_dog_show/paw_print_4.png")
	// https://www.google.com/logos/fnbx/westminster_dog_show/paw_print_4.png
	if err != nil {
		log.Fatalln(err)
	}
	// //We Read the response body on the line below.
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // convert []byte datatype to Image datatype
	// img, _, err := image.Decode(bytes.NewReader(body))
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// save an image to file
	f, err := os.Create(path)
	if err != nil {
		panic(err)

	}
	defer f.Close()
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")

	// open windows program...
	// c := exec.Command("cmd", "/C", "del", "D:\\a.txt")
	c := exec.Command("C:\\Program Files (x86)\\Google\\Picasa3\\Picasa3.exe", path)
	// "C:\Program Files (x86)\Google\Picasa3\Picasa3.exe"

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	//Convert the body to type string
	// sb := string(body)
	// log.Printf(body)
}
