package main

import (
	"fmt"
	"github.com/heroyan/go-cmd/cb"
	"os"

	"github.com/c-bata/go-prompt"
	_ "github.com/heroyan/go-cmd/cb"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func ppt()  {
	for {
		fmt.Println("Please select table.")
		t := prompt.Input("> ", completer)
		fmt.Println("You selected " + t)
		if t == "quit" || t == "exit" {
			os.Exit(0)
		}
	}
}

func chTest()  {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}

func cmdTest()  {
	cb.Execute()
}

func main() {
	cmdTest()
}