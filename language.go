package main

type SentencesList struct {
	PathRoot             string
	Dates                []int
	Files                string
	CurrentFileIndex     int
	CurrentSentenceIndex int
	CurrentDocument      TgtDocument
	CurrentSentence      Sentence
}

type TgtDocument struct {
	Sentences []Sentence `xml:"body>"`
}

type Sentence struct {
	Words []Word `xml:"S>"`
}

type Word struct {
	Word     string `xml:",chardata"`
	Id       int    `xml:"ID,attr"`
	Parent   int    `xml:"DOM,attr"`
	Features string `xml:"FEAT,attr"`
	Lemma    string `xml:"LEMMA,attr"`
}
