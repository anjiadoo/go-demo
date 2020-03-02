package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

const topicMapperFilePath = "./topic_mapper_config.xml"

type TopicMapper struct {
	XMLName xml.Name `xml:"Event"`
	Topics  []Topic  `xml:"Topic"`
}

type Topic struct {
	TopicID   string     `xml:"topicID,attr"`
	Topic     string     `xml:"topic,attr"`
	ChanName  string     `xml:"chanName,attr"`
	Desc      string     `xml:"desc,attr"`
	SubTopics []SubTopic `xml:"SubTopic"`
}

type SubTopic struct {
	SubTopic    string `xml:"subTopic,attr"`
	FirstField  string `xml:"firstField,attr"`
	SecondField string `xml:"secondField,attr"`
	Desc        string `xml:"desc,attr"`
}

func NewTopicMapper() *TopicMapper {
	return loadXMLFile()
}

func loadXMLFile() *TopicMapper {
	file, err := ioutil.ReadFile(topicMapperFilePath)
	if err != nil {
		log.Fatalf("load topic_mapper_config.xml fail, err:%v", err)
	}
	var Mapper = &TopicMapper{}
	if err := xml.Unmarshal(file, Mapper); err != nil {
		log.Fatalf("xml file unmarshal fail, err:%v", err)
	}
	log.Print(Mapper)
	return Mapper
}

func main() {
	NewTopicMapper()
}
