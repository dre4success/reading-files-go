package main

import (
	"log"
	"os"
	"github.com/dre4success/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("../blogposts/posts"))
	if err != nil {
		log.Fatal("logging fatal err", err)
	}
	log.Println(posts)
}
