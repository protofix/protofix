// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package codecfix

var (
	// ASCII encrypts/decrypts ASCII value of the FIX field.
	ASCII = NOPCrypt{}
)

// Crypt encrypts/decrypts a value of a FIX field.
type Cryptor interface {
	// Encrypt encrypts a field of the FIX.
	Encrypt([]byte) ([]byte, error)
	// Decrypt decrypts an encrypted field of the FIX.
	Decrypt([]byte) ([]byte, error)
}

// NOPCrypt is the "no operation" encryptor/decryptor for a FIX values.
type NOPCrypt struct{}

func (NOPCrypt) Encrypt(p []byte) ([]byte, error) { return p, nil }
func (NOPCrypt) Decrypt(p []byte) ([]byte, error) { return p, nil }
