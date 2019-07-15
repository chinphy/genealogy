package util

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
)

// ClanNode 节点
type ClanNode struct {
	ID       uint        `json:"id"`
	Name     string      `json:"name"`
	Avatar   string      `json:"avatar"`
	Children []*ClanNode `json:"children"`
}

// AddChild 添加节点
func (srv *ClanNode) AddChild() *ClanNode {
	b := make([]byte, 1)
	rand.Read(b)

	// rand.Seed(time.Now().UnixNano())

	node := &ClanNode{
		ID:   uint(b[0]),
		Name: "xx",
	}
	srv.Children = append(srv.Children, node)

	return node
}

// Gen something do
func Gen() {
	root := ClanNode{
		ID:   0,
		Name: "root",
	}
	root.AddChild()
	root.AddChild()
	first := root.AddChild()

	first.AddChild()
	first.AddChild()
	first.AddChild()

	v, _ := json.Marshal(root)

	fmt.Println(string(v))

}
