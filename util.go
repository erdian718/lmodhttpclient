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
	"io"
	"net/http"

	"ofunc/lua"
)

func toReader(l *lua.State, i int) io.Reader {
	if v, ok := l.GetRaw(i).(io.Reader); ok {
		return v
	} else {
		panic("http/client: not a reader")
	}
}

func toResp(l *lua.State, i int) response {
	if v, ok := l.GetRaw(i).(response); ok {
		return v
	} else {
		panic("http/client: not a response")
	}
}

func result(l *lua.State, resp *http.Response, err error) int {
	if err != nil {
		if resp != nil {
			resp.Body.Close()
		}
		l.Push(nil)
		l.Push(err.Error())
		return 2
	}
	l.Push(response{resp})
	l.PushIndex(lua.FirstUpVal - 1)
	l.SetMetaTable(-2)
	return 1
}
