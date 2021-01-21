package com

import (
	"regexp"

	"github.com/hi-iwi/AaGo/ae"
	"github.com/hi-iwi/dtype"
)

/*
@param pattern  e.g. `[[:word:]]+` `\w+`
Filter(pattern string, required bool)
Filter(required bool)
Filter(pattern string)
Filter(default dtype.Dtype)

*/
func (p *ReqProp) Filter(patterns ...interface{}) *ae.Error {
	required := true
	pattern := ""

	for i := 0; i < len(patterns); i++ {
		pat := patterns[i]
		if s, ok := pat.(string); ok {
			pattern = s
		} else if b, ok := pat.(bool); ok {
			required = b
		} else if d, ok := pat.(*dtype.Dtype); ok && p.String() == "" {
			p.Value = d.Value
		}
	}
	if p.String() == "" {
		if required {
			return ae.NewError(400, "Parameter `"+p.param+"` is required!")
		}
	} else if pattern != "" {
		re, _ := regexp.Compile(pattern)
		m := re.FindStringSubmatch(p.String())
		if m == nil || len(m) < 1 {
			return ae.NewError(400, "Parameter `"+p.param+"`=`"+p.String()+"` dose not match `"+pattern+"`")
		}
	}
	return nil
}
