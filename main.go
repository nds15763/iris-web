package main

import (
	"strconv"

	"../TBEEGO/HongShaoRouFrameWork"
	"../TBEEGO/Route"
	"github.com/kataras/iris"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	app := iris.New()
	// lt1 := ListNode{2, nil}
	// lt := ListNode{4, &lt1}
	// lm := ListNode{3, &lt}
	// lh := ListNode{1, &lm}
	lr := singleNumber([]int{2, 2, 1})
	println(lr)
	//numJewelsInStones()
	//logger，但是不知道在哪
	app.Logger().SetLevel("debug")

	//注册模板
	app.RegisterView(iris.HTML("./Views", ".html"))
	HongShaoRouFrameWork.DbInit()
	//注册路由HongShaoRouFrameWork
	// .Init(app)
	Route.RouteInit(app)

	//配置地址
	app.Run(iris.Addr(":8081"), iris.WithConfiguration(iris.Configuration{
		TimeFormat: "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:    "UTF-8",
	}))
}
func singleNumber(nums []int) int {
	c := strconv.Itoa(int(nums[:]))
	for _, v := range nums {
		if string.Count(v, c) {
			return v
		}
	}
	return 0
}

// func numJewelsInStones() {
// 	var wg sync.WaitGroup
// 	wg.Add(10)
// 	go wgFunc(&wg)
// 	wg.Wait()
// }

// func wgFunc(wg *sync.WaitGroup) {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			fmt.Println(err) //这里的err其实就是panic传入的内容，bug
// 		}
// 	}()
// 	sli := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	for i := 0; i < 10; i++ {
// 		println(sli[i])
// 	}
// 	wg.Done()
// }

// func findKthLargest(nums []int, k int) int {

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] < nums[j] {
// 				nums[i], nums[j] = nums[j], nums[i]
// 			}
// 		}
// 	}
// 	return nums[k]
// }

// func sortList(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	var tail *ListNode
// 	cur := head
// 	for tail != cur.Next {
// 		for cur.Next != tail {
// 			if cur.Val > cur.Next.Val {
// 				cur.Val, cur.Next.Val = cur.Next.Val, cur.Val
// 			}
// 			cur = cur.Next
// 		}
// 		tail = cur //下一次遍历的尾结点是当前结点
// 		cur = head //遍历起始结点重置为头结点
// 	}
// 	return head
// }

// func quickSort(values []int, left, right int) {

// 	temp := values[left]
// 	p := left
// 	i, j := left, right

// 	for i <= j {
// 		for j >= p && values[j] >= temp {
// 			j--
// 		}
// 		if j >= p {
// 			values[p] = values[j]
// 			p = j
// 		}

// 		for i <= p && values[i] <= temp {
// 			i++
// 		}
// 		if i <= p {
// 			values[p] = values[i]
// 			p = i
// 		}

// 	}

// 	values[p] = temp

// 	if p-left > 1 {
// 		quickSort(values, left, p-1)
// 	}
// 	if right-p > 1 {
// 		quickSort(values, p+1, right)
// 	}

// }

// func QuickSort(values []int) {
// 	quickSort(values, 0, len(values)-1)
// }

// func gowg(wg *sync.WaitGroup, i int) {
// 	println(i)
// 	fmt.Println(time.Now())
// 	wg.Done()
// }
// func isBigger(front string, back string) bool {
// 	var flist []byte = []byte(front)
// 	var blist []byte = []byte(back)

// 	//位数不相等
// 	if len(flist) != len(blist) {
// 		return false
// 	}
// 	//按字节比较
// 	for i, v := range flist {
// 		if int(v) > int(blist[i]) {
// 			return true
// 		}
// 	}
// 	return false
// }

// func twoSum(nums []int, target int) []int {
// 	tMap := make(map[int]int)
// 	for i, v := range nums {
// 		tarI := target - nums[i]
// 		if tarValue, ok := tMap[tarI]; ok {
// 			//存在
// 			return []int{tarValue, i}
// 		}
// 		tMap[v] = i
// 	}
// 	return nil
// }
// func removeDuplicates(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	i := 0
// 	for j := 1; j < len(nums); j++ {
// 		if nums[i] != nums[j] {
// 			i++
// 			nums[i] = nums[j]
// 		}
// 	}
// 	return i + 1
// }

// // func popSort(nums []int) []int {
// // 	for i := 0; i < len(nums)-1; i++ {
// // 		for j := 0; j < len(nums)-1-i; j++ {
// // 			if nums[j] > nums[j+1] {
// // 				swap(nums[j], nums[j+1])
// // 			}
// // 		}
// // 	}

// // }
// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
// // type ListNode struct {
// // 	Val  int
// // 	Next *ListNode
// // }

// // type Object interface{}
// // type Node struct {
// // 	data Object
// // 	next *Node
// // }

// // type List struct {
// // 	size uint64 // 车辆数量
// // 	head *Node  // 车头
// // 	tail *Node  // 车尾
// // }

// // func (list *List) Init() {
// // 	(*list).size = 0   // 此时链表是空的
// // 	(*list).head = nil // 没有车头
// // 	(*list).tail = nil // 没有车尾
// // }
// // func (list *List) Append(node *Node) bool {
// // 	if node == nil {
// // 		return false
// // 	}
// // 	(*node).next = nil
// // 	// 将新元素放入单链表中
// // 	if (*list).size == 0 {
// // 		(*list).head = node
// // 	} else {
// // 		oldTail := (*list).tail
// // 		(*oldTail).next = node
// // 	}

// // 	// 调整尾部位置，及链表元素数量
// // 	(*list).tail = node // node成为新的尾部
// // 	(*list).size++      // 元素数量增加

// // 	return true
// // }
// //结点
// type Node struct {
// 	Value       int
// 	Left, Right *Node
// }

// //打印结点值
// func (node *Node) Print() {
// 	fmt.Print(node.Value, " ")
// }

// //设置结点值
// func (node *Node) SetValue(value int) {
// 	if node == nil {
// 		fmt.Println("setting value to nil " +
// 			"node.Ignored.")
// 		return
// 	}
// 	node.Value = value
// }

// //前序遍历
// func (node *Node) preOrder() {
// 	if node == nil {
// 		return
// 	}
// 	node.Print()
// 	node.Left.preOrder()
// 	node.Right.preOrder()
// }

// //中序遍历
// func (node *Node) middleOrder() {
// 	if node == nil {
// 		return
// 	}
// 	node.Left.middleOrder()
// 	node.Print()
// 	node.Right.middleOrder()
// }

// //后序遍历
// func (node *Node) postOrder() {
// 	if node == nil {
// 		return
// 	}
// 	node.Left.postOrder()
// 	node.Right.postOrder()
// 	node.Print()
// }

// //创建结点
// func CreateNode(value int) *Node {
// 	return &Node{Value: value}
// }
