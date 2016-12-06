package main

import "fmt"

/**
Node nthoeunth
*/
type Node struct {
	value int
	right *Node
	left  *Node
}

func main() {
	head := generateTree()

	fmt.Printf("Starting preorder ")
	preOrderTraversal(&head)
	fmt.Println()

	fmt.Printf("Starting inorder ")
	inOrderTraversal(&head)
	fmt.Println()
}

// 5! = 5 * 4 * 3 * 2 * 1
func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}

func postOrderTraversal(current *Node) {
	if current == nil {
		return
	}

	inOrderTraversal(current.left)
	inOrderTraversal(current.right)
	fmt.Printf("%v,", current.value)
}

func inOrderTraversal(current *Node) {
	if current == nil {
		return
	}

	fmt.Printf("%v,", current.value)
	inOrderTraversal(current.left)
	inOrderTraversal(current.right)
}

// current = head
// inOrderTraversal(current) // 5
//   -> prints 5
//   -> inOrderTraversal(current.left)
//         -> prints 3
//         -> inOrderTraversal(current.left)
//             -> prints 6
//             -> inOrderTraversal(current.left)
//                 nil return
//             -> inOrderTraversal(current.right)
//                 nil return
//         -> inOrderTraversal(current.right)
//             -> prints 8
//             -> inOrderTraversal(current.left)
//                 -> prints 10
//                 -> inOrderTraversal(current.left)
//                     nil return
//                 -> inOrderTraversal(current.right)
//                     nil return
//             -> inOrderTraversal(current.right)
//                 nil return

func preOrderTraversal(current *Node) {
	if current == nil {
		return
	}

	preOrderTraversal(current.left)
	fmt.Printf("%v,", current.value)
	preOrderTraversal(current.right)
}

func generateTree() Node {
	head := Node{value: 5}

	n3 := Node{value: 3}
	n6 := Node{value: 6}
	n8 := Node{value: 8}
	n10 := Node{value: 10}
	n7 := Node{value: 7}
	n9 := Node{value: 9}

	head.left = &n3
	head.right = &n7

	n3.left = &n6
	n3.right = &n8

	n8.left = &n10

	n7.left = &n9

	return head
}
