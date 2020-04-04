// jsonrpc(socket连接), rpc(go 之间调用)， grpc
https://studygolang.com/articles/14336

########################## 安装  ######################

// 下载指定平台的解码器放到go的bin目录下
https://github.com/protocolbuffers/protobuf/releases
# liunx 下
wget https://github.com/google/protobuf/archive/v3.5.0.tar.gz
tar -zxvf v3.5.0.tar.gz
cd protobuf-3.5.0
./autogen.sh
./configure
make
make check
make install

// 安装 grpc
cd 到 $GOPATH\src
git clone https://github.com/grpc/grpc-go.git google.golang.org/grpc
git clone https://github.com/golang/net.git golang.org/x/net
git clone https://github.com/golang/text.git golang.org/x/text
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git google.golang.org/genproto

go install google.golang.org/grpc

########################## 使用  ######################

// protoc --go_out=plugins=grpc:{输出目录}  {proto文件}
// plugins可以替换为其他插件 如 protorpc
// proto文件语法详解参阅：https://blog.csdn.net/u014308482/article/details/52958148
cd 到 grpc
protoc --go_out=plugins=grpc:./ ./test/test.proto


########################## 注意  ######################

1. 基于 tcp/http2 协议 , Protocol Buffer 3数据序列化协议
2. proto 文件为基础规范 （各个服务的接口请求、返回的数据格式），编译为各种语言，实现跨语言调用
3. 服务端：
    1. 创建 tcp 服务监听
    2. 注册服务 (login, register, test ...)： 注册合法的函数或者类到 grpc 服务上 （请求参数固定， 返回参数必须是proto定义的数据结构指针）
    3. 在gRPC服务器上注册反射服务,将监听交给gRPC服务处理
4. 客户端
    1. 建立 grpc 连接
    2. 创建服务 (login, register, test ...) 客户端
    3. 调用gRPC接口