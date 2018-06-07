### go文件上传模块

基于[github.com/axetroy/gin-uploader](https://github.com/axetroy/gin-uploader)

- [x] 文件上传
- [x] 限制上传文件的后缀名
- [x] 限制上传文件的大小
- [x] Hash去重，防止重复上传
- [x] 图片自动生成缩略图
- [x] 全由配置

### 如何运行

```bash
go get -v -u github.com/axetroy/go-upload
cd $GOPATH/src/github.com/axetroy/go-upload
make build
./bin/server
# 或者以生产环境运行
GO_ENV=production ./bin/server
```

### API

```bash

# 上传相关
[POST]  /upload/image                           # 图片上传
[POST]  /upload/file                            # 其他文件上传
[GET]   /upload/example                         # 上传demo，仅在开发模式下
# 下载相关
[GET]   /download/image/origin/:filename        # 获取上传的原始图片
[GET]   /download/image/thumbnail/:filename     # 获取上传的缩略图片
[GET]   /download/file/download/:filename       # 下载文件
[GET]   /download/file/raw/:filename            # 获取文件
```

### 配置文件

配置文件与二进制同一目录

```yaml
# config.yaml

# 在GO中，所有属性绑定到结构体，都是小写

# HTTP模块设置
http:
  host: localhost                               # 监听地址
  port: 9044                                    # 监听端口

# 上传模块的相关设置

upload:
  path: uploads                                 # 文件上传的根目录
  # 普通文件上传
  file:
    path: files                                 # 文件上传的目录
    maxsize: 10485760                           # 上传文件的最大大小
    allowtype:                                  # 允许上传的文件后缀名
      - .log
      - .txt
      - .text
      - .md
  # 图片上传
  image:
    path: image                                 # 图片上传的目录
    maxsize: 10485760                           # 上传图片的最大大小
    thumbnail:
      path: thumbnail                           # 缩略图存放的目录
      maxwidth: 300                             # 缩略图最大宽度
      maxheight: 300                            # 缩略图最大高度
```


## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faxetroy%2Fgo-upload.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Faxetroy%2Fgo-upload?ref=badge_large)
