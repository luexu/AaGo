package aa

import (
	"log"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/hi-iwi/AaGo/dtype"
)

type Config interface {
	Get(key string, defaultValue ...interface{}) *dtype.Dtype
}

func parseToDuration(d string) time.Duration {
	if len(d) < 2 {
		return 0
	}
	var t int
	if d[len(d)-2:] == "ms" {
		t, _ = strconv.Atoi(d[0 : len(d)-2])
		return time.Duration(t) * time.Millisecond
	}

	if d[len(d)-1:] == "s" {
		t, _ = strconv.Atoi(d[0 : len(d)-1])
	} else {
		t, _ = strconv.Atoi(d)
	}
	return time.Duration(t) * time.Second
}


func splitDots(keys ...string) []string {
	n := make([]string, 0)
	for _, key := range keys {
		n = append(n, strings.Split(key, ".")...)
	}
	return n
}

func defaultDtype(key string, defaultValue ...interface{}) *dtype.Dtype {
	dv := parseDefaultValue(defaultValue...)
	if len(defaultValue) == 0 {
		log.Println("not found config " + key)
	}
	return dtype.New(dv)
}

func parseDefaultValue(vs ...interface{}) interface{} {
	if len(vs) > 0 {
		return vs[0]
	}
	return ""
}

func (app *Aa) ParseConfig(filename string) error {
	switch path.Ext(filename) {
	case ".ini":
		app.ParseIni(filename)
		// case ".yml", ".yaml":
		// 	a.ParseYml(filename)
	}

	// a.mu.Lock()
	// for k, v := range c {
	// 	a.Config[k] = NewDtype(v)
	// }
	// a.mu.Unlock()

	app.ParseToConfiguration()
	return nil
}
