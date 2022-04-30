# 深职院自动健康填报助手

## 使用方法：
### Python语言实现（使用python3运行）：
  #### 第一步、安装模块：  
    pip3 install js2py
    pip3 install requests
    pip3 install urllib3

  #### 第二步、在2.py里填写学号和密码  
    
  #### 第三步运行代码  
    python3 2.py 

### Go语言实现
  #### 第一步安装包  
    go get github.com/dop251/goja

  #### 第二步编译  
    go build main.go

  #### 第三步修改config.ini配置文件  
    在config.ini填写账号和密码
  
  #### 第四步运行  
    Windows：main.exe
    Linux: chmod +x main;./main

GO文件里面有编译好的文件，可以直接使用  
如果运行的时候显示{"datas":1,"code":"0"}，则表示已经填报成功，显示其他东西则填报失败  
### 实现无人工操作定时填报  
方法一：在Linux上可以使用crontab来实现每日填报  
方法二：使用screen命令  
# 声明  
### 项目用于学习交流，仅用于各项无异常时打卡，建议如实打卡  
### 使用此脚本进行违法行为与制作人无关，相关后果自负
