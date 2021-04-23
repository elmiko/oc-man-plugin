package main

import (
	_ "embed"

    "bytes"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

//go:embed help.md
var helpBytes []byte
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

func getTopicContent(title string, topics TopicsList) bytes.Buffer {
    var buf bytes.Buffer
    switch title {
    case "topics":
        fmt.Fprintln(&buf, "I know about the following topics:")
        for _, t := range topics.Topics {
            fmt.Fprintln(&buf, "*", t.Title)
        }
    case "help":
        fmt.Fprintln(&buf, string(helpBytes))
    default:
        fmt.Fprintln(&buf, "I don't have a page for", title)
    }

    return buf
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("What manual page do you want?")
		fmt.Println("For example, try 'oc man topics'")
	} else {
		topics := getTopics()
		title := os.Args[1]
        content := getTopicContent(title, topics)
        content.WriteTo(os.Stdout)
	}
}
