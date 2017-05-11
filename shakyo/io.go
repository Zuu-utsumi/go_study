package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Readln(reader *bufio.Reader) (string, error) {
	var (
		isPrefix    bool  = true
		err         error = nil
		line, lines []byte
	)

	// NOTE: Reading line is limited 4096byte. left parts of the line can be get
	// 			 next read. At that time isPrefix bool will be true.
	for isPrefix && err == nil {
		line, isPrefix, err = reader.ReadLine()
		lines = append(lines, line...)
	}

	return string(lines), err
}

func openFile(path string) *os.File {
	f, e := os.OpenFile(path, os.O_RDONLY, 0644)

	if e != nil {
		log.Fatal(e)
	}

	return f
}

func main1() {
	path := "./io.go"
	// f, e := os.Open(path)
	f := openFile(path)

	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		l, err := Readln(reader)

		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			log.Fatal(err)
		}

		fmt.Println(l)
	}
}

func newFile(path string) *os.File {
	f, e := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if e != nil {
		log.Fatal(e)
	}

	return f
}

func main2() {
	path := "./tmp.txt"
	// f, e := os.Create(path)
	f := newFile(path)

	defer f.Close()
	// defer os.Remove(path)

	writer := bufio.NewWriter(f)
	str := "hogehogeogehogagaeg aaaaaa"
	_, err := writer.WriteString(str)

	if err != nil {
		log.Fatal(err)
	}

	// release buffer memory
	writer.Flush()
}

func main3() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println("input text>>" + scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main4() {
	path := "./tmp.txt"
	fp, _ := os.Create(path)
	bw := bufio.NewWriter(fp)
	bw.WriteString("Hello world!")
	bw.Flush()
}

func main5() {
	main4()
	bs, err := ioutil.ReadFile("./tmp.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}

func main() {
	// main1()
	// main2()
	// main3()
	// main4()
	main5()
}
