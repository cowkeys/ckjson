package ckjson_test

import (
	"testing"

	"github.com/cowkeys/ckjson"
)

func TestCkJson(t *testing.T) {
	baseuse()
	nesting()
	nesting2()
	tag()
	cases()
	big()
}

func baseuse() {
	json := `{"employees": [
    { "firstName":"Bill" , "salary":10000,"flag":true },
    { "first-Name":"George" , "lastName":"Bush","flag":false },
    { "firstName":"Thomas" , "depart_tag":"IT" }]}`
	ck := ckjson.NewCkj("TestCK", json)
	ck.JsonTag = true
	ck.JsonToStruct()
}

func nesting() {
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
	ck := ckjson.NewCkj("TestCK", json)
	ck.JsonTag = true
	ck.JsonToStruct()
}

func nesting2() {
	json := `{
    "button": [
        {
            "name": "MAg",
            "sub_button": [
                {
                    "type": "view",
                    "name": "magzine",
                    "url": "test1"
                },
                {
                    "type": "view",
                    "name": "magzine2",
                    "url": "test2"
                },
                {
                    "type": "view",
                    "name": "times",
                    "url": "test3"
                }
            ]
        },
        {
            "name": "favor",
            "sub_button": [
                {
                    "type": "click",
                    "name": "book",
                    "key": "v_1_2"
                },
                {
                    "type": "click",
                    "name": "book2",
                    "key": "v_1_1"
                },
                {
                    "type": "click",
                    "name": "web",
                    "key": "v_2_1",
                    "price":1000
                }
            ]
        },
        {
            "name": "account",
            "sub_button": [
                {
                    "type": "click",
                    "name": "count",
                    "key": "v_3_1"
                },
                {
                    "type": "click",
                    "name": "cs",
                    "key": "v_3_2"
                }
            ]
        }
    ]
}`

	ck := ckjson.NewCkj("TestCK", json)
	ck.JsonToStruct()

	ck.JsonTag = false // set not display tag
	ck.JsonToStruct()

}

func tag() {
	json := `{ "first_name":"Bill" }`

	ck := ckjson.NewCkj("TestCK", json)
	ck.JsonToStruct()

	ck.JsonTag = false // set not display tag
	ck.JsonToStruct()

}

func cases() {
	json := `{ "first_name":"Bill" }`
	ck := ckjson.NewCkj("TestCK", json)
	ck.JsonToStruct() //default : aaa_BCD -> AaaBcd (Camel case)

	ck.LowerCase = true //aaa_BCD -> Aaabcd (except first char)
	ck.JsonToStruct()

	ck.UpperCase = true
	ck.JsonToStruct() //aaa_BCD -> AAABCD (except first char)
}

func big() {
	json := `{"cowkys":[{"code":"AUDCAD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"AUDCHF","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"AUDHKD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"AUDHKD.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"AUDJPY","digits":3,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"AUDNZD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"AUDTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"AUDUSD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"AUS200","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1},{"code":"BRENT","digits":2,"stopslevel":10,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":100},{"code":"BRENTX7","digits":2,"stopslevel":10,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":100},{"code":"BUNDZ7","digits":2,"stopslevel":10,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1000},{"code":"CAC40","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1},{"code":"CADCHF","digits":5,"stopslevel":25,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"CADJPY","digits":3,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"CADTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"CHFJPY","digits":3,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"CHFTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"DAX30","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1},{"code":"DKKTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"ESP35","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1},{"code":"ESX50","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1},{"code":"EURAUD","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURCAD","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURCHF","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURCNH","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURCZK","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURDKK","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURDKK.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"EURGBP","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURHKD","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"EURHKD.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"EURHKD.vip","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"EURHUF","digits":3,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURJPY","digits":3,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURMXN","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"EURMXN.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"EURMXN.vip","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"EURNOK","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURNZD","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURPLN","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURSEK","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURSGD.vip","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"EURTRY","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURUSD","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"EURZAR","digits":4,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"EURZAR.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"EURZAR.vip","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPAUD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"GBPCAD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"GBPCHF","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"GBPCNH","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPCNH.raw","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPCNH.vip","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPCZK","digits":4,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPCZK.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPCZK.vip","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPDKK","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPDKK.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPDKK.vip","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPEUR","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPHKD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPHKD.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPHKD.vip","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPHUF","digits":4,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPHUF.raw","digits":3,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPHUF.vip","digits":3,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPJPY","digits":3,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"GBPMXN","digits":4,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPMXN.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPMXN.vip","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPNOK","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPNOK.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPNOK.vip","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPNZD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"GBPPLN","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPPLN.raw","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPPLN.vip","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPSEK","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPSEK.raw","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPSEK.vip","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPSGD.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPSGD.vip","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"GBPUSD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"GBPZAR","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPZAR.raw","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GBPZAR.vip","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"GOLDZ7","digits":1,"stopslevel":10,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":100},{"code":"HGCZ7","digits":4,"stopslevel":5,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":10000},{"code":"HK50","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":10},{"code":"HUFTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"JPN225","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":100},{"code":"NAS100","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1},{"code":"NOKTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"NZDCHF","digits":5,"stopslevel":5,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"NZDJPY","digits":3,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"NZDTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"NZDUSD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"PLNTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"SEKTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"SGDTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"SILVERZ7","digits":3,"stopslevel":10,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":5000},{"code":"SMI20","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1},{"code":"SP500","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1},{"code":"TRYJPY","digits":3,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"UK100","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1},{"code":"USDCAD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDCHF","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDCNH","digits":5,"stopslevel":20,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDCZK","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDDKK","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDHKD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"USDHUF","digits":3,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDJPY","digits":3,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDMXN","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDNGN","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000},{"code":"USDNOK","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDPLN","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDSEK","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDSGD","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"USDZAR","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100000},{"code":"WIGZ6","digits":1,"stopslevel":10,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1},{"code":"WS30","digits":1,"stopslevel":1,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":1},{"code":"WTI","digits":2,"stopslevel":10,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":100},{"code":"WTIV7","digits":2,"stopslevel":10,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":100},{"code":"WTIX7","digits":2,"stopslevel":10,"min_lots":0.1,"max_lots":100,"lots_step":0.1,"contractsize":100},{"code":"XAGUSD","digits":2,"stopslevel":4,"min_lots":0.1,"max_lots":100,"lots_step":0.01,"contractsize":5000},{"code":"XAUUSD","digits":2,"stopslevel":5,"min_lots":0.01,"max_lots":100,"lots_step":0.01,"contractsize":100},{"code":"ZARTRY","digits":5,"stopslevel":10,"min_lots":0.01,"max_lots":1000,"lots_step":0.1,"contractsize":100000}]}`

	ck := ckjson.NewCkj("TestCK", json)
	ck.JsonToStruct()
}
