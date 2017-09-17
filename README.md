# ckjson
[![Build Status](https://travis-ci.org/cowkeys/ckjson.png?branch=master)](https://travis-ci.org/cowkeys/ckjson)
[![Coverage Status](https://coveralls.io/repos/github/cowkeys/ckjson/badge.png?branch=master)](https://coveralls.io/github/cowkeys/ckjson?branch=master)

#### ckjson can be used to convert json string to go-struct quickly

**Features:**

-   No other dependency.
-   Adapter to decode Deep nesting json.
-   Show on cmd


Installation
------------

```
$ go get github.com/cowkeys/ckjson
```

Quick start
-----------
### base use
```go
package main

import (
   "github.com/cowkeys/ckjson"
)

func main() {
	json := `{"employees": [
    { "firstName":"Bill" , "salary":10000.00,"flag":true },
    { "first-Name":"George" , "lastName":"Bush","flag":false },
    { "firstName":"Thomas" , "depart_tag":"IT" }]}`
    
	ck := ckjson.NewCkj("TestCK", json)
	ck.JsonToStruct()
}
```
##### output:

```go
type Employees struct {
	FirstName string  `json:"firstName"`
	Salary    float64 `json:"salary"`
	Flag      bool    `json:"flag"`
	LastName  string  `json:"lastName"`
	DepartTag string  `json:"depart_tag"`
}

type TestCK struct {
	Employees Employees `json:"employees"`
}
```

### deep nesting
```go
    json := `{"company": {
        "name": "google",
        "address": {
            "country": "us",
            "code": 1234,
            "loop_one": {
                "loop_two": [
                    1,
                    2
                ]
            }}}}`

	ck3 := ckjson.NewCkj("TestCK", json)
	ck3.JsonToStruct()
```
##### output:

```go
type LoopTwo struct {
}

type LoopOne struct {
	LoopTwo LoopTwo `json:"loop_two"`
}

type Address struct {
	Country string  `json:"country"`
	Code    float64 `json:"code"`
	LoopOne LoopOne `json:"loop_one"`
}

type Company struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

type TestCK struct {
	Company Company `json:"company"`
}

```
### name cases (upper/lower/camel)
```go
//auto filter '_' , '-' , ':' , '.' ...
json :=    `{ "first_name":"Bill" }`

ck := ckjson.NewCkj("TestCK", t4)
ck.JsonToStruct()//default : aaa_BCD -> AaaBcd (Camel case)

ck.LowerCase = true//aaa_BCD -> Aaabcd (except first char)
ck.JsonToStruct()

ck.UpperCase = true
ck.JsonToStruct()//aaa_BCD -> AAABCD (except first char) 
```
##### output:
```go
type TestCK struct {
	FirstName string
}

type TestCK struct {
	Firstname string
}

type TestCK struct {
	FIRSTNAME string
}
```
### tag display
```go
json :=    `{ "first_name":"Bill" }`

ck := ckjson.NewCkj("TestCK", json)
ck.JsonToStruct()

ck.JsonTag = false // set not display tag
ck.JsonToStruct()
```
##### output:
```go
type TestCK struct {
	FirstName string `json:"first_name"`
}

type TestCK struct {
	FirstName string
}

```

Document
--------

-   `func NewCkj(name,msg string) *ckjson`

    new a ckjson client

-   `func (*NewCkj)JsonToStruct()`

    print go struct


