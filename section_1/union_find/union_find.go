//package union_find
package main

import (
	"fmt"
)

type union_t struct {
	n_connections int
	connection []int
}

/**
 *	Initializes n locations
 **/
func (u *union_t) UN(n int) {
	u.n_connections = n
	u.connection = make([]int, n)
	for i := 0; i < n; i++ {
		u.connection[i] = i
	}
}

/**
 *	Adds connection between p and q
 **/
func (u *union_t)union(p, q int) {
	pid := u.connection[p]	// lol
	qid := u.connection[q]

	// if theyre connected already sweet
	if pid == qid {
		return
	}

	u.connection[p] = qid
	
	// decrement number of connections by 1
	u.n_connections--
}

/**
 *	Identifier for p
 **/
func (u *union_t)find(p int) int{
	for p != u.connection[p] {
		p = u.connection[p]
	}
	return p
}

/**
 *	Bool returning if two edges are connected	
 **/
func (u *union_t)connected(p, q int) bool {
	return u.find(p) == u.find(q)
}

/**
 *	Returns number of edges in union
 **/
func (u union_t)count() int {
	return u.n_connections
}

func main() {
	fmt.Println("Hello world")
}