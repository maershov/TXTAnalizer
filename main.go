package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"regexp"
)

var flagPadej = flag.Bool("p", false, "flag for padej forms accounting")
var flagHist = flag.Bool("h", false, "flag for histogram")

func FilePosition() int {
	var counter int = 0

	if *flagPadej {
		counter++
	}

	if *flagHist {
		counter++
	}

	return counter
}

func checkNarechie(str string) bool {

	if wordsNarechie[str] {
		return true
	} else {
		return false
	}
}

func checkPredlog(str string) bool {

	if wordsPredlogi[str] {
		return true
	} else {
		return false
	}
}

func checkSouz(str string) bool {

	if wordsSouzi[str] {
		return true
	} else {
		return false
	}
}

func onlyLetters(str string) string {
	re, err := regexp.Compile(`[^а-яА-ЯёЁ]`)

	if err != nil {
		log.Fatal(err)
	}

	str = re.ReplaceAllString(str, "")
	//fmt.Println(str1)
	return str
}

func counterWorker(str string, mapa *map[string]int) {
	_, ok := (*mapa)[str]

	if ok {
		(*mapa)[str] += 1
	} else {
		(*mapa)[str] = 1
	}
}

func fileReader(input io.Reader, output io.Writer) error {
	var err error
	in := bufio.NewScanner(input)
	in.Split(bufio.ScanWords)
	in.Buffer(nil, 1024)
	counter := make(map[string]int)

	for in.Scan() {
		txt := in.Text()
		txt = onlyLetters(txt)
		txt = strings.ToLower(txt)

		if *flagPadej {

			if !checkNarechie(txt) && !checkPredlog(txt) {
				if !checkSouz(txt) {
					err, txt = padej(txt)
				}
				counterWorker(txt, &counter)

				if err != nil {
					fmt.Println(err)
					return err
				}
			} else {
				//fmt.Println("deleted narechie or predlog")
			}
		} else {
			counterWorker(txt, &counter)
		}
	}

	var maxValue = 0

	for index, value := range counter {
		if value > maxValue {
			maxValue = value
		}
		fmt.Println(index, "->", value)
	}

	if *flagHist {
		leng := len(counter)
		err := Histogram(counter, leng, maxValue)

		if err != nil {
			fmt.Println(err)
			return err
		}

	}

	return nil
}

func main() {
	flag.Parse()
	filePos := FilePosition()
	file, err := os.Open(os.Args[filePos+1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	createLibrary()

	err = fileReader(file, os.Stdout)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

}
