/*
Copyright 2019 by ofunc

This software is provided 'as-is', without any express or implied warranty. In
no event will the authors be held liable for any damages arising from the use of
this software.

Permission is granted to anyone to use this software for any purpose, including
commercial applications, and to alter it and redistribute it freely, subject to
the following restrictions:

1. The origin of this software must not be misrepresented; you must not claim
that you wrote the original software. If you use this software in a product, an
acknowledgment in the product documentation would be appreciated but is not
required.

2. Altered source versions must be plainly marked as such, and must not be
misrepresented as being the original software.

3. This notice may not be removed or altered from any source distribution.
*/

// http.Client bindings for Lua.
package lmodhttpclient

import (
	"io"
	"net/http"

	"ofunc/lua"
)

// Open opens the module.
func Open(l *lua.State) int {
	m := metaResp(l)
	l.NewTable(0, 8)

	l.Push("version")
	l.Push("0.0.1")
	l.SetTableRaw(-3)

	l.Push("head")
	l.PushClosure(lHead, m)
	l.SetTableRaw(-3)

	l.Push("get")
	l.PushClosure(lGet, m)
	l.SetTableRaw(-3)

	l.Push("post")
	l.PushClosure(lPost, m)
	l.SetTableRaw(-3)

	return 1
}

func lHead(l *lua.State) int {
	resp, err := http.Head(l.ToString(1))
	return result(l, resp, err)
}

func lGet(l *lua.State) int {
	resp, err := http.Get(l.ToString(1))
	return result(l, resp, err)
}

func lPost(l *lua.State) int {
	var resp *http.Response
	var err error
	if r, ok := l.GetRaw(3).(io.Reader); ok {
		resp, err = http.Post(l.ToString(1), l.ToString(2), r)
	} else {
		resp, err = http.PostForm(l.ToString(1), toForm(l, 3))
	}
	return result(l, resp, err)
}
