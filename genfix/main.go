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

	"github.com/protofix/protofix/genfix"
	"github.com/protofix/protofix/specfix"
)

var specs = []spec{
	spec{pkg: genfix.Package{Name: "fix40", Format: "FIX.4.0", Info: "FIX.4.0", Spec: specfix.FIX40}, opts: genfix.FIX40()},
	spec{pkg: genfix.Package{Name: "fix41", Format: "FIX.4.1", Info: "FIX.4.1", Spec: specfix.FIX41}, opts: genfix.FIX41()},
	spec{pkg: genfix.Package{Name: "fix42", Format: "FIX.4.2", Info: "FIX.4.2", Spec: specfix.FIX42}, opts: genfix.FIX42()},
	spec{pkg: genfix.Package{Name: "fix43", Format: "FIX.4.3", Info: "FIX.4.3", Spec: specfix.FIX43}, opts: genfix.FIX43()},
	spec{pkg: genfix.Package{Name: "fix44", Format: "FIX.4.4", Info: "FIX.4.4", Spec: specfix.FIX44}, opts: genfix.FIX44()},
	spec{pkg: genfix.Package{Name: "fix50", Format: "FIX.5.0", Info: "FIX.5.0", Spec: specfix.FIX50}, opts: genfix.FIX50()},
	spec{pkg: genfix.Package{Name: "fix50sp1", Format: "FIX.5.0", Info: "FIX.5.0 servicepack 1", Spec: specfix.FIX50SP1}, opts: genfix.FIX50SP1()},
	spec{pkg: genfix.Package{Name: "fix50sp2", Format: "FIX.5.0", Info: "FIX.5.0 servicepack 2", Spec: specfix.FIX50SP2}, opts: genfix.FIX50SP2()},
	spec{pkg: genfix.Package{Name: "fixt11", Format: "FIXT.1.1", Info: "FIXT.1.1 (transport) ", Spec: specfix.FIXT11}, opts: genfix.FIXT11()},
	spec{pkg: genfix.Package{Name: "moex44", Format: "FIX.4.4", Info: "MOEX.4.4 (FIX.4.4)", Spec: specfix.MOEX44}, opts: genfix.MOEX44()},
}

type spec struct {
	pkg  genfix.Package
	opts []genfix.Option
}

//go:embed test.xml
var testXML []byte

var testSpec = genfix.Package{Name: "testspec", Format: "FIX.0.1", Info: "test FIX.0.1", Spec: testXML}

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

	// Generate specs.
	for _, s := range specs {
		err = genfix.Generate(pth, s.pkg, s.opts...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nGenerate FIX format by specification: %s\n", err)
			os.Exit(1)
		}
	}

	// Generate test spec.
	err = genfix.Generate(path.Join(pth, "genfix/internal"), testSpec, genfix.FIX44()...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nGenerate test FIX format by test specification: %s\n", err)
		os.Exit(1)
	}
}
