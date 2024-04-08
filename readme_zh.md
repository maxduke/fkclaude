### 项目介绍：
claude 官网镜像 1:1版本，完全复刻～

不惧怕官方资源更新，自动同步官方资源

演示站点: [https://www.atvai.com](https://www.atvai.com)

演示站点不做任何技术支持，仅供学习交流使用，随缘维护～

tg 交流频道 ：[https://t.me/cliptalk](https://t.me/cliptalk)

### 项目部署
#### 1.1 准备工作
1. 安装 1panel、宝塔也可以，如果你是服务器小白，建议安装一下，这里我用的是 1panel（如果你是大佬，可以直接使用 nginx 来做反向代理～）
2. 安装 docker 和 docker-compose （这里建议使用 docker-compose，具体怎么安装，请自行谷歌或者gpt）
3. 将项目克隆到本地
```shell
git clone https://github.com/petiky/fkclaude.git
```
#### 1.2 部署
```shell
cd fkclaude
docker-compose up -d
```
或者
```shell
./fkclaude
```
#### 1.3 配置
`我这里使用的 1panel 的反向代理，如果你使用的是宝塔，也可以使用宝塔的反向代理，这里就不多说了`
1. 新建一个网站，选择反向代理，填写你的域名，前端请求路径填写`/api`，名称填写 `root` (你也可以随便写) 代理地址填写你的服务器地址，端口填写 `claude.ai`,选择 https，然后保存
2. 给你的域名配置 ssl 证书（不多赘述，不懂的自行谷歌）
   3. 修改反向代理配置文件
      - 点击反向代理
      - 点击创建反向代理
        - 名称填写 api，前端请求路径填写`/api` 后端反代地址填写 `0.0.0.0:3650/api` 点击保存即可
        - 点击源文
          - 将以下内容复制进去
       ```shell
       location ^~ /api {
           proxy_pass http://0.0.0.0:3650/api; 
           proxy_set_header Host $host; 
           proxy_set_header X-Real-IP $remote_addr; 
           proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for; 
           proxy_set_header Upgrade $http_upgrade; 
           proxy_set_header Connection "upgrade"; 
           proxy_set_header X-Forwarded-Proto $scheme; 
           proxy_http_version 1.1; 
           proxy_buffering off; # 关闭代理缓冲
           chunked_transfer_encoding on; # 开启分块传输编码
           keepalive_timeout 300; # 设定keep-alive超时时间为65秒
           add_header X-Cache $upstream_cache_status; 
           add_header Strict-Transport-Security "max-age=31536000"; 
           add_header Cache-Control no-cache; 
       }
       ```
         - 点击保存
      - 点击名称为 root 的反向代理的源文：
           - 将一下内容复制进去
      ```shell
      location ^~ / {
      proxy_pass https://claude.ai;
      proxy_set_header Host claude.ai;
      proxy_set_header X-Real-IP $remote_addr;
      proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
      proxy_set_header X-Forwarded-Proto $scheme;
      proxy_set_header X-Forwarded-Host claude.ai;
      proxy_set_header X-Forwarded-Port 443;
      proxy_set_header X-Cloudflare-Proto https;
      proxy_set_header Cf-Cache-Status DYNAMIC;
      proxy_set_header CF-Visitor '{"scheme":"https"}';
      proxy_set_header CF-IPCountry US;
      #proxy_set_header CF-Connecting-IP $remote_addr;  # 使用用户的IP
      proxy_set_header Cf-Cache-Status $http_cf_cache_status;
      proxy_set_header User-Agent $http_user_agent; # 转发用户的 User-Agent
      proxy_set_header Accept $http_accept; # 转发用户的 Accept 头
      proxy_set_header Accept-Language $http_accept_language; # 转发用户的 Accept-Language 头
      proxy_set_header Referer $http_referer; # 转发用户的 Referer 头
      proxy_set_header Cookie $http_cookie; # 转发用户的 Cookie
      proxy_set_header Connection "";
      proxy_set_header Cache-Control $http_cache_control;
      proxy_cookie_domain claude.ai $host; # 修改 Set-Cookie 头中的域名
      proxy_cookie_path / "/"; # 修改 Set-Cookie 头中的路径
      proxy_set_header Sec-Ch-Ua $http_sec_ch_ua;
      proxy_set_header Sec-Ch-Ua-Mobile $http_sec_ch_ua_mobile;
      proxy_set_header Sec-Ch-Ua-Platform $http_sec_ch_ua_platform;
      proxy_set_header Sec-Fetch-Mode $http_sec_fetch_mode;
      proxy_set_header Sec-Fetch-Site $http_sec_fetch_site;
      proxy_set_header Upgrade-Insecure-Requests $http_upgrade_insecure_requests;
      if ($request_method = 'OPTIONS' ) {
      add_header 'Access-Control-Allow-Origin' '*';
      add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS';
      add_header 'Access-Control-Allow-Headers' 'DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range';
      add_header 'Access-Control-Max-Age' 1728000;
      add_header 'Content-Type' 'text/plain; charset=utf-8';
      add_header 'Content-Length' 0;
      return 200;
      }
      add_header 'Access-Control-Allow-Origin' '*';
      add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS';
      add_header 'Access-Control-Allow-Headers' 'DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range';
      add_header 'Access-Control-Expose-Headers' 'Content-Length,Content-Range';
      proxy_set_header CF-RAY $http_cf_ray;
      proxy_set_header Upgrade $http_upgrade;
      proxy_set_header Connection "upgrade";
      proxy_http_version 1.1;
      proxy_ssl_server_name on;
      proxy_ssl_protocols TLSv1.2 TLSv1.3;
      proxy_ssl_ciphers EECDH+AESGCM:EDH+AESGCM;
      add_header Strict-Transport-Security "max-age=31536000";
      add_header Cache-Control no-cache;
      }
      ```
       - 点击保存
      
现在重新打开你的域名，应该就可以看到 与claude一模一样的官网了，注册，登录，对话 什么都是一模一样的～

### 注意事项
1. 请不要修改项目中的任何文件，否则可能会导致项目无法正常运行
2. 请不要修改项目中的任何文件，否则可能会导致项目无法正常运行
3. 如果出现 403 等错误，请检查你的反向代理配置是否正确或者你的 ip 是否干净等等因素
4. 请不要直接使用登录接口以及注册接口进行登录或者注册，大概率会直接封号，如果使用，请自行使用 sessionKey 登录～
5. 本项目仅供学习交流使用，不得用于商业用途，否则后果自负
6. 本项目仅供学习交流使用，不得用于商业用途，否则后果自负
7. 本项目仅供学习交流使用，不得用于商业用途，否则后果自负
8. 本项目仅供学习交流使用，不得用于商业用途，否则后果自负

### 打赏作者
如果你觉得这个项目对你有帮助，可以给我打赏一下，谢谢～

USDT TRC地址：`TBZdWC2y1b2DPLK6awnEfShUu7x9XRY4xp`