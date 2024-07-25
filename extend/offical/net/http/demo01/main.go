package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"learn-go/common/funcs"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	host = "http://local.lewaimai.com"
)

// 测试加密签名
func main() {
	type articleExtend struct {
		Last_mod  int      `json:"last_mod"`
		First_mod int      `json:"first_mod"`
		Users     []string `json:"users"`
	}
	extend, err := json.Marshal(articleExtend{
		Last_mod:  int(time.Now().Unix()),
		First_mod: int(time.Now().Add(time.Hour*-24 + time.Minute*30).Unix()),
		Users:     []string{"lizhi", "lizhu"},
	})
	if err != nil {
		log.Fatalln("article extend json Mardahl error:", err)
	}
	postParms := map[string]string{
		"title":   "测试",
		"content": `真的就测试下"yes" what's your name? \t\s*&{}\`,
		"desc":    "666",
		"tags":    strings.Join([]string{"tennis", "pingpong"}, ","),
		"extend":  string(extend),
		"time":    strconv.FormatInt(time.Now().Unix(), 10),
		"nonce":   funcs.RandomStr(12),
	}

	sign := funcs.Sign(postParms)

	// 几种提交方式(post postform)及参数接受和验签 考虑特殊字符和大json
	// urlStr := funcs.BuildQuery(postParms, true) + "&sign=" + sign
	// log.Println("post string: ", urlStr)
	// resp, err := http.Post(host+"/test1.php", "application/x-www-form-urlencoded", strings.NewReader(urlStr))
	// if err != nil {
	// 	fmt.Printf("%s 请求错误 ...... \n", err)
	// 	return
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln("解析body错误", err)
	// }
	// log.Println("response body", string(body))

	// postform
	postParms["sign"] = sign
	resp, err := http.PostForm(host+"/test1.php", funcs.GetFormValue(postParms))
	if err != nil {
		fmt.Printf("%s 请求错误 ...... \n", err)
		return
	}
	defer resp.Body.Close()
	dealBody(*resp)

	// ``````````` php `````````````
	// 	echo "\n\n php return \n";
	// var_export($_POST);
	// if ($_POST) {
	//     ksort($_POST);
	//     $sign = $_POST['sign'];
	//     unset($_POST['sign']);
	//     $signStr = urldecode(http_build_query($_POST));
	//     if ($sign != md5(md5($signStr) . "3214e23sa12wsased34ewdxw1edwxs")) {
	//         echo "\n sign error: " . $signStr . "\n";
	//     } else {
	//         echo "\n sign success \n";
	//     }
	// }
	// ``````````` php `````````````

	// 参数处理
}

func dealBody(resp http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("解析body错误", err)
	}
	log.Println("\n\n response body: ", string(body))
}
