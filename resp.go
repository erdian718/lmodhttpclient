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

package lmodhttpclient

import (
	"net/http"

	"ofunc/lua"
)

type response struct {
	*http.Response
}

func (resp response) Read(p []byte) (int, error) {
	return resp.Body.Read(p)
}

func metaResp(l *lua.State) int {
	l.NewTable(0, 2)
	idx := l.AbsIndex(-1)

	l.Push("__index")
	l.Push(lRespIndex)
	l.SetTableRaw(idx)

	return idx
}

func lRespIndex(l *lua.State) int {
	resp := toResp(l, 1)
	switch key := l.ToString(2); key {
	case "status":
		l.Push(resp.StatusCode)
	case "header":
		l.Push(resp.Header) // TODO
	case "close":
		l.Push(lRespClose)
	}
	return 1
}

func lRespClose(l *lua.State) int {
	if err := toResp(l, 1).Body.Close(); err == nil {
		return 0
	} else {
		l.Push(err.Error())
		return 1
	}
}

func toResp(l *lua.State, i int) response {
	if v, ok := l.GetRaw(i).(response); ok {
		return v
	} else {
		panic("http/client: not a response")
	}
}
