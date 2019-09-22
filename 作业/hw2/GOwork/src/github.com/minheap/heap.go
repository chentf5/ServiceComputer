package main

import "fmt"

type Node struct {
	Value int
}


// 用于构建结构体切片为最小堆，需要调用down函数
func Init(nodes []Node) {
    var currentsize int = len(nodes)
    var pos int = (currentsize-1)/2
    for pos >= 0   {
        down(nodes,pos,currentsize-1)
        pos = pos-1
    }
    
}

// 需要down（下沉）的元素在切片中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i, n int) {
    j := 2*i +1
    temp := nodes[i]
    for j <= n    {
        if j < n && nodes[j].Value > nodes[j+1].Value {
            j++
		}
		
        if nodes[j].Value >= temp.Value {
            break
        } else {
            nodes[i] = nodes[j]
            i = j
            j = 2*i+1
        }
    }
    nodes[i] = temp
}

// 用于保证插入新元素(j为元素的索引,切片末尾插入，堆底插入)的结构体切片之后仍然是一个最小堆
func up(nodes []Node, j int) {
    i := (j-1) /2
    temp := nodes[j]
    for j > 0 {
        if nodes[i].Value <= temp.Value {
            break
        } else {
            nodes[j] = nodes[i]
            j = i
            i = (j-1)/2
        }
    }
    nodes[j] = temp
}

// 弹出最小元素，并保证弹出后的结构体切片仍然是一个最小堆，第一个返回值是弹出的节点的信息，第二个参数是Pop操作后得到的新的结构体切片
func Pop(nodes []Node) (Node, []Node) {
	var zero Node
	zero.Value = 0
    if len(nodes) == 0  {
        return zero,nodes
    }
    min_X := nodes[0]
    nodes = nodes[1:]
    down(nodes,0,len(nodes)-1)
    return min_X,nodes 

}

// 保证插入新元素时，结构体切片仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) []Node {
    
    nodes = append(nodes,node)
    up(nodes,len(nodes)-1)
    return nodes
}

// 移除切片中指定索引的元素，保证移除后结构体切片仍然是一个最小堆
func Remove(nodes []Node, node Node) []Node {
    i := 0
    for i < len(nodes)    {
        if node.Value == nodes[i].Value {
            break
		}
		i++
	}
	//fmt.Println(i)
	if i != len(nodes)	{
		nodes = append(nodes[:i],nodes[i+1:]...)
	}
	
	//fmt.Println(nodes)
    Init(nodes)
    return nodes
}

func main() {
	var node1 Node
	node1.Value = 53
	var node2 Node
	node2.Value = 17
	var node3 Node
	node3.Value = 78
	var node4 Node
	node4.Value = 9
	var node5 Node
	node5.Value = 45
	var node6 Node
	node6.Value = 65
	var node7 Node
	node7.Value = 87
	var node8 Node
	node8.Value = 23
	var nodes []Node
	nodes = append(nodes,node1,node2,node3,node4,node5,node6,node7,node8)
	fmt.Println(nodes)
	Init(nodes)
	fmt.Println(nodes)
	var temp Node
	temp.Value = 11
	
	node_temp := Push(temp,nodes)
	fmt.Println(node_temp)
	
	
	node_temp = Remove(node_temp,temp)
	fmt.Println(node_temp)
    //fmt.Println(nodes)
    v ,node_pop := Pop(nodes)
    fmt.Println(v,node_pop)
}