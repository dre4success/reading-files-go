package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

func newPost(postfile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postfile)

	readMetaLine := func(tagname string) string {
		scanner.Scan()
		log.Println(scanner.Text())
		return strings.TrimPrefix(scanner.Text(), tagname)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagSeparator), ", ")

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        readBody(scanner),
	}, nil

}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
