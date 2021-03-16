package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"encoding/json"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {
	flag.Parse()
	path := flag.Args()[0]
	r, _ := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: path,
		SingleBranch: true,
		NoCheckout: false,
	})

	ref, _ := r.Head()
	
	cIter, _ := r.Log(&git.LogOptions{From: ref.Hash()})

	commitListMap := map[string][]gitLogStruct{}


	cIter.ForEach(func(c *object.Commit) error {
		message := strings.Fields(c.Message)
		index, _ := strconv.Atoi(message[1])
		log := gitLogStruct{ Hash: message[0], HistoryIndex: index }
		commitListMap[c.Author.Name] = append(commitListMap[c.Author.Name], log)
		return nil
	})

	json, _ := json.Marshal(commitListMap)
	fmt.Println(string(json))
}

type gitLogStruct struct {
	Hash string `json:"hash"`
	HistoryIndex int `json:"historyIndex"`
}
