package main

func main() {
	buffer := bytes.NewBufferString("")
	numBytes, err := buffer.WriteString("Hello World")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Wrote %v bytes: %c\n", numBytes, buffer)
	}
}