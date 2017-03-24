/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var stack []*TreeNode
	depth := make(map[*TreeNode]int)
	depth[root] = 1
	mini := 0x7fffffff
	for len(stack) != 0 || root != nil { //基于栈实现的中序遍历
		if root != nil {
			stack = append(stack, root)
			if root.Left != nil {
				depth[root.Left] = depth[root] + 1
			}
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root.Left == nil && root.Right == nil && mini > depth[root] { // leafNode
				mini = depth[root]
			}
			if root.Right != nil {
				depth[root.Right] = depth[root] + 1
			}
			root = root.Right
		}
	}
	return mini
}
