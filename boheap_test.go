package BrodalOkasakiHeap


import (
    "testing"
    "math/rand"
)


func Test_adopt_rankorder_1(t *testing.T) {
    root := newBONode(0)

    n1 := newBONode(1)
    n1.rank = 1

    n2 := newBONode(2)
    n2.rank = 2

    n3 := newBONode(3)
    n3.rank = 3

    n4 := newBONode(4)
    n4.rank = 4

    root.adopt(n1)
    root.adopt(n2)
    root.adopt(n3)
    root.adopt(n4)



    child := root.children_head
    if child != n1 {t.Errorf("Expected %d, got %d", n1.value, child.value)}

    child = child.rightsibling
    if child != n2 {t.Errorf("Expected %d, got %d", n2.value, child.value)}

    child = child.rightsibling
    if child != n3 {t.Errorf("Expected %d, got %d", n3.value, child.value)}

    child = child.rightsibling
    if child != n4 {t.Errorf("Expected %d, got %d", n4.value, child.value)}
}


func Test_adopt_rankorder_2(t *testing.T) {
    root := newBONode(0)

    n1 := newBONode(1)
    n1.rank = 1

    n2 := newBONode(2)
    n2.rank = 2

    n3 := newBONode(3)
    n3.rank = 3

    n4 := newBONode(4)
    n4.rank = 4

    n4_2 := newBONode(5)
    n4_2.rank = 4

    root.adopt(n4_2)
    root.adopt(n4)
    root.adopt(n3)
    root.adopt(n2)
    root.adopt(n1)

    child := root.children_head
    if child != n1 {t.Errorf("Expected %d, got %d", n1.value, child.value)}

    child = child.rightsibling
    if child != n2 {t.Errorf("Expected %d, got %d", n2.value, child.value)}

    child = child.rightsibling
    if child != n3 {t.Errorf("Expected %d, got %d", n3.value, child.value)}

    child = child.rightsibling
    if child != n4 {t.Errorf("Expected %d, got %d", n4.value, child.value)}
}


func Test_adopt_rankorder_3(t *testing.T) {
    root := newBONode(0)

    n1 := newBONode(1)
    n1.rank = 0

    n2 := newBONode(2)
    n2.rank = 0

    n3 := newBONode(3)
    n3.rank = 3

    n4 := newBONode(4)
    n4.rank = 1

    root.adopt(n1)
    root.adopt(n2)
    root.adopt(n3)
    root.adopt(n4)

    if !rankOrdered(root) {t.Errorf("not rank ordered")}
}


func Test_smallestrank_1(t *testing.T) {
    root := newBONode(0)

    n1 := newBONode(1)
    n1.rank = 1

    n2 := newBONode(2)
    n2.rank = 1

    n3 := newBONode(3)
    n3.rank = 5

    n4 := newBONode(4)
    n4.rank = 6

    root.adopt(n1)
    root.adopt(n2)
    root.adopt(n3)
    root.adopt(n4)

    m1, m2 := root.getSmallestRankChildren()

    if (m1 != n1) && (m1 != n2) {t.Errorf("false smallest rank children")}
    if (m1 != n2) && (m2 != n2) {t.Errorf("false smallest rank children.")}

}


func Test_insert_1(t *testing.T) {
    heap := NewBOHeap()

    n1 := newBONode(1)
    n2 := newBONode(2)

    heap.insert(n1)
    heap.insert(n2)

    if heap.root != n1 {t.Errorf("wrong root.")}
    if !isChild(heap.root, n2) {t.Errorf("n2 not a child of root.")}

    if n1.children_head != n2 {t.Errorf("relationship error.")}
    if n2.parent != heap.root {t.Errorf("relationship error.")}
}


func Test_insert_2(t *testing.T) {
    heap := NewBOHeap()

    n1 := newBONode(1)
    n2 := newBONode(2)
    n3 := newBONode(3)
    n4 := newBONode(4)

    heap.insert(n1)
    heap.insert(n2)
    heap.insert(n3)
    heap.insert(n4)

    if heap.root != n1 {t.Errorf("wrong root.")}
    if !isChild(n2, n3) {t.Errorf("relationship error.")}
    if !isChild(n2, n4) {t.Errorf("relationship error.")}

    if n2.parent != n1 {t.Errorf("relationship error.")}
    if n3.parent != n2 {t.Errorf("relationship error.")}
    if n4.parent != n2 {t.Errorf("relationship error.")}
}


func Test_pop_1(t *testing.T) {
    heap := NewBOHeap()

    heap.Insert(1)
    heap.Insert(2)
    heap.Insert(3)
    heap.Insert(4)

    if heap.Pop() != 1 {t.Errorf("wrong pop value")}
}


func Test_rankorder_1 (t *testing.T) {
    const SIZE = 30

    heap := NewBOHeap()
    arr := interval(0, SIZE)
    insert_mult(heap, arr)

    for i:=0; i<SIZE; i++ {
        heap.Pop()
        if !rankOrdered(heap.root) {t.Errorf("not rank ordered")}
    }
}


func Test_heapsort_1(t *testing.T) {
    heap := NewBOHeap()

    for i:=30; i>=0; i-- {
        heap.Insert(i)
    }

    for i:=0; i<=30; i++ {
        pval := heap.Pop()
        if pval != i {t.Errorf("expected %d, got %d", i, pval)}
    }
}


func Test_heapsort_missing_1(t *testing.T) {
    const SIZE = 25
    rand.Seed(1)

    nums := interval(0, SIZE)
    nums = shuffle(nums)

    heap := NewBOHeap()

    for _, elem := range nums {
        heap.Insert(elem)
    }

    for i:=0; i<5; i++ {
        heap.Pop()
    }

    if !isChild_int(heap.root, 6) {t.Errorf("relationship error.")}
}


func Test_insert_shuffle(t *testing.T) {
    const SIZE = 25
    rand.Seed(1)

    nums := interval(0, SIZE)
    nums = shuffle(nums)

    heap := NewBOHeap()

    for _, elem := range nums {
        heap.Insert(elem)
    }

}


func Test_heapsort_shuffle_1(t *testing.T) {
    const SIZE = 25
    rand.Seed(1)

    nums := interval(0, SIZE)
    nums = shuffle(nums)

    heap := NewBOHeap()

    for _, elem := range nums {
        heap.Insert(elem)
    }

    for i:=0; i<SIZE; i++ {
        pval := heap.Pop()
        if pval != i {t.Errorf("expected %d, got %d", i, pval)}
    }
}


func Test_insert_binomial(t *testing.T) {
    heap := NewBOHeap()

    n0 := newBONode(0)

    n1 := newBONode(1)
    n1.rank = 3

    n2 := newBONode(2)
    n2.rank = 4

    n3 := newBONode(3)
    n3.rank = 5

    n4 := newBONode(4)
    n4.rank = 6

    catalyst := newBONode(100)
    catalyst.rank = 3

    heap.insert(n0)
    heap.insert(n1)
    heap.insert(n2)
    heap.insert(n3)
    heap.insert(n4)
    heap.insert(catalyst)

    if n1.rank != 7 {t.Errorf("rank error.")}
    if catalyst.parent != n1 {t.Errorf("relationship error.")}

}


func Test_merge_1(t *testing.T) {
    const (
        SIZE1 = 10
        SIZE2 = SIZE1 * 2
    )

    h1 := NewBOHeap()
    h2 := NewBOHeap()

    insert_mult(h1, interval(0, SIZE1))
    insert_mult(h2, interval(SIZE1, SIZE2))

    h1.Merge(h2)

    if h1.size != SIZE2 {t.Errorf("size error, expected %d, got %d", SIZE2, h1.size)}

    for i:=0; i<SIZE2; i++ {
        pval := h1.Pop()
        if pval != i {t.Errorf("expected %d, got %d", i, pval)}
    }
}


func Test_merge_shuffle(t *testing.T) {
    const (
        SIZE1 = 1000
        SIZE2 = SIZE1 * 2
    )

    h1 := NewBOHeap()
    h2 := NewBOHeap()

    s1 := interval(0, SIZE1)
    s1 = shuffle(s1)
    insert_mult(h1, s1)

    s2 := interval(SIZE1, SIZE2)
    s2 = shuffle(s2)
    insert_mult(h2, s2)

    h1.Merge(h2)

    if h1.size != SIZE2 {t.Errorf("size error, expected %d, got %d", SIZE2, h1.size)}

    for i:=0; i<SIZE2; i++ {
        pval := h1.Pop()
        if pval != i {t.Errorf("expected %d, got %d", i, pval)}
    }
}


// ====== Helpers ======
func interval(start int, end int) []int {
    // [start, end)
    slice := make([]int, 0, end-start)

    for i:=start; i<end; i++ {
        slice = append(slice, i)
    }
    return slice
}

func insert_mult(bq* BOHeap, values []int) {
    for _, elem := range values {
        bq.Insert(elem)
    }
}


func shuffle(slice []int) []int {
    shuffled := make([]int, len(slice), len(slice))
    copy(shuffled, slice)

    for i:=0; i<len(shuffled); i++ {
        index := rand.Intn(len(shuffled)-i)
        index += i

        // swap i'th and index'th element
        tmp := shuffled[i]
        shuffled[i] = shuffled[index]
        shuffled[index] = tmp
    }
    return shuffled
}


func isChild(parent *BONode, canditate *BONode) bool {
    child := parent.children_head
    for child != nil {
        if child == canditate {
            return true
        }
        child = child.rightsibling
    }
    return false
}


func isChild_int(parent *BONode, value int) bool {
    child := parent.children_head
    for child != nil {
        if child.value == value {
            return true
        }
        child = child.rightsibling
    }
    return false
}


func rankOrdered(parent *BONode) bool {
    if parent == nil || !parent.hasChildren() || parent.children_head.rightsibling == nil {
        // no children or 1 child.
        return true
    }

    prev := parent.children_head
    next := parent.children_head.rightsibling

    for next != nil {
        if prev.rank > next.rank {
            return false
        }
        prev = next
        next = next.rightsibling
    }
    return true
}