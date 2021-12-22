package main

import "container/list"
type TreeNode struct{
	Value int
	Left *TreeNode
	Right *TreeNode
}
func postorderTraversal(root *TreeNode) []int {
	if root==nil{
		return []int{}
	}
	stack:= list.New()
	res:=[]int{}
	pre:=&TreeNode{}
	//迭代的方式。
	for stack.Len()!=0||root!=nil{
		for root!=nil{
			stack.PushBack(root)
			root=root.Left
		}
		root=stack.Remove(stack.Back()).(*TreeNode)
		//当一个节点没有右子树，或者它的右子树已经遍历过了，那么就可以输出这个节点了
		if root.Right==nil||pre==root.Right{
			res= append(res, root.Value)
			//那么现在的pre就是刚才放进去的那个root了。
			pre=root
			root=nil
		}else{
			stack.PushBack(root)
			root=root.Right
		}
	}
	return res
}
