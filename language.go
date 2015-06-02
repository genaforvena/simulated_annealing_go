package main

import (
	"encoding/xml"
)

type WholeXmlText struct {
	XMLName     xml.Name    `xml:"text"`
	TgtDocument TgtDocument `xml:"body"`
}

type TgtDocument struct {
	XMLName   xml.Name   `xml:"body"`
	Sentences []Sentence `xml:"S"`
}

type Sentence struct {
	XMLName xml.Name `xml:"S"`
	Words   []Word   `xml:"W"`
}

type Word struct {
	XMLName  xml.Name `xml:"W"`
	Word     string   `xml:",chardata"`
	Id       int      `xml:"ID,attr"`
	Parent   int      `xml:"DOM,attr"`
	Features string   `xml:"FEAT,attr"`
	Lemma    string   `xml:"LEMMA,attr"`
}
