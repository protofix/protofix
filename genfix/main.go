// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/danil/protofix/genfix"
	"github.com/danil/protofix/specfix"
)

var specs = []genfix.Package{
	genfix.Package{Name: "moex44", Info: "MOEX.4.4 (FIX.4.4)", Spec: specfix.MOEX44},
}

//go:embed test.xml
var testXML []byte

var testSpec = genfix.Package{Name: "testspec", Info: "test FIX", Spec: testXML}

func main() {
	log.SetFlags(0)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "os: get current/working directory: %s\n", err)
		os.Exit(1)
	}

	dirs := strings.Split(dir, string(os.PathSeparator))
	dirs = append(dirs[:0], append([]string{"/"}, dirs...)...)

	pth := path.Join(dirs[:len(dirs)-1]...)

	err = genfix.Generate(pth, specs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nGenerate FIX format by specification: %s\n", err)
		os.Exit(1)
	}

	err = genfix.Generate(path.Join(pth, "genfix/internal"), []genfix.Package{testSpec})
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nGenerate test FIX format by test specification: %s\n", err)
		os.Exit(1)
	}
}
