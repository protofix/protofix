// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/danil/protofix/genfix"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "os: get current/working directory: %s\n", err)
		os.Exit(1)
	}

	dirs := strings.Split(dir, string(os.PathSeparator))
	dirs = append(dirs[:0], append([]string{"/"}, dirs...)...)

	pth := path.Join(dirs[:len(dirs)-1]...)

	err = genfix.Generate(pth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nGenerate FIX format by specification: %s\n", err)
		os.Exit(1)
	}
}
