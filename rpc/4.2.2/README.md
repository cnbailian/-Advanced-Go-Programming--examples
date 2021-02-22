当我在当前目录中使用 `protoc --go-netrpc_out=plugins=netrpc:. hello.proto` 命令时，会提示一条警告：
```
WARNING: Package "github.com/golang/protobuf/protoc-gen-go/generator" is deprecated.
	A future release of golang/protobuf will delete this package,
	which has long been excluded from the compatibility promise.
```
于是我找到了这个 issue：https://github.com/golang/protobuf/issues/1104

这是否意味着不应该依赖于内部包实现 protoc generator？
