// web 开发 net.http
// 1. 注册路由处理 （静态资源可以定向到指定目录）
// 2. 创建服务监听

http1.0 一个tcp链接只能发送一个请求
http1.1 一个tcp链接可以发送多个请求 Connection： keep-alive
http2.0 （1）协议解析由文本改为二进制 （2) 服务端推送  (3) 与1.1不同，不是排队发送，而是并发 （4) 头部压缩

// 请求处理
// request  请求行  请求头  请求体
// response 响应行  响应头  响应体

// 用户识别及状态记录 userData (文件， 数据库，redis, map)
// 1. session -> uid -> userData
// 2. post -> token + sign -> userData
// 3. jwt header + userdata + sign