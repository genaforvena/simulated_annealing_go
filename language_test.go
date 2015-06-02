package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

func TestWord(t *testing.T) {
	xmlFile, err := os.Open("test_sentence.xml")
	if err != nil {
		t.Fatal("Unable to open file with test sentence!")
	}
	defer xmlFile.Close()

	XMLdata, _ := ioutil.ReadAll(xmlFile)

	var tgtDocument TgtDocument
	xml.Unmarshal(XMLdata, &tgtDocument)

	if tgtDocument.Sentences == nil {
		t.Error("There are no words read!")
	}

}
