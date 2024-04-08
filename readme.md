# Claude Official Website Mirror (1:1 Version)

<div align="center">
  <a href="readme_zh.md">中文版 README</a>
  </br>
  <p>My communication channel :<a href="https://t.me/cliptalk">https://t.me/cliptalk</a>
</div>

### Project Introduction:

This is a 1:1 mirror of the Claude official website, completely replicated.

It is not afraid of official resource updates and automatically synchronizes with official resources.

Demo site: [https://www.atvai.com](https://www.atvai.com)

The demo site does not provide any technical support and is only for learning and exchange purposes. It is maintained casually.

### Project Deployment

#### 1.1 Prerequisites

1. Install 1panel or BT panel. If you are a server novice, it is recommended to install one of these. Here, I am using 1panel. (If you are an expert, you can directly use Nginx for reverse proxy.)
2. Install Docker and Docker Compose. (It is recommended to use Docker Compose. For specific installation instructions, please search on Google or consult GPT.)
3. Clone the project to your local machine.

```shell
git clone https://github.com/petiky/fkclaude.git
```

#### 1.2 Deployment

```shell
cd fkclaude
docker-compose up -d
```
or
```shell
./fkclaude
```
#### 1.3 Configuration

`I am using 1panel's reverse proxy here. If you are using BT panel, you can also use its reverse proxy. I won't go into details here.`

1. Create a new website, select reverse proxy, fill in your domain name, set the frontend request path to `/api`, name it `root` (you can name it anything), fill in your server address for the proxy address, set the port to `claude.ai`, select HTTPS, and save.
2. Configure an SSL certificate for your domain name. (I won't elaborate on this. If you don't know how, please search on Google.)
3. Modify the reverse proxy configuration file:
    - Click on "Reverse Proxy"
    - Click on "Create Reverse Proxy"
        - Fill in the name as "api", set the frontend request path to `/api`, set the backend reverse proxy address to `0.0.0.0:3650/api`, and click save.
        - Click on "Source"
            - Copy the following content into it:
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
              proxy_buffering off; # Disable proxy buffering
              chunked_transfer_encoding on; # Enable chunked transfer encoding
              keepalive_timeout 300; # Set keep-alive timeout to 65 seconds
              add_header X-Cache $upstream_cache_status;
              add_header Strict-Transport-Security "max-age=31536000";
              add_header Cache-Control no-cache;
          }
          ```
            - Click save
    - Click on the "Source" of the reverse proxy named "root":
        - Copy the following content into it:
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
      proxy_set_header Cf-Cache-Status $http_cf_cache_status;
      proxy_set_header User-Agent $http_user_agent; # Forward user's User-Agent
      proxy_set_header Accept $http_accept; # Forward user's Accept header
      proxy_set_header Accept-Language $http_accept_language; # Forward user's Accept-Language header
      proxy_set_header Referer $http_referer; # Forward user's Referer header
      proxy_set_header Cookie $http_cookie; # Forward user's Cookie
      proxy_set_header Connection "";
      proxy_set_header Cache-Control $http_cache_control;
      proxy_cookie_domain claude.ai $host; # Modify the domain name in the Set-Cookie header
      proxy_cookie_path / "/"; # Modify the path in the Set-Cookie header
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
        - Click save

Now, reopen your domain, and you should see an official website identical to Claude's. Registration, login, and conversation are all exactly the same.

### Precautions

1. Please do not modify any files in the project, as it may cause the project to malfunction.
2. Please do not modify any files in the project, as it may cause the project to malfunction.
3. If you encounter 403 or other errors, please check if your reverse proxy configuration is correct or if your IP is clean, among other factors.
4. Please do not directly use the login or registration interfaces for logging in or registering. There is a high probability of account bans. If you use them, please log in using sessionKey.
5. This project is for learning and exchange purposes only and should not be used for commercial purposes. Otherwise, you will bear the consequences.
6. This project is for learning and exchange purposes only and should not be used for commercial purposes. Otherwise, you will bear the consequences.
7. This project is for learning and exchange purposes only and should not be used for commercial purposes. Otherwise, you will bear the consequences.
8. This project is for learning and exchange purposes only and should not be used for commercial purposes. Otherwise, you will bear the consequences.

### Reward the Author

If you find this project helpful, you can reward me. Thank you!

USDT TRC address: `TBZdWC2y1b2DPLK6awnEfShUu7x9XRY4xp`
