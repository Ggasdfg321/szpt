package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/dop251/goja"
)

func readJS() string {
	path := "en.js"
	file, _ := os.Open(path)
	script, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	return string(script)
}

func getConfig(path string) map[string]string {
	config := make(map[string]string)
	f, _ := os.Open(path)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}

func getEncry() (string, string, string) {
	client := &http.Client{
		Jar: cookies_jar,
	}
	req, err := http.NewRequest(http.MethodGet, login_url, nil)

	if err != nil {
		fmt.Println("错误")
		return "err", "err", "err"
	}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	reg_encry := regexp.MustCompile(`var pwdDefaultEncryptSalt = "(.*)";`)
	reg_lt := regexp.MustCompile(`<input type="hidden" name="lt" value="(.*)"/>`)
	reg_execution := regexp.MustCompile(`<input type="hidden" name="execution" value="(.*)"/>`)
	if reg_encry == nil || reg_lt == nil || reg_execution == nil {
		fmt.Println("regexp err")
		return "err", "err", "err"
	}
	encry := reg_encry.FindStringSubmatch(string(body))
	lt := reg_lt.FindStringSubmatch(string(body))
	execution := reg_execution.FindStringSubmatch(string(body))
	return encry[len(encry)-1], lt[len(lt)-1], execution[len(execution)-1]
}

const (
	login_url       string = "https://authserver.szpt.edu.cn/authserver/login"
	getCookie_url   string = "https://ehall.szpt.edu.cn/publicappinternet/sys/szptpubxsjkxxbs/index.do?t_s=1646488570022#/mrxxbs"
	getInfo_url     string = "https://ehall.szpt.edu.cn/publicappinternet/sys/szptpubxsjkxxbs/mrxxbs/getSaveReportInfo.do"
	autoTianBao_url string = "https://ehall.szpt.edu.cn/publicappinternet/sys/szptpubxsjkxxbs/mrxxbs/saveReportInfo.do"
)

var encrypt func(string, string) string

func runJS() {
	vm := goja.New()
	_, err := vm.RunString(readJS())
	if err != nil {
		fmt.Println("JS代码有问题！")
		return
	}
	err = vm.ExportTo(vm.Get("encryptAES"), &encrypt)
	if err != nil {
		fmt.Println("JS函数映射失败")
		return
	}
}

func Login(username string, password string, encry_pwd string, lt string, execution string) {
	pwd_en := encrypt(password, encry_pwd)
	values := strings.NewReader(url.Values{"username": {username}, "password": {pwd_en}, "lt": {lt}, "dllt": {"userNamePasswordLogin"}, "execution": {execution}, "_eventId": {"submit"}, "rmShown": {"1"}}.Encode())
	client := &http.Client{
		Jar: cookies_jar,
	}
	req, _ := http.NewRequest(http.MethodPost, login_url, values)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}

//
var cookies_jar *cookiejar.Jar

func main() {

	cookies_jar, _ = cookiejar.New(nil)
	runJS()
	encry_pwd, lt, execution := getEncry()
	conf := getConfig("config.ini")
	fmt.Println(conf["username"], conf["password"])
	Login(conf["username"], conf["password"], encry_pwd, lt, execution)

	func() {
		client := &http.Client{
			Jar: cookies_jar,
		}
		resp, _ := client.Get(getCookie_url)
		defer resp.Body.Close()
		resp, _ = client.Get(getInfo_url)
		data, _ := ioutil.ReadAll(resp.Body)
		tmp_json := func() map[string]interface{} {
			var tmp_jsonu = []byte(`{"WID": "", "REPORT_DATE": "", "USER_ID": "", "USER_NAME": "", "DEPT_CODE": "", "DEPT_NAME": "", "XB": "", "NL": "", "SJHM": "", "JG": "", "DWDZ": "", "XZZSF_DISPLAY": "", "XZZSF": "", "XZZCS_DISPLAY": "", "XZZCS": "", "XZZDQ_DISPLAY": "", "XZZDQ": "", "XZZXXDZ": "", "JTZZSF_DISPLAY": "", "JTZZSF": "", "JTZZCS_DISPLAY": "", "JTZZCS": "", "JTZZDQ_DISPLAY": "", "JTZZDQ": "", "JTZZXXDZ": "", "JZSJHM": "", "SFWZXS_DISPLAY": "", "SFWZXS": "", "SSXQ_DISPLAY": "", "SSXQ": "", "SSLD_DISPLAY": "", "SSLD": "", "SSFJH_DISPLAY": "", "SSFJH": "", "ZRCXFHJZXCLX": "", "JRFXXCLX": "", "ZSDZ": "", "BJ": "", "FDYXM": "", "FDYSJHM": "", "BZRXM": "", "BZRSJHM": "", "SFZZCJSX_DISPLAY": "", "SFZZCJSX": "", "SXFS_DISPLAY": "", "SXFS": "", "SFZZSXDWSS_DISPLAY": "", "SFZZSXDWSS": "", "SFYFS_DISPLAY": "", "SFYFS": "", "FSSJ": "", "SFYFX_DISPLAY": "", "SFYFX": "", "FXSJ": "", "QYTZWTW": "", "QYTWSTW": "", "DTZSTW": "", "SFJFFH_DISPLAY": "", "SFJFFH": "", "FHTJGJ": "", "QTXYSMDJWQK": "", "CQWYQFKZDDQ_DISPLAY": "", "CQWYQFKZDDQ": "", "SSSQ": "", "XSQBDSJ": "", "JCGYQFKZDDQGWRY_DISPLAY": "", "JCGYQFKZDDQGWRY": "", "HSJCBG": "", "JCGYSHQZBL_DISPLAY": "", "JCGYSHQZBL": "", "GTSHQSGQYWYQFKZDDQLJS_DISPLAY": "", "GTSHQSGQYWYQFKZDDQLJS": "", "JTCYJKQK_DISPLAY": "", "JTCYJKQK": "", "STJKZK_DISPLAY": "", "STJKZK": "", "STYCZK": "", "STYXZK_DISPLAY": "", "STYXZK": "", "SFJSJJGC_DISPLAY": "", "SFJSJJGC": "", "JSJJGCJTSJ": "", "SFJSJTGC_DISPLAY": "", "SFJSJTGC": "", "JSJTGCJTSJ": "", "JSJJJTGCYY": "", "SFWZZGRZ_DISPLAY": "", "SFWZZGRZ": "", "SFJXHSJC_DISPLAY": "", "SFJXHSJC": "", "ZJYCHSJCSJ": "", "HSJCJG_DISPLAY": "", "HSJCJG": "", "XGYMJZQK_DISPLAY": "", "XGYMJZQK": "", "XGYMJZJJ": "", "SFYYYXGYMJZ_DISPLAY": "", "SFYYYXGYMJZ": "", "FXZB1_DISPLAY": "", "FXZB1": "", "FXZB2_DISPLAY": "", "FXZB2": "", "FXZB3_DISPLAY": "", "FXZB3": "", "FXZB4_DISPLAY": "", "FXZB4": "", "FXZB5_DISPLAY": "", "FXZB5": "", "FXZB6_DISPLAY": "", "FXZB6": "", "FXZB7_DISPLAY": "", "FXZB7": "", "FXZB8_DISPLAY": "", "FXZB8": "", "FXZB9_DISPLAY": "", "FXZB9": "", "BZ": ""}`)
			var tmp_json map[string]interface{}
			_ = json.Unmarshal(tmp_jsonu, &tmp_json)

			return tmp_json
		}()

		szpt_json := func(data string) map[string]interface{} {
			var tmp_jsonu = []byte(data)
			var tmp_json map[string]interface{}
			_ = json.Unmarshal(tmp_jsonu, &tmp_json)
			return tmp_json
		}(string(data))

		func() {
			datas := szpt_json["datas"].(map[string]interface{})
			for i, j := range datas {
				tmp_json[i] = j
			}
			return_json, _ := json.Marshal(tmp_json)
			values := strings.NewReader(url.Values{"formData": {string(return_json)}}.Encode())
			client := &http.Client{
				Jar: cookies_jar,
			}
			req, _ := http.NewRequest(http.MethodPost, autoTianBao_url, values)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			resp, _ := client.Do(req)
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(data))
		}()
	}()
}
