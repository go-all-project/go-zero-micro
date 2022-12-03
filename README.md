## 命令

### rpc
* 创建 rpc
  ```shell
  goctl rpc new rpc
  ```
* 根据 proto文件 生成
  ```shell
  goctl rpc protoc *.proto --go_out . --go-grpc_out . --zrpc_out . --style=goZero
  ```
* 生成 proto 拆分的 message
  ```shell
  protoc -I ./xx  --go_out=paths=source_relative:.  --go-grpc_out=paths=source_relative:. orders.rpc.proto
  ```
  
### api
* 创建 api
    ```shell
    goctl api new admin
    ```
* 根据 api 文件生成
    ```shell
    goctl api go -api xx.api --dir .
    ```
* 生成 markdown 文档
    ```shell
    goctl api doc --dir ./
    ```

