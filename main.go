package main

import (
	_ "embed"

	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

//go:embed topics.yaml
var topicsBytes []byte

type TopicsEntry struct {
	Title    string `yaml:"title"`
	Location string `yaml:"location"`
}

type TopicsList struct {
	Topics []TopicsEntry `yaml:"topics"`
}

func getTopics() TopicsList {
	var topics TopicsList

	if err := yaml.Unmarshal(topicsBytes, &topics); err != nil {
		log.Fatalf("cannot unmarshal topics data: %v", err)
	}

	return topics
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("What manual page do you want?")
		fmt.Println("For example, try 'oc man topics'")
	} else {
		topics := getTopics()
		target := os.Args[1]
		if target == "topics" {
			fmt.Println("I know about the following topics:")
			for _, t := range topics.Topics {
				fmt.Println(t.Title)
			}
		} else {
			fmt.Println("I don't have a page for", target)
		}
	}
}
