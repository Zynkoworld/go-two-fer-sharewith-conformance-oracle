package main

import (
	__json "encoding/json"
	__fmt "fmt"
)



// ShareWith apportions the goods between two people.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}

type R struct {
	Ok bool        `json:"ok"`
	V  interface{} `json:"v"`
}

func main() {
	inputs := []string{"", "Alice", "Bob"}
	out := []R{}
	for _, x := range inputs {
		func() {
			defer func() { if r := recover(); r != nil { out = append(out, R{false, __fmt.Sprint(r)}) } }()
			out = append(out, R{true, ShareWith(x)})
		}()
	}
	b, _ := __json.Marshal(map[string]interface{}{"out": out})
	__fmt.Println(string(b))
}
