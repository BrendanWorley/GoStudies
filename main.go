package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// creating counter 
	counter := 0

	// reading files in the directory and assigning the array 'files'

	files, err := ioutil.ReadDir("..")
	if err != nil {
		log.Fatal(err)
	}

	// going through the array 'files' printing their names and counting the number of files

	for _, file := range files {

		fmt.Println(file.Name())
		counter = counter +1
		//arrayFiles[i] = file.Name()

	}

	// posting the counted number of files

	fmt.Printf("Number of files is %d.\n", counter)

	// creating an empty array 'filesList' with the length equal to the 'counter'

	filesList := [14]string{}

	// going throu the directory of files ('files' array) and copying files' names into the 'filesList'

	for i, file := range files {
		//fmt.Println(file.Name())
		filesList[i] = file.Name()	
	}

	// posting the 'filesList' and its length

	fmt.Printf("Here's the file list: %s\n", filesList)
	fmt.Printf("The length of the file list is %d\n", len(filesList))

	// creatign the empty slice 'filesSlice'

	filesSlice := []string{}

	// again going through the same files directory and copying the files' names into the slice 'filesSlice'

	for _, file := range files {

		filesSlice = append(filesSlice, file.Name())	

	}

	// posting the slice 'filesSlice' created and its length

	fmt.Printf("And here's the Slice we've made: %s\n", filesSlice)
	fmt.Printf("The length of the Slice created is: %d\n", len(filesSlice))

}