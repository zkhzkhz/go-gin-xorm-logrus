package main

import "fmt"

type welcome string

type Handler interface {
	Do(k, v interface{})
}

func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for key, value := range m {
			h.Do(key, value)
		}
	}
}

func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandleFunc(f))
}

type HandleFunc func(k, v interface{})

func (f HandleFunc) Do(k, v interface{}) {
	f(k, v)
}

func (w welcome) selfInfo(k, v interface{}) {
	fmt.Printf("%s,我叫%s,今年%d岁\n", w, k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 22
	persons["李四"] = 23
	persons["王五"] = 24

	var w welcome = "大家好"
	EachFunc(persons, w.selfInfo)
}
