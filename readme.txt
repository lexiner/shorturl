常用命令

运行项目：bee run 
打包项目:   bee pack （指定系统，-be GOOS=windows/linux）

下载依赖包： go get https://*********
安装依赖包:   go mod tidy

安装注意事项:

conf/app.go  配置数据库mysql信息 （数据库表在根目录，需要提前导入）

创建短链接接口：

/create  post（json）
参数：longurl （长链接）

域名也在app.go里修改为自己域名
