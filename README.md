### go文件上传模块

主要针对图片上传

- [x] 文件上传
- [x] 限制上传文件的后缀名
- [x] 限制上传文件的大小
- [x] Hash去重，防止重复上传
- [x] 图片自动生成缩略图
- [x] 全由配置

### 如何运行

```bash
go get -v github.com/axetroy/go-upload
cd $GOPATH/src/github.com/axetroy/go-upload
make build
./bin/server
```

### API

```bash
[POST]  /upload                                 # 图片上传POST方法
[GET]   /download/:size/:filename               # 获取的上传的图片, size为配置文件[upload.image, upload.thumbnail]
[GET]   /upload/example                         # 图片上传的demo，简单的表单post
```

### 配置文件

配置文件与二进制同一目录

```yaml
# config.yaml
# HTTP模块设置
http:
  host: localhost                               # 监听地址
  port: 9044                                    # 监听端口

# 上传模块的相关设置

upload:
  path: uploads                                 # 上传文件的保存目录名，相对于当前目录
  image: image                                  # path的子目录，图片保存的目录名，
  thumbnail: thumbnail                          # path的字目录，缩略图保存的目录名
  thumbnailmaxwidth: 300                        # 缩略图的最大宽度
  thumbnailmaxheight: 300                       # 缩略图的最大高度
  allowtype:                                    # 如果是空数组，那么会允许任意文件后缀名
    - .jpg
    - .jpeg
    - .png
    - .svg
    - .gif
    - .bmp
    - .ico
  maxsize: 10485760                             # 上传文件的最大大小，默认10M
```


## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Faxetroy%2Fgo-upload.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Faxetroy%2Fgo-upload?ref=badge_large)