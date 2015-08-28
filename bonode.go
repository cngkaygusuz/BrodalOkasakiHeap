package BrodalOkasakiHeap
import "fmt"


/*
This is the structure that denotes a single node of the Brodal-Okasaki heap.

I liberally added new fields to this struct. The original description does not need that much information per node.
 */
type BONode struct {
    // This is the value that a node has.
    value			int

    // Subqueue mechanism implements "data-structure-bootstrapping". While merging, children_head field is cleared
    // and moved to subqueue_head.
    subqueue_head	*BONode

    // Children of a node is held in a doubly-linked-list fashion. The parent has a reference to the head of the list.
    children_head	*BONode

    // Every children has a reference to her parent. Leftsibling is the previous, rightsibling is the next element.
    parent 			*BONode
    rightsibling	*BONode
    leftsibling		*BONode

    // Rank of a node.
    rank			int
}


/*
Create a new node.
 */
func newBONode(value int) *BONode {
    return &BONode {
        value: value,

        subqueue_head: nil,
        children_head: nil,

        parent: nil,
        rightsibling: nil,
        leftsibling: nil,

        rank: 0,
    }
}


/*
A node adopts the other node, being its parent.
Here we adjust the necessary fields accordingly to make this connection.
 */
func (bon* BONode) adopt(other *BONode) {

    // Parent relations
    other.parent = bon

    // Sibling relations
    bon.putNodeAmongChildren(other)
}


/*
Inserting the newly-parented node into the doubly-linked-list of children.
 */
func (bon* BONode) putNodeAmongChildren(other *BONode) {
    var prev *BONode
    var next *BONode

    prev = nil
    next = bon.children_head

    for next != nil && other.rank > next.rank {
        prev = next
        next = next.rightsibling
    }

    if prev == nil && next == nil {
        // This means children list is empty.
        bon.children_head = other
    } else if prev == nil && next != nil {
        // We get to this state when there are only 1 child on the list, and our new node has smaller rank than
        // the existing one.
        bon.children_head = other

        other.rightsibling = next
        next.leftsibling = other
    } else if prev != nil && next == nil {
        // We got to the end of the list, our new node has the highest rank.
        prev.rightsibling = other
        other.leftsibling = prev
    } else if prev != nil && next != nil {
        // Standard case, we hit somewhere in between the list.
        prev.rightsibling = other
        next.leftsibling = other

        other.leftsibling = prev
        other.rightsibling = next
    }
}


/*
A node goes rogue, severing its ties with its parent and siblings.
*/
func (bon* BONode) rogue() {
    if bon.parent == nil {
        return  // If no parent, this node hasn't been adopted yet. No need to go through.
    }
    parent := bon.parent

    if parent.children_head == bon {
        parent.children_head = bon.rightsibling  // This can set "nil" to parent.children_head
    } else {
        bon.leftsibling.rightsibling = bon.rightsibling
        if bon.rightsibling != nil {
            bon.rightsibling.leftsibling = bon.leftsibling
        }
    }
    bon.parent = nil
    bon.leftsibling = nil
    bon.rightsibling = nil
}


/*
Essentially same functionality as "rogue", but this one works for subqueued children.
 */
func (bon* BONode) rogue_subqueue() {
    if bon.parent == nil {
        return
    }
    parent := bon.parent

    if parent.subqueue_head == bon {
        parent.subqueue_head = bon.rightsibling
    } else {
        bon.leftsibling.rightsibling = bon.rightsibling
        if bon.rightsibling != nil {
            bon.rightsibling.leftsibling = bon.leftsibling
        }
    }
    bon.parent = nil
    bon.leftsibling = nil
    bon.rightsibling = nil

}

/*
Essential skew-linking procedure is described for "BOHeap.insert_skew"

This function simply performs the linking procedure given the nodes.
 */
func skewLink(firnode *BONode, secnode *BONode, newnode *BONode) *BONode {
    if firnode == nil || secnode == nil {
        // This happens when the parent has less than 2 children.
        return newnode
    } else if firnode.rank != secnode.rank {
        // Inequal ranks. We are going to just simply insert the new node.
        return newnode
    } else {
        // Here is the skew linking.
        // Get the minimum valued node among those 3, make her parent and the other two her children.
        currRank := firnode.rank  // We can also use secnode to get this value.

        minnode, node1, node2 := min_of_3(firnode, secnode, newnode)

        minnode.rogue()
        node1.rogue()
        node2.rogue()

        minnode.adopt(node1)
        minnode.adopt(node2)
        minnode.rank = currRank + 1

        return minnode
    }
}


/*
This function links the two nodes according to binomilal heap linking procedure. The description is given for
"BOHeap.insert_binomial"

The returning node is assumed to be rogue for code-simplifying reasons.
 */
func simpleLink(existingnode *BONode, newnode *BONode) *BONode {
    existingnode.rogue()

    if existingnode.value < newnode.value {
        existingnode.adopt(newnode)
        existingnode.rank += 1
        return existingnode
    } else {
        newnode.adopt(existingnode)
        newnode.rank += 1
        return newnode
    }
}


/*
Return the minimum-valued child.
 */
func (bon* BONode) getMinChild() *BONode {
    if !bon.hasChildren() {
        return nil
    }

    minchild := bon.children_head
    checknode := minchild.rightsibling

    for checknode != nil {
        if checknode.value < minchild.value {
            minchild = checknode
        }
        checknode = checknode.rightsibling
    }

    return minchild
}


func (bon* BONode) hasChildren() bool {
    return bon.children_head != nil
}


/*
Return the smallest ranked 2 children of a given node.

The doubly-linked-list of children is rank ordered, so we don't need to do any searching.
 */
func (bon* BONode) getSmallestRankChildren() (*BONode, *BONode) {
    if bon.children_head == nil {
        return nil, nil
    } else {
        return bon.children_head, bon.children_head.rightsibling
    }
}


/*
Simple.
 */
func (bon* BONode) getSameRankChild(rank int) *BONode {
    child := bon.children_head
    for child != nil {
        if child.rank == rank {
            return child
        }
        child = child.rightsibling
    }
    return nil
}


/*
Print a single node.
 */
func (bon* BONode) print_singular() {
    /*
        fmt.Printf("Node: %d, Rank: %d, address: %p left: %p right: %p parent: %p child_head: %p",
            bon.value, bon.rank, bon, bon.leftsibling, bon.rightsibling ,bon.parent, bon.children_head)
    */
    var sqstr string

    if bon.subqueue_head != nil {
        sqstr = "has subqueue"
    } else {
        sqstr = ""
    }

    str := fmt.Sprintf("Node: %d, Rank: %d %s", bon.value, bon.rank, sqstr)
    fmt.Print(str)
}


/*
Print a node and all of its children recursively in depth-first fashion.
 */
func (bon* BONode) print_recursive(level int) {
    printSpace(level)
    bon.print_singular()
    fmt.Println()

    children := bon.children_head
    for children != nil {
        children.print_recursive(level+4)
        children = children.rightsibling
    }
}


/*
Minimum of 3 algorithm.

First returning value is the minimum, other two are the rest.

I decided to return all of them in a sense, because if we don't do that, we need to do additional work back in the
calling function to determine which one was the smallest.
 */
func min_of_3(n1 *BONode, n2 *BONode, n3 *BONode) (*BONode, *BONode, *BONode) {
    if n1.value < n2.value {
        if n1.value < n3.value {
            return n1, n2, n3
        } else {  // n3 <= n1 <= n2
            return n3, n1, n2
        }
    } else {  // n2 <= n1 ? n3
        if n2.value < n3.value {
            return n2, n1, n3
        } else {  // n3 <= n2 <= n1
            return n3, n1, n2
        }
    }
}


/*
Put subqueue elements to a container, and return them.
I use this to iterate over elements. We can't trust traversing the children-linked-list because the operations that
may be performed can modify this list, hence we may not traverse all of the elements.
 */
func (bon* BONode) subqueueIterator() []*BONode {
    itercont := make([]*BONode, 0, 8)

    child := bon.subqueue_head
    for child != nil {
        itercont = append(itercont, child)
        child = child.rightsibling
    }
    return itercont
}


/*
Read subqueueIterator. Same thing for the children.
 */
func (bon* BONode) childrenIterator() []*BONode {
    itercont := make([]*BONode, 0, 8)

    child := bon.children_head
    for child != nil {
        itercont = append(itercont, child)
        child = child.rightsibling
    }
    return itercont
}


func (bon* BONode) moveChildrenToSubqueue() {
    // There is a bug here. Care to find out?

    bon.subqueue_head = bon.children_head
    bon.children_head = nil
}


// ====== Container helpers ======
func appendList(container []*BONode, head *BONode) []*BONode {
    node := head
    newc := container
    for node != nil {
        newc = append(newc, node)
        node = node.rightsibling
    }
    return newc
}
