# go-seed

## What's it

This is seed project of Go, if you are interested in it, u can contact me by QQ

QQ： MTQzNzg3NjA3Mw==

## Usage



### Web

#### 异常处理



#### 解决跨域问题

在开发时，经常会遇到跨域的问题。很简单， **只需要添加一个中间件**就可以解决跨域了：

```go
func main() {
    app = gin.Default()
    // 使用解决跨域的中间件
    app.Use(goseed.CorsHandler())
    // ... 其他注册逻辑
}
```

