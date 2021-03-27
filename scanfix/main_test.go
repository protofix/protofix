// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanfix_test

import "runtime"

func line() int { _, _, l, _ := runtime.Caller(1); return l }
