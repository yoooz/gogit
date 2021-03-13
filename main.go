package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {
	r, _ := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: "../test",
	})

	ref, _ := r.Head()
	
	cIter, _ := r.Log(&git.LogOptions{From: ref.Hash()})

	authorMap := map[string][]string{}


	cIter.ForEach(func(c *object.Commit) error {
		authorMap[c.Author.Name] = append(authorMap[c.Author.Name], c.Message)
		return nil
	})

	for key, value := range authorMap {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Println()
	}
}
