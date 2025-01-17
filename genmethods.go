// Copyright (c) 2014 Dataence, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build ignore

package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	methods = []string{
		"GET",
		"PUT",
		"POST",
		"DELETE",
		"CONNECT",
		"OPTIONS",
		"TRACE",
		"PATCH",
		"PROPFIND",
		"PROPPATCH",
		"MKCOL",
		"COPY",
		"MOVE",
		"LOCK",
		"UNLOCK",
		"VERSION_CONTROL",
		"CHECKOUT",
		"UNCHECKOUT",
		"CHECKIN",
		"UPDATE",
		"LABEL",
		"REPORT",
		"MKWORKSPACE",
		"MKACTIVITY",
		"BASELINE_CONTROL",
		"MERGE",
		"INVALID",
	}
)

type node struct {
	b byte    // node value
	f bool    // end of string
	s int     // state
	c []*node // children
	m int     // suffix length
	w string  // suffix
}

func outputFile(fname string) *os.File {
	var (
		ofile *os.File = os.Stdin
		err   error
	)

	if fname != "" {
		// Open output file
		ofile, err = os.OpenFile(fname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		if err != nil {
			logger.HandleFatal(err)
		}
	}

	return ofile
}

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		logger.HandleFatal("Invalid command. Must be 'genmethods <outfile>.'")
	}

	ofile := outputFile(flag.Arg(0))
	defer ofile.Close()

	s := 0 // state
	root := &node{s: s}
	nodes := append(make([]*node, 0, 10), root)

	for _, ss := range methods {
		ss += " "
		cur := root

		// node before last dot
		l := len(ss)

		for i, b := range []byte(ss) {
			found := false
			var n *node

			for _, n = range cur.c {
				if n.b == b {
					found = true
					break
				}
			}

			if !found {
				s++
				n = &node{b: b, s: s, m: i + 1}
				cur.c = append(cur.c, n)
				nodes = append(nodes, n)
			}

			cur = n
			if i == l-1 {
				n.w = ss
				n.f = true
			}
		}
	}

	fmt.Fprintf(ofile, `// Copyright (c) 2014 Dataence, LLC. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file is automatically generated by genfsm.go. Please DO NOT modify.
// This file was generated at %s.

package sequence

func matchRequestMethods(ss string) int {
	var (
		m int			// match length
		s int			// state
	)

loop:
	for _, b := range ss {
		switch s {
`, time.Now())

	var lcdiff byte = 'a' - 'A'

	for _, n := range nodes {
		if len(n.c) > 0 {
			fmt.Fprintf(ofile, "\t\tcase %d:\n", n.s)
			fmt.Fprintf(ofile, "\t\t\tswitch b {\n")

			for _, c := range n.c {
				if c.b != ' ' {
					fmt.Fprintf(ofile, "\t\t\tcase '%c', '%c':\n", c.b, c.b|lcdiff)
				} else {
					fmt.Fprintf(ofile, "\t\t\tcase '%c':\n", c.b)
				}
				fmt.Fprintf(ofile, "\t\t\t\ts = %d\n", c.s)
				if c.f {
					// in case we hit something like "zblogspot.com.ar", result
					// should be "com.ar" but there's also "blogspot.com.ar", so
					// we need to make sure we are either behind a "." or it's
					// the beginning.
					fmt.Fprintf(ofile, `m = %d// %s - final`, c.m-1, c.w)
				}
			}

			fmt.Fprintf(ofile, "\t\t\tdefault:\n\t\t\t\tbreak loop\n\t\t\t}\n")
		}
	}

	fmt.Fprintln(ofile, `		default:
			break loop
		}
	}

	//fmt.Println("m =", m, "l =", l)

	return m
}`)
}
