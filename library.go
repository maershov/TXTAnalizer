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

	libHandler("wordsO.txt", wordsO, 2)
	libHandler("wordsA.txt", wordsA, 2)
	libHandler("wordsYA.txt", wordsYA, 2)
	libHandler("words2scl.txt", words2scl, 2)
	libHandler("wordsMZ.txt", wordsMZ, 2)
	libHandler("wordsNarechie.txt", wordsNarechie, 0)
	libHandler("wordsPredlogi.txt", wordsPredlogi, 0)
	libHandler("wordsSouzi.txt", wordsSouzi, 0)
	libHandler("wordsE.txt", wordsE, 2)
	libHandler("wordsL.txt", wordsL, 0)
	libHandler("wordsII.txt", wordsII, 4)
	libHandler("wordsIII.txt", wordsIII, 4)
	libHandler("wordsOI.txt", wordsOI, 4)
	libHandler("wordsMestoimenia.txt", wordsMest, 0)
	libHandler("wordsChastitsa.txt", wordsChast, 0)
	libHandler("wordsNoscl.txt", wordsNoscl, 0)

}
