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
	"net/http/cookiejar"
	"net/url"

	"ofunc/lua"

	"golang.org/x/net/publicsuffix"
)

func init() {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	http.DefaultClient.Jar = jar
}

func toResp(l *lua.State, i int) response {
	if v, ok := l.GetRaw(i).(response); ok {
		return v
	} else {
		panic("http/client: not a response")
	}
}

func toForm(l *lua.State, i int) url.Values {
	form := make(url.Values, l.Count(i))
	l.ForEach(i, func() bool {
		k := l.ToString(-2)
		if l.GetMetaField(-1, "__pairs") != lua.TypeNil {
			l.Pop(1)
			l.ForEach(-1, func() bool {
				form.Add(k, l.ToString(-1))
				return true
			})
		} else if l.GetMetaField(-1, "__len") != lua.TypeNil {
			l.PushIndex(-2)
			l.Call(1, 1)
			n := int(l.ToInteger(-1))
			l.Pop(1)
			for i := 1; i <= n; i++ {
				l.Push(i)
				l.GetTable(-2)
				form.Add(k, l.ToString(-1))
				l.Pop(1)
			}
		} else if l.TypeOf(-1) == lua.TypeTable {
			l.ForEachRaw(-1, func() bool {
				form.Add(k, l.ToString(-1))
				return true
			})
		} else {
			form.Add(k, l.ToString(-1))
		}
		return true
	})
	return form
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

func header(l *lua.State, h http.Header) {
	l.NewTable(0, len(h))
	for k, v := range h {
		l.Push(k)
		l.NewTable(len(v), 0)
		for i, x := range v {
			l.Push(i + 1)
			l.Push(x)
			l.SetTableRaw(-3)
		}
		l.SetTableRaw(-3)
	}
}
