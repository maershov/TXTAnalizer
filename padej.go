package main

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strings"
)

var uniqueSuf1 = []string{"ули", "уле", "улю", "улей", "уле"} //уль килоджоуль
var uniqueSuf2 = []string{"ого", "ому", "ым", "ыми"}          //ый  веселый без ом
var uniqueSuf3 = []string{"ов", "ам", "ами", "ах"}            //ы  компы
var uniqueSuf4 = []string{"ями", "ях", "ям"}                  //и  ноздри
var glagolEnding1 = []string{"у", "ю", "ешь", "ет", "ем", "ете", "ут", "ют"}
var glagolEnding2 = []string{"у", "ю", "ишь", "ит", "им", "ите", "ат", "ю", "ят"}
var glagolEndingPast = []string{"ала", "ал", "ял", "яла", "яли", "ул", "ел", "ёл", "ли", "лся", "лась", "ла", "али", "ало", "ил"}
var glagolEndingSYA = []string{"лся", "лась", "лись"}
var prilagatEnding = []string{"ыми", "ьими", "ьих", "ые", "ие", "ьи", "ых", "их", "ого", "ьего", "его", "ой", "ьей", "ей", "ому", "ьему", "ему", "ым", "ьим", "им", "ые", "ие", "ое", "ее", "ье", "ую", "ью", "ом", "ем", "ая", "ое"}
var soglasnie = []string{"б", "в", "щ", "г", "д", "л", "ж", "з", "к", "н", "п", "т", "ф", "ч", "ц", "щ", "р", "х"}
var shSoglasnie = []string{"ш", "щ", "ч"}
var glasnie = []string{"е", "ё", "о", "а", "ы", "я", "и", "ю"}
var meForms = []string{"меня", "себя", "мне", "себе", "мной", "собой"}
var youForms = []string{"тебя", "тебе", "тобой"}
var heForms = []string{"его", "ему", "нем"}
var sheForms = []string{"ее", "её", "ей", "ней"}
var theyForms = []string{"их", "им", "ими", "них"}
var usForms = []string{"нас", "нам", "нами"}
var allForms = []string{"всех", "всем", "всеми"}

func isMestoimenie(str string) bool {
	if wordsMest[str] {
		return true
	}
	return false
}

func isChastitsa(str string) bool {
	if wordsChast[str] {
		return true
	}
	return false
}

func isNoscl(str string) bool {
	if wordsNoscl[str] {
		return true
	}
	return false
}

func changeHandler(str string, value, mest string) (error, string) {
	//fmt.Println(str)
	re, err := regexp.Compile(value + `$`)
	if err != nil {
		log.Fatal(err)
	}
	str = re.ReplaceAllString(str, mest)
	//fmt.Println(str)
	return nil, str
}

func padej(str string) (error, string) {

	if isChastitsa(str) {
		//fmt.Println(str)
		return nil, str
	}

	if isNoscl(str) {
		//fmt.Println(str)
		return nil, str
	}

	if isMestoimenie(str) {
		for _, value := range meForms {
			if value == str {
				return changeHandler(str, value, "я")
			}
		}
		for _, value := range youForms {
			if value == str {
				return changeHandler(str, value, "ты")
			}
		}
		for _, value := range heForms {
			if value == str {
				return changeHandler(str, value, "он")
			}
		}
		for _, value := range sheForms {
			if value == str {
				return changeHandler(str, value, "она")
			}
		}
		for _, value := range theyForms {
			if value == str {
				return changeHandler(str, value, "они")
			}
		}
		for _, value := range usForms {
			if value == str {
				return changeHandler(str, value, "мы")
			}
		}
		for _, value := range allForms {
			if value == str {
				return changeHandler(str, value, "все")
			}
		}

		return nil, str
	}

	for _, value := range prilagatEnding {
		if strings.HasSuffix(str, value) {
			delta := len(value)
			if wordsII[str[:len(str)-delta]] {
				return changeHandler(str, value, "ий")
			}
			if wordsIII[str[:len(str)-delta]] {
				return changeHandler(str, value, "ый")

			}

			if wordsOI[str[:len(str)-delta]] {
				return changeHandler(str, value, "ой")
			}
		}
	}

	for _, value := range uniqueSuf1 {
		if strings.HasSuffix(str, value) {
			return changeHandler(str, value, "уль")
		}
	}

	for _, value := range uniqueSuf2 { //веселому
		if strings.HasSuffix(str, value) {
			return changeHandler(str, value, "ый")

		}
	}

	for _, value := range uniqueSuf3 { //компами парками
		if strings.HasSuffix(str, value) {
			var err error
			err, str = changeHandler(str, value, "")
			if err != nil {
				//fmt.Print(err)
				return err, ""
			}
		}
	}

	for _, value := range uniqueSuf4 { //ноздри
		if strings.HasSuffix(str, value) {
			//fmt.Println("testim yami  ", value, len(value))
			if wordsE[str[:len(str)-len(value)]] {
				return changeHandler(str, value, "е")

			}
			//fmt.Println(`hello from ` + value)
			return changeHandler(str, value, "и")
		}
	} // закончили проверять уникальные окончания

	if strings.HasSuffix(str, "ой") { // скакалкой

		return changeHandler(str, "ой", "а")

	}

	if strings.HasSuffix(str, "ом") {

		if wordsO[str[:len(str)-4]] {
			return changeHandler(str, "ом", "о")

		}

		if len(str) > 6 {

			return changeHandler(str, "ом", "")

		}
	}

	if strings.HasSuffix(str, "ем") { //решением
		s := []rune(str)
		length := len(s)
		if strings.HasSuffix(str, "енем") {
			return changeHandler(str, "енем", "ый")

		}
		if wordsE[string(s[:length-2])] {
			return changeHandler(str, "м", "")

		}
	}

	if strings.HasSuffix(str, "ью") { //мякотью
		return changeHandler(str, "ю", "")

	}

	if strings.HasSuffix(str, "ей") { // чащей или яблоней
		s := []rune(str)
		length := len(s)
		condition := false
		for _, value := range shSoglasnie {
			if value == string(s[length-3]) {
				condition = true
			}
		}
		if condition {
			return changeHandler(str, "ей", "а")
		} else {
			if words2scl[str[:len(str)-4]] {
				return changeHandler(str, "ей", "ь")

			}
			return changeHandler(str, "ей", "я")

		}
	}
	//проверили сложные окончания, начинаем проверять одиночные

	if strings.HasSuffix(str, "е") {
		s := []rune(str)
		length := len(s)

		if wordsA[str[:len(str)-2]] {
			return changeHandler(str, "е", "а")

		}

		if wordsO[str[:len(str)-2]] {
			return changeHandler(str, "е", "о")

		}

		if wordsYA[str[:len(str)-2]] {
			return changeHandler(str, "е", "я")

		}

		if words2scl[str[:len(str)-2]] {
			return changeHandler(str, "е", "ь")

		}

		conditionGlas := false
		for _, value := range glasnie {
			if value == string(s[length-2]) {
				conditionGlas = true
			}
		}

		if conditionGlas {
			return changeHandler(str, "е", "й")

		}
		return changeHandler(str, "е", "")

	}

	if strings.HasSuffix(str, "ы") {
		if wordsA[str[:len(str)-2]] {
			return changeHandler(str, "ы", "а")
		} else {
			return changeHandler(str, "ы", "")

		}
	}

	if strings.HasSuffix(str, "у") {

		if wordsA[str[:len(str)-2]] {
			return changeHandler(str, "у", "а")

		}

		if wordsO[str[:len(str)-2]] {
			return changeHandler(str, "у", "о")

		}

		return changeHandler(str, "у", "")

	}

	if strings.HasSuffix(str, "и") {

		if wordsA[str[:len(str)-2]] {
			return changeHandler(str, "и", "а")

		}

		if wordsYA[str[:len(str)-2]] {
			return changeHandler(str, "и", "я")

		}

		if wordsMZ[str[:len(str)-2]] {
			return changeHandler(str, "и", "ь")
		}

		if wordsE[str[:len(str)-2]] {
			return changeHandler(str, "и", "е")

		}
		var err error
		err, str = changeHandler(str, "и", "")

		if err != nil {
			return err, ""
		}

	}

	if strings.HasSuffix(str, "а") {

		if wordsA[str[:len(str)-2]] {
			//fmt.Println("hello from wordA", str)
			return nil, str
		}

		if wordsO[str[:len(str)-2]] {
			return changeHandler(str, "а", "о")

		}

		s := []rune(str)
		length := len(s)
		condition := true
		if "л" == string(s[length-2]) {
			if !wordsL[str[:len(str)-2]] {
				//fmt.Println("hello from wordL")
				condition = false
			}
		}
		fmt.Println(str)
		if condition {
			return changeHandler(str, "а", "")

		}
	}
	//fmt.Print("aaa")
	if strings.HasSuffix(str, "ю") {

		if wordsYA[str[:len(str)-2]] {
			return changeHandler(str, "ю", "я")
		}

		s := []rune(str)
		length := len(s)
		condition := false
		for _, value := range glasnie {
			if value == string(s[length-2]) {
				condition = true
			}
		}
		if condition {
			return changeHandler(str, "ю", "й")
		}

		if strings.HasSuffix(str, "ую") {
			return changeHandler(str, "ую", "ий")

		}

		return changeHandler(str, "ю", "")

	}

	if strings.HasSuffix(str, "я") {
		if strings.HasSuffix(str, "ться") || strings.HasSuffix(str, "аяся") || strings.HasSuffix(str, "ая") || strings.HasSuffix(str, "яя") {
			//fmt.Println(str)
			return nil, str
		}

		s := []rune(str)
		length := len(s)
		condition := false
		for _, value := range glasnie {
			if value == string(s[length-2]) {
				condition = true
			}
		}
		if condition {
			return changeHandler(str, "я", "й")

		}

		if wordsMZ[str[:len(str)-2]] {
			return changeHandler(str, "я", "ь")

		}
	}

	for _, value := range glagolEnding1 {
		if strings.HasSuffix(str, value) {
			return changeHandler(str, value, "ть")
		}

		//глаголы
		for _, value := range glagolEnding2 {
			if strings.HasSuffix(str, value) {
				return changeHandler(str, value, "ть")
			}
		}
		//проверим окончания прошедшего времени для глаголов
		for _, value := range glagolEndingPast { //гуляли
			if strings.HasSuffix(str, value) {
				delta := len(value) / 2
				if !(delta%2 == 0) {
					delta--
				}
				if value == "ли" || value == "лся" || value == "лась" || value == "ла" {
					delta = 0
				}
				if value[2:] == "л" {
					if wordsL[str] {
						return nil, str
					}
				}
				//fmt.Println("end  ", str,  value ,value[delta :], delta)
				return changeHandler(str, value[delta:], "ть")

			}
		}

		for _, value := range glagolEndingSYA { //ноздри
			if strings.HasSuffix(str, value) {
				return changeHandler(str, value, "ться")

			}
		}
	}

	for _, value := range soglasnie { //гуляли
		if strings.HasSuffix(str, value) {
			if wordsA[str] {
				return changeHandler(str, "", "а")

			}
			if wordsO[str] {
				return changeHandler(str, "", "о")

			}
		}
	}
	//fmt.Println(str)

	return nil, str

}

func padejTester(str string, input io.Reader, output io.Writer) error {
	err, strin := padej((str))

	if err != nil {
		fmt.Fprint(output, err)
		return err
	}

	fmt.Fprintln(output, strin)
	return nil
}
