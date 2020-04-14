package main

import (
	"Login/Utils"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var COOKIES []*http.Cookie

var USERAGENT = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36`

func main() {

	username := ""
	password := ""
	randnumber := ""

	key, nowtime := getDeskeyNowtime()

	sessionid := COOKIES[0].Value

	v := url.Values{}
	v.Add("dateTime", time.Now().Format("2006/01/02 15:04:05"))
	get_pic("http://211.86.128.194/suzxyjw/cas/genValidateCode?" + v.Encode())

	fmt.Scan(&randnumber)

	var txt_mm_expression = get_txt_mm_expression(password)
	var txt_mm_length = get_txt_mm_length(password)
	var txt_mm_userzh = get_txt_mm_userzh(username, password)

	password = md5sum(md5sum(password) + md5sum(strings.ToLower(randnumber)))

	var p_username = "_u" + randnumber
	var p_password = "_p" + randnumber

	username = base64_encode(username + ";;" + sessionid)

	var params = p_username + "=" + username +
		"&" + p_password + "=" + password +
		"&randnumber=" + randnumber +
		"&isPasswordPolicy=" + "1" +
		"&txt_mm_expression=" + txt_mm_expression +
		"&txt_mm_length=" + txt_mm_length +
		"&txt_mm_userzh=" + txt_mm_userzh
	params, token, timestamp := getEnparams(params, key, nowtime)
	login("http://211.86.128.194/suzxyjw/cas/logon.action", params, token, timestamp, sessionid)
}

// 获取验证码
func get_pic(urls string) {

	resp, err := grequests.Get(urls, &grequests.RequestOptions{Cookies: COOKIES})
	if err != nil {
		log.Println(err)
	}
	err = resp.DownloadToFile("./code.png")
	if err != nil {
		log.Println(err)
	}
}

// 登陆
func login(urls, params, token, timestamp, sessionid string) {

	resp, err := grequests.Post("http://211.86.128.194/suzxyjw/cas/logon.action", &grequests.RequestOptions{
		UserAgent: USERAGENT,
		Cookies:   COOKIES,
		Headers: map[string]string{
			"Referer": "http://211.86.128.194/suzxyjw/cas/login.action",
		},
		IsAjax: true,
		Data: map[string]string{
			"params":    params,
			"token":     token,
			"timestamp": timestamp,
		}})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
	if gjson.GetBytes(resp.Bytes(), "status").Int() == 200 {
		fmt.Println(strings.Split(COOKIES[0].String(), ";")[0])
		fmt.Println(gjson.GetBytes(resp.Bytes(), "message"), gjson.GetBytes(resp.Bytes(), "result"))
		//getkb(2)
	}
}

//// 获取课表
//func getkb(weeks int) {
//	resp, err := grequests.Post("http://211.86.128.194/suzxyjw/frame/desk/showLessonScheduleDetail.action", &grequests.RequestOptions{
//		Cookies:   COOKIES,
//		UserAgent: USERAGENT,
//		Data: map[string]string{
//			"weeks": strconv.Itoa(weeks),
//		},
//		Headers: map[string]string{
//			"Content-Type": "text/html; charset=utf-8",
//		}})
//
//	query, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Bytes()))
//	if err != nil {
//		log.Println(err)
//	}
//	query.Find("tbody").Find(".mykb").Find("td[style*=background-color]").Each(func(i int, selection *goquery.Selection) {
//		fmt.Println(i, selection.Find("ul").Text())
//	})
//}

func getEnparams(data, key, nowtime string) (string, string, string) {

	token := md5sum(md5sum(data) + md5sum(nowtime))
	var params = base64_encode(utf16to8(Utils.StrEnc(data, key)))
	return params, token, nowtime
}

// 获取个参数
func getDeskeyNowtime() (string, string) {

	resp, err := grequests.Get("http://211.86.128.194/suzxyjw/cas/login.action", &grequests.RequestOptions{
		UserAgent: USERAGENT,
		Headers: map[string]string{
			"Referer": "http://www.ahszu.edu.cn/",
		},
	})
	if err != nil {
		log.Println(err)
	}
	COOKIES = resp.RawResponse.Cookies()

	regexp, err := regexp.Compile(`/suzxyjw/custom/js/SetKingoEncypt.jsp\?t=[0-9]+`)
	if err != nil {
		log.Println(err)
	}
	url_sufix := regexp.FindString(resp.String())

	resp, err = grequests.Get("http://211.86.128.194"+url_sufix, &grequests.RequestOptions{
		Cookies:   COOKIES,
		UserAgent: USERAGENT,
	})
	if err != nil {
		log.Println(err)
	}

	html2str := resp.String()

	var FindTwoIndex = func(str string, subchar int32) int {
		w := 0
		for i, itoa := range str {
			if itoa == subchar {
				w++
				if w == 2 {
					return i
				}
			}
		}
		return -1
	}

	var splitx = func(str string) string {
		temp_str := strings.Split(str, "=")[1]
		// 去空格
		temp_str = strings.TrimSpace(temp_str)
		// 去符号
		temp_str = strings.ReplaceAll(temp_str, "'", "")
		return temp_str
	}
	// 两部分 keys, nowtime
	two_parts := strings.Split(html2str[:FindTwoIndex(html2str, ';')], ";")
	keys := splitx(two_parts[0])
	nowtime := splitx(two_parts[1])
	return keys, nowtime
}

//utf16to8
func utf16to8(str string) string {
	var out string
	for _, itoa := range str {
		if itoa >= 0x0001 && itoa <= 0x007f {
			out += string(itoa)
		} else if itoa > 0x07ff {
			out += string(0xe0 | ((itoa >> 12) & 0x0f))
			out += string(0x80 | ((itoa >> 6) & 0x3f))
			out += string(0x80 | ((itoa >> 0) & 0x3f))
		} else {
			out += string(0xc0 | ((itoa >> 6) & 0x1f))
			out += string(0x80 | ((itoa >> 0) & 0x3f))
		}
	}
	return out
}

// 密码规则
func get_txt_mm_expression(x string) string {
	result := 0
	charType := func(x byte) int {
		if x >= 48 && x <= 57 {
			return 8
		}
		if x >= 97 && x <= 122 {
			return 4
		}
		if x >= 65 && x <= 90 {
			return 2
		}
		return 1
	}

	for _, itoa := range x {
		result |= charType(byte(itoa))
	}
	return strconv.Itoa(result)
}

// 密码长度
func get_txt_mm_length(x string) string {
	return fmt.Sprintf("%v", len(x))
}

// 判断密码是否在账号里
func get_txt_mm_userzh(user, pass string) string {
	s_user := strings.TrimSpace(strings.ToLower(user))
	s_pass := strings.TrimSpace(strings.ToLower(pass))
	result := strings.Index(s_pass, s_user)
	if result > -1 {
		return "1"
	}
	return "0"

}

// base64 encode
func base64_encode(x interface{}) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v", x)))
}

// base64_decode
func base64_decode(x string) string {
	result, err := base64.StdEncoding.DecodeString(x)
	if err != nil {
		log.Println(err)
	}
	return string(result)
}

// md5sum
func md5sum(x interface{}) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%v", x))))
}
