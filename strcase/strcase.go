// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strcase

import (
	"regexp"
	"strings"
)

var camelToSnakeMatchFirst = regexp.MustCompile("(.)([A-Z][a-z]+)")
var camelToSnakeMatchAll = regexp.MustCompile("([a-z0-9])([A-Z])")

func CamelToSnake(str string) string {
	s := camelToSnakeMatchFirst.ReplaceAllString(str, "${1}_${2}")
	s = camelToSnakeMatchAll.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(s)
}
