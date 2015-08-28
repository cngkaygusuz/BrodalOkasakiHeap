package BrodalOkasakiHeap

/*
Generic operations to be performed over heap structure.
 */
type PriorityQueue interface {
    Insert(int)				// Insert an element into the pqueue.
    Pop()    	int			// Return and remove the topmost key, determined by the implementation.
    Peek()    	int			// Return the topmost key
    Size()		int			// Get the size of pqueue.
    Merge(PriorityQueue)	// Merge two of the same-type priority queue (this means you shouldn't attempt merging a binary heap with a binomial heap).

    // Brodal-Okasaki heaps does not support DecreaseKey() operation, but in the discussion section of the paper,
    // the author mentions a few ideas on how it might be possible.
    // DecreaseKey()
}

