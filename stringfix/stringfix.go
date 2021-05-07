// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stringfix

import (
	"regexp"
	"strings"
)

var toSnakeCaseMatchFirst = regexp.MustCompile("(.)([A-Z][a-z]+)")
var toSnakeCaseMatchAll = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := toSnakeCaseMatchFirst.ReplaceAllString(str, "${1}_${2}")
	snake = toSnakeCaseMatchAll.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
