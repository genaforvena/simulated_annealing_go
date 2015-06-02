package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func TestWord(t *testing.T) {
	xmlFile, err := os.Open("test_sentence.xml")
	if err != nil {
		t.Fatal("Unable to open file with test sentence!")
	}
	defer xmlFile.Close()

	XMLdata, _ := ioutil.ReadAll(xmlFile)

	var parsedXml WholeXmlText
	xml.Unmarshal(XMLdata, &parsedXml)

	t.Log(parsedXml)

	if parsedXml.TgtDocument.Sentences == nil {
		t.Error("There are no sentences!")
	}

	sentences := parsedXml.TgtDocument.Sentences
	firstSentence := sentences[0]

	if len(firstSentence.Words) != 10 {
		t.Error("Invalid number of sentences " +
			strconv.Itoa(len(firstSentence.Words)))
	}

	word2 := firstSentence.Words[1]

	if word2.Id != 2 {
		t.Error("Invalid word2 id: " + strconv.Itoa(word2.Id))
	}

	if word2.Parent != 0 {
		t.Error("Invalid word2 parent: " + strconv.Itoa(word2.Parent))
	}

	if word2.Features != "V НЕСОВ ИЗЪЯВ НЕПРОШ ЕД 3-Л" {
		t.Error("Invalid word2 features: " + word2.Features)
	}

	if word2.Word != "кажется" {
		t.Error("Invalid word2 word: " + word2.Word)
	}

	if word2.Lemma != "КАЗАТЬСЯ" {
		t.Error("Invalid word2 lemma: " + word2.Lemma)
	}
}
