package com

import (
	"github.com/hi-iwi/AaGo/dtype"
	"strconv"
	"strings"
)

type ReqProp struct {
	dtype.Dtype
	param string
}

func NewReqProp(param string, data interface{}) *ReqProp {
	var p ReqProp
	p.param = param
	p.Value = data
	return &p
}

func (p *ReqProp) Default(v interface{}) {
	if p.IsEmpty() {
		p.Value = v
	}
}

func rexp(elems string) string {
	return "^(" + elems + ")$"
}

func UintsRegExp(set ...interface{}) string {
	elems := make([]string, len(set))
	for i, v := range set {
		w, err := dtype.Uint64(v)
		if err != nil {
			continue
		}
		elems[i] = strconv.FormatUint(w, 10)
	}
	return StringsRegExp(elems)
}

func IntsRegExp(ies []interface{}) string {
	elems := make([]string, len(ies))
	for i, v := range ies {
		w, err := dtype.Int64(v)
		if err != nil {
			continue
		}
		elems[i] = strconv.FormatInt(w, 10)
	}
	return StringsRegExp(elems)
}

func StringsRegExp(elems []string) string {
	switch len(elems) {
	case 0:
		return rexp("")
	case 1:
		return rexp(dtype.String(elems[0]))
	}
	n := len(elems) - 1
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}
	var b strings.Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for _, s := range elems[1:] {
		b.WriteString("|")
		b.WriteString(s)
	}
	return rexp(b.String())
}
