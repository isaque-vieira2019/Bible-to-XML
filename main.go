package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	nameFile := "Biblia"
	writeConvertFile(readTxt(nameFile), nameFile)
}

func readTxt(nameFile string) []string {
	file, err := ioutil.ReadFile("txt/" + nameFile + ".txt")

	if err != nil {
		log.Fatal("ERROR ao Abrir o arquivo. Message:" + err.Error())
	}

	text := string(file)

	rows := strings.Split(text, "\n")

	/*f, err := os.Create("txtConvert/" + nameFile + ".txt")

	if err != nil {
		log.Fatal("ERROR ao Criar o arquivo. Message:" + err.Error())
	}

	defer f.Close()
	*/
	var convertString []string

	for _, row := range rows {
		_, content, _ := strings.Cut(row, "@")
		convertString = append(convertString, content)
		//f.WriteString(content)
	}

	return convertString
}

func writeConvertFile(txt []string, nameFile string) {
	f, err := os.Create("txtConvert/" + nameFile + ".xml")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var livroAtual = ""
	var livroCount = 0
	var capAtual = ""
	var versCount = 0

	for _, row := range txt {
		livro, text, _ := strings.Cut(row, "@")

		if livro != livroAtual {
			var text string

			if livroAtual != "" {
				text += "\t\t\t\t</CHAPTER>\n"
				text += "\t\t</BIBLEBOOK>\n"
			}

			livroAtual = livro
			livroCount++
			capAtual = ""
			text += "\t\t" + `<BIBLEBOOK bnumber="` + strconv.Itoa(livroCount) + `" bname="- ` + livroAtual + `">` + "\n"
			f.WriteString(text)
		}

		cap, text2, _ := strings.Cut(text, "@")

		if cap != capAtual {
			var text string

			if capAtual != "" {
				text += "\t\t\t\t</CHAPTER>\n"
			}

			capAtual = cap
			versCount = 0
			text += "\t\t\t\t" + `<CHAPTER cnumber="` + capAtual + `">` + "\n"
			f.WriteString(text)
		}

		versCount++
		vers, _, _ := strings.Cut(text2, "@")

		versText := "\t\t\t\t\t\t" + `<VERS vnumber="` + strconv.Itoa(versCount) + `, KJV Fiel 1611">`
		versText += vers
		versText += " </VERS>\n"
		f.WriteString(versText)
	}
}
