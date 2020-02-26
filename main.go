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
)

func main() {
	w := colorable.NewColorableStdout()
	for i, arg := range os.Args {
		argColor := color.New(color.BgBlue, color.FgWhite)
		fmt.Fprintf(w, "%s%s\n",
			color.RedString("Arg %d: ", i),
			argColor.Sprint(arg))
	}
}
