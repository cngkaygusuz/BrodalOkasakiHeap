package BrodalOkasakiHeap
import "fmt"


/*
"BOHeap" is a wrapper around "BONode". This structure is defines an entrypoint for a Brodal-Okasaki heap and
implements the priority queue interface using operations defined for "BONode" structure.
 */
type BOHeap struct {
    root	*BONode
    size 	int
}


/*
Create a new Brodal-Okasaki heap.
 */
func NewBOHeap() *BOHeap {
    return &BOHeap {
        root: nil,
        size: 0,
    }
}


/*
Insert a new value to the heap.

This function depends on another "insert" function, but executing this function always performs the insertion procedure
of the skew-binomial heaps, which is simply described as follows:
    1. Get two nodes from the children of root node that has the minimum ranks.
    2. Skew-link those two nodes and the new node.
    3. Insert this newly linked node among the children of root.

Much of the logical complexity is hidden in the "skew-linking" procedure. Read "BONode.skewLink".
 */
func (bq *BOHeap) Insert(value int) {
    newnode := newBONode(value)
    bq.size += 1
    bq.insert(newnode)
}

/*
Delete and return the minimum valued node.

Since we do almost everything under Pop(), this function is a bit complicated. I implemented Pop() as follows:

    1. Get the minimum node among the children of root.
    2. If this node has elements in its subqueue, add them to the original queue.
    3. For every children of the minimum node;
        3a. Insert the child again to the original heap.

All suboperations (1, 2, 3) take O(logn) time. As you can see, Pop() does every operation that takes O(logn) time, hence
the asymptotical optimality.

In the original paper, re-inserting the children of a node involves partitioning the children. I didn't bother understanding
what the authors really meant, and made a slight modification here. Figure out the difference.
 */
func (bq *BOHeap) Pop() int {
    bq.size -= 1
    retval := bq.root.value

    if bq.root.subqueue_head != nil {
        bq.merge_subqueue()
    }

    minchild := bq.root.getMinChild()

    if minchild == nil {
        // minchild == nil signifies the heap is empty.
        bq.root = nil
    } else {
        minchild.rogue()
        bq.reInsertChildren(minchild)
        bq.promoteToRoot(minchild)
    }

    return retval
}


/*
Return the minimum valued element in the queue.
 */
func (bq *BOHeap) Peek() int {
    // Minimum is the global root so we have O(1) access time.
    return bq.root.value
}


/*
Return the total keys present in the queue.
 */
func (bq *BOHeap) Size() int {
    return bq.size
}


/*
Instead of doing the merge operation immediately, we do it when we need it. This is also called "lazy-evaluation".
Implementation of this operation is:
    * Move the children head to the subqueue head.
    * Insert the root of other queue as if it is a singleton node.
 */
func (bq* BOHeap) Merge(other *BOHeap) {
    bq.size += other.size

    oroot := other.root
    oroot.moveChildrenToSubqueue()

    bq.insert(oroot)
}


// ====== Helper functions =======


/*
This method performs the actual insertion operation. It is a bit complicated because it has to deal with 4 cases:
    * Insertion when the queue is empty
    * Insertion when the new node's value is smaller than the root.
    * Insertion of nodes with rank 0
        * We perform the insertion procedure of skew binomial heaps.
    * Insertion of nodes with rank >0
        * We perform the insertion procedure of ordinary binomial heaps.
 */
func (bq* BOHeap) insert(newnode *BONode) {
    if bq.root == nil {
        // When the queue is empty.
        bq.root = newnode
    } else if newnode.value < bq.root.value {  // If you put "<=" instead of "<", an infinite loop occurs. Find why.
        oldroot := bq.swapWithRoot(newnode)
        bq.insert(oldroot)
    } else {
        if newnode.rank > 0 {
            bq.insert_binomial(newnode)
        } else {
            bq.insert_skew(newnode)
        }
    }
}


/*
Merge the subqueue. Essentially, re-insert the immediate children, but for the subqueue.
 */
func (bq* BOHeap) merge_subqueue() {
    for _, node := range bq.root.subqueueIterator() {
        node.rogue_subqueue()
        bq.insert(node)
    }
}

/*
Re-insert the children. Simple.
 */
func (bq* BOHeap) reInsertChildren(bon *BONode) {
    // Reinsert the children of the GIVEN node.
    for _, node := range bon.childrenIterator() {
        node.rogue()
        bq.insert(node)
    }
}


/*
This function performs skew-binomial insertion, which is;
    * Get the minimum ranked 2 children.
    * Compare the ranks of those 2 children
        * If they have the same rank, skew-link them, and simply insert the resulting node.
        * If they have different ranks, simply insert the newnode.
    It is important to note that "newnode" has a rank of 0. This method should not run for nodes that has rank <>0
 */
func (bq* BOHeap) insert_skew(newnode *BONode) {
    node1, node2 := bq.root.getSmallestRankChildren()
    mergednode := skewLink(node1, node2, newnode)
    bq.root.adopt(mergednode)
}


/*
This is the insertion procedure for ordinary binomial heaps, which is;
    * Get a child that has the same rank for the other node.
        * If such node exists, among those two who has smaller value adopts the other one. Insert the newly made node
          into the queue.
        * If not, simply insert the other node.

This operation can possibly take O(logn) time under normal implementation.
Consider the case, inserting a node of rank 3 to the binomial heap that has the following state (numbers imply ranks):

    * (4 3) <-- 3

This will cause another insertion operation

    * (4) <-- 4

In the worst case, this operation cascades up to "logn" time, but this fact doesn't harm the asymptotical optimality,
since it can only be executed under Pop().
 */
func (bq* BOHeap) insert_binomial(other *BONode) {
    // "other" is assumed to be rogue.

    srnode := bq.root.getSameRankChild(other.rank)

    if srnode == nil {
        bq.root.adopt(other)
    } else {
        newnode := simpleLink(srnode, other)
        // There is a recursion here. If the aforementioned cascading to be occur, it shall be done so by recursing
        // the next line of code. Figure it out yourself.
        bq.insert(newnode)
    }
}


/*
Swap the other node with the existing root.
 */
func (bq* BOHeap) swapWithRoot(newroot *BONode) *BONode {
    // "newroot" is assumed to be rogue and childless.
    oldroot := bq.root

    // Since the children has a reference to their parents, we need to iterate through all of them anyway.
    // Simply transferring "children_head" won't work.
    for _, node := range oldroot.childrenIterator() {
        node.rogue()
        newroot.adopt(node)
    }

    oldroot.rank = 0

    bq.root = newroot
    return oldroot
}


/*
Promote the selected node to be the new root.
 */
func (bq* BOHeap) promoteToRoot(minnode *BONode) {
    // There is a bug. What is it?
    bq.root.value = minnode.value
    bq.root.subqueue_head = minnode.subqueue_head
}


func (bq* BOHeap) print() {
    if bq.root == nil {
        fmt.Println("heap is empty.")
    } else {
        bq.root.print_recursive(0)
    }
}