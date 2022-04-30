import js2py,requests,re,urllib3,json
from urllib.parse import unquote
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
def password_en(pwd,en):
    js_obj=js2py.EvalJs()
    with open("en.js","r",encoding='utf-8') as f:
        js_obj.execute(f.read())
    return js_obj.encryptAES(pwd,en)

def main(user,pwd):
    username=user
    password=pwd
    url=requests.session()
    url_login="https://authserver.szpt.edu.cn/authserver/login"
    a=url.get(url_login,verify=False)
    encry=re.search('var pwdDefaultEncryptSalt = "(.*)";',a.text).group(1)
    encry_pwd=password_en(password,encry)
    lt=re.search('<input type="hidden" name="lt" value="(.*)"/>',a.text).group(1)
    execution=re.search('<input type="hidden" name="execution" value="(.*)"/>',a.text).group(1)

    data={"username":username,
        "password":encry_pwd,
        "lt":lt,
        "dllt":"userNamePasswordLogin",
        "execution":execution,
        "_eventId":"submit",
        "rmShown":"1"
        }

    url.post(url_login,data=data,verify=False).text
    url.get('https://ehall.szpt.edu.cn/publicappinternet/sys/szptpubxsjkxxbs/index.do?t_s=1646488570022#/mrxxbs',verify=False)
    tijiao=json.loads('{"WID": "", "REPORT_DATE": "", "USER_ID": "", "USER_NAME": "", "DEPT_CODE": "", "DEPT_NAME": "", "XB": "", "NL": "", "SJHM": "", "JG": "", "DWDZ": "", "XZZSF_DISPLAY": "", "XZZSF": "", "XZZCS_DISPLAY": "", "XZZCS": "", "XZZDQ_DISPLAY": "", "XZZDQ": "", "XZZXXDZ": "", "JTZZSF_DISPLAY": "", "JTZZSF": "", "JTZZCS_DISPLAY": "", "JTZZCS": "", "JTZZDQ_DISPLAY": "", "JTZZDQ": "", "JTZZXXDZ": "", "JZSJHM": "", "SFWZXS_DISPLAY": "", "SFWZXS": "", "SSXQ_DISPLAY": "", "SSXQ": "", "SSLD_DISPLAY": "", "SSLD": "", "SSFJH_DISPLAY": "", "SSFJH": "", "ZRCXFHJZXCLX": "", "JRFXXCLX": "", "ZSDZ": "", "BJ": "", "FDYXM": "", "FDYSJHM": "", "BZRXM": "", "BZRSJHM": "", "SFZZCJSX_DISPLAY": "", "SFZZCJSX": "", "SXFS_DISPLAY": "", "SXFS": "", "SFZZSXDWSS_DISPLAY": "", "SFZZSXDWSS": "", "SFYFS_DISPLAY": "", "SFYFS": "", "FSSJ": "", "SFYFX_DISPLAY": "", "SFYFX": "", "FXSJ": "", "QYTZWTW": "", "QYTWSTW": "", "DTZSTW": "", "SFJFFH_DISPLAY": "", "SFJFFH": "", "FHTJGJ": "", "QTXYSMDJWQK": "", "CQWYQFKZDDQ_DISPLAY": "", "CQWYQFKZDDQ": "", "SSSQ": "", "XSQBDSJ": "", "JCGYQFKZDDQGWRY_DISPLAY": "", "JCGYQFKZDDQGWRY": "", "HSJCBG": "", "JCGYSHQZBL_DISPLAY": "", "JCGYSHQZBL": "", "GTSHQSGQYWYQFKZDDQLJS_DISPLAY": "", "GTSHQSGQYWYQFKZDDQLJS": "", "JTCYJKQK_DISPLAY": "", "JTCYJKQK": "", "STJKZK_DISPLAY": "", "STJKZK": "", "STYCZK": "", "STYXZK_DISPLAY": "", "STYXZK": "", "SFJSJJGC_DISPLAY": "", "SFJSJJGC": "", "JSJJGCJTSJ": "", "SFJSJTGC_DISPLAY": "", "SFJSJTGC": "", "JSJTGCJTSJ": "", "JSJJJTGCYY": "", "SFWZZGRZ_DISPLAY": "", "SFWZZGRZ": "", "SFJXHSJC_DISPLAY": "", "SFJXHSJC": "", "ZJYCHSJCSJ": "", "HSJCJG_DISPLAY": "", "HSJCJG": "", "XGYMJZQK_DISPLAY": "", "XGYMJZQK": "", "XGYMJZJJ": "", "SFYYYXGYMJZ_DISPLAY": "", "SFYYYXGYMJZ": "", "FXZB1_DISPLAY": "", "FXZB1": "", "FXZB2_DISPLAY": "", "FXZB2": "", "FXZB3_DISPLAY": "", "FXZB3": "", "FXZB4_DISPLAY": "", "FXZB4": "", "FXZB5_DISPLAY": "", "FXZB5": "", "FXZB6_DISPLAY": "", "FXZB6": "", "FXZB7_DISPLAY": "", "FXZB7": "", "FXZB8_DISPLAY": "", "FXZB8": "", "FXZB9_DISPLAY": "", "FXZB9": "", "BZ": ""}')
    baocun=json.loads(url.post("https://ehall.szpt.edu.cn/publicappinternet/sys/szptpubxsjkxxbs/mrxxbs/getSaveReportInfo.do",verify=False).text)
    for i in baocun['datas']:
        if i in tijiao:
            tijiao[i]=baocun['datas'][i]
    tijiao=json.dumps(tijiao)
    data={"formData":tijiao}

    print(url.post("https://ehall.szpt.edu.cn/publicappinternet/sys/szptpubxsjkxxbs/mrxxbs/saveReportInfo.do",data=data,verify=False).text)
    
if __name__ == "__main__":
    username="" #填写你的学号
    password="" #填写你的密码
    main(username,password)
