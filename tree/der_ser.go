package tree

import (
	"fmt"
	"strconv"
	"strings"
)

func deser(treeString string) *TreeNode {
	fields := strings.FieldsFunc(treeString, func(ch rune) bool {
		return ch == ','
	})

	var idx int
	return buildTree(fields, &idx)
}

func buildTree(fields []string, idx *int) *TreeNode {
	index := *idx
	if fields[index] == "-" {
		return nil
	}

	atoi, err := strconv.Atoi(fields[index])
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: atoi}
	*idx = index + 1
	root.Left = buildTree(fields, idx)
	*idx = index + 1
	root.Right = buildTree(fields, idx)

	return root
}

func ser(root *TreeNode) string {
	var stringBuffer strings.Builder
	rser(root, stringBuffer)
	return strings.Trim(stringBuffer.String(), ",")
}

func rser(root *TreeNode, sBuilder strings.Builder) {
	if root == nil {
		sBuilder.WriteString("-")
		return
	}
	sBuilder.WriteString(fmt.Sprintf("%d,", root.Val))
	rser(root.Left, sBuilder)
	rser(root.Right, sBuilder)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
