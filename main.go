//
// Copyright 2020 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"golang.org/x/exp/slices"
)

func main() {
	w := colorable.NewColorableStdout()
	for i, arg := range os.Args {
		argColor := color.New(color.BgBlue, color.FgWhite)
		fmt.Fprintf(w, "%s%s\n",
			color.RedString("Arg %d: ", i),
			argColor.Sprint(arg))
	}

	fmt.Println()
	fmt.Println("Environment:")
	envs := os.Environ()
	slices.Sort(envs)
	for _, env := range envs {
		argColor := color.New(color.BgBlue, color.FgWhite)
		fmt.Fprintf(w, "  - %s\n", argColor.Sprint(env))
	}
}
