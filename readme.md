Introduction
============
Brodal-Okasaki heap is a tree structure that satisfies the heap property. It is asympotically optimal, supporting

* Inserting a new value
* Finding the minimum value
* Merging with another queue
* Decreasing an element within the queue

on O(1) worst-case execution time, and supporting

* Deleting the minimum value

on O(logn) time.

The heap is first described in the paper that may be accessed from [here](http://www.brics.dk/RS/96/37/BRICS-RS-96-37.pdf).
This implementation follows the original description as close as possible, but there are a few differences/ambigiuous
points that is explained in the code.


Prerequisites
=============
One who wish to venture further should get familiar with

* Heaps in general and
* Binomial Heap  (I have an implementation using Go in [this repository](https://github.com/ckaygusu/BinomialHeap).)

You are free to try to make sense what is present in this repository, but I think getting familiar with these concepts
will help you have easier time with it.


Brodal-Okasaki heap
===================
This data structure is a heavily modified version of the binomial heap;

* A different subtree linking mechanism is introduced to make Insert() perform in O(1).
    * Takes O(logn) in ordinary binomial heap.
    * This variant is called "Skew-Binomial Heap".
* Adding a root node that holds the minimum element.
* The data structure is allowed to contain itself to make Merge() perform in O(1).
    * Takes O(logn) in skew-binomial and ordinary binomial heaps.

These three modifications yields us an asymptotically optimal data structure, which means we cannot improve any of the
worst-case bounds without increasing the bound for some other operation.


Asymptotic Optimality
=======================
Skew-Binomial heaps do really reduce the work done over the binomial heaps, but Brodal-Okasaki heaps achieve the same
asymptotical improvement by doing most of the actual work under Pop().

* Peek() ("findMin" operation) takes O(1) time because we determine the next minimum element within the runtime of
Pop() operation. Determining the minimum element takes O(logn) time.
* Merge() takes O(1) time, because this operation simply puts the other queue under the carpet and inserts only
the root node of other queue. Actual merging operation is done within the runtime of Pop() when we pop a node that
happens to have nodes "under the carpet".
* Actual Pop() operation (as in binomial heap) already takes O(logn) time.

Since it does not matter whether we perform 3*logn or 50*logn operations, under "O" notation they are all O(logn).
By this little hack, we achieve asymptotical optimality, where everything except Pop() is done under O(1).

As you can see, we don't really reduce the amount of work done here, we just push them under the same umbrella so
the chart looks nice and dandy. Because of this, I agree with the authors stating this data structure is of little
practical interest but serves as an theoretical example that this kind of asymptotical performance is possible.


Functional vs Imperative
========================
In the paper, Brodal-Okasaki heap has a functional implementation. Here, I made the implementation in a imperative
language (Go) so some of the concepts may not translate directly.


Puzzles
=======
After grasping this data structure, you may play with it by

* Optimizing for performance. This implementation is horribly inefficient, and there are a lot of points that can
be optimized.
* I left out some questions as comments in the code. Try to tackle them, they are not too complicated.