Open-Falcon微信告警组件
---------------------------------------------------
在此
感谢 @laiwei 的https://github.com/open-falcon/mail-provider 项目代码支持		
感谢 @chanxuehong 老司机的微信SDK支持  

##  申请微信企业号
在https://work.weixin.qq.com/ 注册微信企业号，不需要认证即可
### 配置企业号
在企业号里建立应用,依此点击：企业应用-创建应用。可见范围里添加成员
获取企业ID：点击我的企业--获取企业ID
![1](https://img.cactifans.com/wp-content/uploads/2018/07/2.jpg)
获取应用AgentId及应用Secret
![2](https://img.cactifans.com/wp-content/uploads/2018/07/1.jpg)
##  部署Falcon-wechat
 Falcon-wechat可部署在Falcon-Alarm机器，也可部署在独立机器，使用以下命令部署
```bash
 mkdir -p /usr/local/falcon-wechat
 cd /usr/local/falcon-wechat
 wget https://dl.cactifans.com/open-falcon/falcon-wechat-0.0.1.tar.gz
 tar zxvf falcon-wechat-0.0.1.tar.gz
```
修改配置文件cfg.json配置文件,将corpid修改为你的企业ID，secret修改为你应用的secret，agentid修改为你的AgentId，并保存
```json
{
        "debug": true,
        "http": {
                "listen": "0.0.0.0:9527",
                "token": ""
        },
        "wechat": {
	        "corpid": "wxa7c63522727b6bf0",
	        "secret": "d5S-_XGVd-5HA0o9bkvMMK5Wh1qwCC-YQei4WcL9hSM",
	        "agentid": 1
    }
}
```
启动falcon-wechat
```bash
./control start
```
启动信息
>falcon-wechat started..., pid=13875

启动之后可使用以下命令测试是否成功
```bash
curl -d "tos=zhanghao&content=内容" "http://127.0.0.1:9527/wechat"
```
>tos 后面为用户账号
>content 后面为内容

如看到success表示发送成功，如遇到错误可使用以下命令查看日志
查看日志

```
./control tail
```
如看到以下信息表示成功
>2018/07/19 14:44:28 config.go:64: load configuration file cfg.json successfully
2018/07/19 14:44:28 http.go:25: http listening 0.0.0.0:9527

常用命令

>./control start 启动程序     
>./control stop  停止程序     
>./control build 从源码编译程序     
>./control pack  打包程序   

##  配置Open-Falcon
### 配置Alarm组件
修改Open-Falcon的Alarm组件config目录下的配置文件cfg.json，将IM段修改为以下内容
```json
	"im": "http://127.0.0.1:9527/wechat",
```
>如果你修改了falcon-wechat的默认端口，请注意修改。如falcon-wechat和Alarm组件为不同机器，注意修改IP地址。
修改之后重启Alarm服务，使其生效
### 配置IM信息
在Dashboard里，为用户配置IM账号为户**用户账号**！**用户账号**！
**用户账号** 不是微信号，重要事情说三遍！
![3](https://img.cactifans.com/wp-content/uploads/2018/07/3.jpg)
用户账号
![4](https://img.cactifans.com/wp-content/uploads/2018/07/4.jpg)
## 效果
告警
![5](https://img.cactifans.com/wp-content/uploads/2018/07/5.jpg)
微信
![6](https://img.cactifans.com/wp-content/uploads/2018/07/6.png)

##  注意事项
1.由于使用企业微信发送消息接口实现，接口调用速率有限制。注意控制消息发送频率，目前每次发消息都会请求一次Access_token，后续优化。
2.由于未认证企业发送消息数量与人员有关，建议控制频率。具体查看官网API文档https://work.weixin.qq.com/api/doc#10785