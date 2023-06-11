package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Print(n *Node) {
	if n.Data != "" {
		if n.Name != "" {
			fmt.Printf("%s%s: %s\n", strings.Repeat("\t", int(n.Depth)-1), n.Name, n.Data)
		} else {
			fmt.Printf("%s%s\n", strings.Repeat("\t", int(n.Depth)-1), n.Data)
		}
	}
	if len(n.Kids) != 0 {
		if n.Data == "" && n.Name != "" {
			fmt.Printf("%s%s:\n", strings.Repeat("\t", int(n.Depth)-1), n.Name)
		}
		for _, k := range n.Kids {
			Print(k)
		}
	}
}

func main() {
	var n *Node = &Node{}
	pay, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	err = json.Unmarshal(pay, &n)
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	fmt.Printf("%#v\n", n)
	Print(n)
}
