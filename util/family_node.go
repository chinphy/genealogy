package util

import (
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

// FamilyNode 家族树处理
type FamilyNode struct {
	LeafIsFound bool                   // 子节点是否找到
	LinkPath    *lls.Stack             // 子节点到根节点的路径链
	tree        map[string]interface{} // 树结构
}

// SetTree 绑定树结构
func (FN *FamilyNode) SetTree(tree map[string]interface{}) {
	FN.tree = tree
}

// FindPath 找出节点到根节点的路径
func (FN *FamilyNode) FindPath(path *lls.Stack, toFind int) *lls.Stack {
	//stack := lls.New()
	if FN.tree == nil {
		return FN.LinkPath
	}

	path.Push(FN.tree["id"])

	if FN.LeafIsFound == false && FN.tree["id"] == toFind {
		FN.LeafIsFound = true
		FN.LinkPath = FN.PrintPath(path)
	}

	// if FN.tree["children"] != nil {
	// 	m := make(map[string]interface{})
	// 	m = FN.tree["children"]
	// 	for _, xx := range m {
	// 		FN.FindPath(path, toFind)
	// 	}
	// }

	return FN.LinkPath
}

// PrintPath 打印和组装成新数据
func (FN *FamilyNode) PrintPath(path *lls.Stack) *lls.Stack {
	stack := lls.New()

	for i := 0; i < path.Size(); i++ {
		val, _ := path.Pop()
		stack.Push(val)
	}

	return stack
}
