package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const (
	TreeSignDash     string = "─"
	TreeSignVertical string = "│"
	TreeSignUpMiddle string = "├"
	TreeSignUpEnding string = "└"
)

type Node struct {
	Data      string
	Name      string
	Kids      []*Node
	Parent    *Node
	Depth     uint
	Collapsed bool
}

func (n *Node) Collapse() {
	n.Collapsed = true
}

func (n *Node) CollapseAll() {
	n.Collapsed = true
	for _, k := range n.Kids {
		k.Collapsed = true
	}
}

func (n *Node) Expand() {
	n.Collapsed = false
}

func (n *Node) ExpandAll() {
	n.Collapsed = false
	for _, k := range n.Kids {
		k.Collapsed = false
	}
}

func (n *Node) From(N any) {
	switch reflect.TypeOf(N).Kind() {
	case reflect.Map:
		for k, v := range N.(map[string]any) {
			switch reflect.TypeOf(v).Kind() {
			case reflect.Map:
				if n.Depth != 0 {
					kid := &Node{
						Name:   k,
						Parent: n,
						Depth:  n.Depth + 1,
					}
					kid.From(v)
					n.Kids = append(n.Kids, kid)
				} else {
					n.Name = k
					n.From(v)
				}
			case reflect.Array:
				if n.Depth != 0 {
					kid := &Node{
						Name:   k,
						Parent: n,
						Depth:  n.Depth + 1,
					}
					kid.From(v)
					n.Kids = append(n.Kids, kid)
				} else {
					n.Name = k
					n.From(v)
				}
			default:
				n.Kids = append(n.Kids, &Node{
					Data:   fmt.Sprint(v),
					Name:   k,
					Parent: n,
					Depth:  n.Depth + 1,
				})
			}
		}
	case reflect.Array:
		for i, v := range N.([]any) {
			switch reflect.TypeOf(v).Kind() {
			case reflect.Map:
				if n.Depth != 0 {
					kid := &Node{
						Name:   fmt.Sprint(i),
						Parent: n,
						Depth:  n.Depth + 1,
					}
					kid.From(v)
					n.Kids = append(n.Kids, kid)
				} else {
					n.Name = fmt.Sprint(i)
					n.From(v)
				}
			case reflect.Array:
				if n.Depth != 0 {
					kid := &Node{
						Name:   fmt.Sprint(i),
						Parent: n,
						Depth:  n.Depth + 1,
					}
					kid.From(v)
					n.Kids = append(n.Kids, kid)
				} else {
					n.Name = fmt.Sprint(i)
					n.From(v)
				}
			default:
				n.Kids = append(n.Kids, &Node{
					Data:   fmt.Sprint(v),
					Name:   fmt.Sprint(i),
					Parent: n,
					Depth:  n.Depth + 1,
				})
			}
		}
	default:
		n.Data = fmt.Sprint(N)
	}
}

func (n *Node) UnmarshalJSON(b []byte) error {
	var m map[string]any
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	n.Data = ""
	n.Depth = 1
	n.Kids = []*Node{}
	n.Parent = nil
	n.Name = ""
	n.From(m)
	return nil
}
