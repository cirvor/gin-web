package main

import (
	"web/middleware"
	"web/route"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化引擎
	engine := gin.Default()
	err := engine.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}

	// 捕获全局异常
	engine.Use(middleware.HandleErrors())
	// 路由
	route.Web(engine)
	route.Websocket(engine)

	// 绑定端口并启动应用
	engine.Run(":9205")

	//endpoint := "oss1.zarsoso.com"
	//accessKeyID := "PEgsbijzc4qbuUr5"
	//secretAccessKey := "qKGvSpdrmLVoQTg8DMEFSR9fQ55RAaWH"
	//useSSL := true
	//
	//// 初使化 minio client对象。
	//minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Printf("success: %#v\n", minioClient) // minioClient初使化成功
	//
	//bucketName := "free"
	//objectName := "v0300f9c0000but36jq504rg256r3iq0.mp4"
	//filePath := "/Users/xjulien/Downloads/v0300f9c0000but36jq504rg256r3iq0.mp4"
	//contentType := "video/mp4"
	//
	//// 使用FPutObject上传一个zip文件。
	//n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Printf("Successfully uploaded %s of size %d\n", objectName, n)
}
