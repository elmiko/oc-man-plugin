package main

import (
    "bytes"
	"embed"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

//go:embed content/*
var contentFiles embed.FS
//go:embed help.md
var helpBytes []byte

type TopicsList struct {
	Titles []string `yaml:"titles"`
}

func getTopics() TopicsList {
	var topics TopicsList

    content, err := contentFiles.ReadFile(fmt.Sprintf("content/index.yaml"))
    if err != nil {
		log.Fatalf("cannot find title index file: %v", err)
    }

	if err := yaml.Unmarshal(content, &topics); err != nil {
		log.Fatalf("cannot unmarshal topics data: %v", err)
	}

	return topics
}

func getTopicContent(title string, topics TopicsList) bytes.Buffer {
    var buf bytes.Buffer
    switch title {
    case "topics":
        fmt.Fprintln(&buf, "I know about the following topics:")
        for _, t := range topics.Titles {
            fmt.Fprintln(&buf, "*", t)
        }
    case "help":
        fmt.Fprintln(&buf, string(helpBytes))
    default:
        content, err := contentFiles.ReadFile(fmt.Sprintf("content/%s", title))
        if err != nil {
            fmt.Fprintln(&buf, "I don't have a page for", title)
            fmt.Fprintln(&buf, err)
        } else {
            fmt.Fprintln(&buf, string(content))
        }
    }

    return buf
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("What manual page do you want?")
		fmt.Println("For example, try 'oc man topics' or 'oc man help'")
	} else {
		topics := getTopics()
		title := os.Args[1]
        content := getTopicContent(title, topics)
        content.WriteTo(os.Stdout)
	}
}
