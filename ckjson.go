package ckjson

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

const (
	E_map    = "map[string]interface {}"
	E_list   = "[]interface {}"
	E_string = "string"
)

var (
	F_s      = "%s"
	F_begin  = "type %s struct {"
	F_end    = "}\n"
	F_kv     = "%s %s %s"
	F_kv_ori = "%s %s %s"
	F_kv_tag = "%s %s `json:\"%s\"`"
)

type ckjson struct {
	Name      string
	Msg       string
	LowerCase bool
	UpperCase bool // default CamelCase
	JsonTag   bool
	TagMap    map[string]string
	Parent    map[string]interface{}
	Subs      []map[int]string
	Out       []string
}

func NewCkj(name, msg string) *ckjson {
	return &ckjson{
		Name:    name,
		Msg:     msg,
		JsonTag: true,
		TagMap:  map[string]string{},
	}
}

func (c *ckjson) FlushCkj() {
	*c = *&ckjson{
		Name:      c.Name,
		Msg:       c.Msg,
		TagMap:    map[string]string{},
		Parent:    map[string]interface{}{},
		Subs:      []map[int]string{},
		Out:       []string{},
		JsonTag:   true,
		LowerCase: false,
		UpperCase: false,
	}
}

func (c *ckjson) JsonToStruct() *ckjson {
	F_kv = F_kv_ori
	if c.JsonTag {
		F_kv = F_kv_tag
	}
	if err := json.Unmarshal([]byte(c.Msg), &c.Parent); err != nil {
		fmt.Println("parse json failed:", err)
		return nil
	}

	c.appendStr(fmt.Sprintf(F_begin, c.Name))
	for k, v := range c.Parent {
		orik, k := c.keyFilter(k)
		tmptype := fmt.Sprintf(F_s, reflect.TypeOf(v))
		if tmptype == E_map {
			c.subStruct(orik, v)
			tmptype = k
		}
		if tmptype == E_list {
			c.subList(orik, v)
			tmptype = k
		}

		c.appendStr(fmt.Sprintf(F_kv, k, tmptype, c.TagMap[k]))
	}
	c.appendStr(F_end)
	c.printStruct()
	c.FlushCkj()
	return c
}

func (c *ckjson) subStruct(key string, out interface{}) {
	dout := out.(map[string]interface{})
	m := map[int]string{}
	index := 0
	sk := c.keyCase(key)
	m[index] = fmt.Sprintf(F_begin, sk)
	for k, v := range dout {
		orik, k := c.keyFilter(k)
		tmptype := fmt.Sprintf(F_s, reflect.TypeOf(v))
		if tmptype == E_map {
			c.subStruct(orik, v)
			tmptype = k
		}
		if tmptype == E_list {
			c.subList(orik, v)
			tmptype = k
		}

		index++
		m[index] = fmt.Sprintf(F_kv, k, tmptype, c.TagMap[k])
	}
	m[index+1] = F_end
	c.Subs = append(c.Subs, m)
}

func (c *ckjson) subList(key string, out interface{}) {
	dout := out.([]interface{})
	list := map[string]interface{}{}
	for _, vl := range dout {
		vtype := fmt.Sprintf(F_s, reflect.TypeOf(vl))
		if vtype == E_list {
			c.subList(key, vl)
			continue
		}

		if vtype == E_map {
			m := vl.(map[string]interface{})
			for k, v := range m {
				_, k := c.keyFilter(k)
				if _, ok := list[k]; ok {
					continue
				}
				list[k] = v
			}
		}

	}
	c.subStruct(key, list)
}

func (c *ckjson) appendStr(kv string) {
	c.Out = append(c.Out, kv)
}

//output cmd
func (c *ckjson) printStruct() {
	if c == nil {
		return
	}

	for _, value := range c.Subs {
		for i := 0; i < len(value); i++ {

			fmt.Println(value[i])
		}
	}

	for i := 0; i < len(c.Out); i++ {
		fmt.Println(c.Out[i])
	}
}

func (c *ckjson) keyFilter(str string) (string, string) {
	temp := c.keyCase(str)
	if _, ok := c.TagMap[temp]; !ok {
		c.TagMap[temp] = str
	}

	if c.JsonTag {
		return str, temp
	}

	c.TagMap[temp] = ""
	return str, temp

}

//change name to different case
func (c *ckjson) keyCase(str string) string {
	temp := strings.FieldsFunc(str, c.CkSplit)
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 && y == 0 {
				vv[i] = unicode.ToUpper(vv[i])
				upperStr += string(vv[i])
				continue
			}
			if c.UpperCase {
				vv[i] = unicode.ToUpper(vv[i])
				upperStr += string(vv[i])
				continue
			}
			if c.LowerCase {
				vv[i] = unicode.ToLower(vv[i])
				upperStr += string(vv[i])
				continue
			}
			if i == 0 {
				vv[i] = unicode.ToUpper(vv[i])
				upperStr += string(vv[i])
				continue
			}

			upperStr += string(vv[i])
			continue
		}
	}
	return upperStr
}

//name filter from 4 char
func (c *ckjson) CkSplit(r rune) bool {
	return r == ':' || r == '.' || r == '-' || r == '_'
}
