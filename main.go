package main

import (
	"container/list"
	"fmt"
)

func main() {

	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	// //* Performing I/O operations with the "os" and "io/ioutil" packaegs
	// //! Opening a file
	// bs, err := ioutil.ReadFile("test.txt")
	// if err != nil {
	// 	return
	// }
	// str := string(bs)
	// fmt.Println(str)

	// //! Creating a new file
	// file, err := os.Create("hello.txt")
	// if err != nil {
	// 	return
	// }
	// defer file.Close()
	// file.WriteString("Hello my people, how are you doing ?")

	// //! Opening a directory
	// dir, err := os.Open(".")
	// if err != nil {
	// 	return
	// }
	// defer dir.Close()

	// fileInfos, err := dir.ReadDir(-1)
	// if err != nil {
	// 	return
	// }
	// for _, fi := range fileInfos {
	// 	fmt.Println(fi.Name())
	// }

}
