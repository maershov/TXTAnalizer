package main

import (
	"bufio"
	"fmt"
	"os"
)

var wordsO = make(map[string]bool)
var wordsA = make(map[string]bool)
var wordsYA = make(map[string]bool)
var words2scl = make(map[string]bool)
var wordsMZ = make(map[string]bool)
var wordsNarechie = make(map[string]bool)
var wordsPredlogi = make(map[string]bool)
var wordsSouzi = make(map[string]bool)
var wordsE = make(map[string]bool)
var wordsL = make(map[string]bool)
var wordsII = make(map[string]bool)
var wordsIII = make(map[string]bool)
var wordsOI = make(map[string]bool)
var wordsMest = make(map[string]bool)
var wordsChast = make(map[string]bool)
var wordsNoscl = make(map[string]bool)

func libHandler(path string, mapa map[string]bool, delta int){
	file, err := os.Open(path)

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	in := bufio.NewScanner(file)
	in.Split(bufio.ScanWords)
	in.Buffer(nil, 1024)
	for in.Scan() {
		txt := in.Text()
		txt = onlyLetters(txt)
		fillLib(txt[:len(txt) - delta], &mapa)
	}

}

func fillLib(str string, mapa* map[string]bool)  {
	_, ok := (*mapa)[str]

	if !ok {
		(*mapa)[str] = true
	}

}

func createLibrary() {

	libHandler("./librarywords/wordsO.txt", wordsO, 2)
	libHandler("./librarywords/wordsA.txt", wordsA, 2)
	libHandler("./librarywords/wordsYA.txt", wordsYA, 2)
	libHandler("./librarywords/words2scl.txt", words2scl, 2)
	libHandler("./librarywords/wordsMZ.txt", wordsMZ, 2)
	libHandler("./librarywords/wordsNarechie.txt", wordsNarechie, 0)
	libHandler("./librarywords/wordsPredlogi.txt", wordsPredlogi, 0)
	libHandler("./librarywords/wordsSouzi.txt", wordsSouzi, 0)
	libHandler("./librarywords/wordsE.txt", wordsE, 2)
	libHandler("./librarywords/wordsL.txt", wordsL, 0)
	libHandler("./librarywords/wordsII.txt", wordsII, 4)
	libHandler("./librarywords/wordsIII.txt", wordsIII, 4)
	libHandler("./librarywords/wordsOI.txt", wordsOI, 4)
	libHandler("./librarywords/wordsMestoimenia.txt", wordsMest, 0)
	libHandler("./librarywords/wordsChastitsa.txt", wordsChast, 0)
	libHandler("./librarywords/wordsNoscl.txt", wordsNoscl, 0)

}
