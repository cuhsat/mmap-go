// Copyright 2011 Evan Shaw. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || openbsd || solaris || netbsd || aix

package mmap

import (
	"golang.org/x/sys/unix"
)

func mmap(len int, _, inflags, fd uintptr, off int64) ([]byte, error) {
	flags := unix.MAP_SHARED
	prot := unix.PROT_READ
	if inflags&ANON != 0 {
		flags |= unix.MAP_ANON
	}

	b, err := unix.Mmap(int(fd), off, len, prot, flags)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (m *MMap) lock() error {
	return unix.Mlock(*m)
}

func (m *MMap) unlock() error {
	return unix.Munlock(*m)
}

func (m *MMap) unmap() error {
	return unix.Munmap(*m)
}
