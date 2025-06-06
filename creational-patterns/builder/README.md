# Builder生成器模式

Builder模式包含以下几个关键角色：
### 1.产品(Product)
这是最终要创建的复杂对象。它包含多个组成部分，通常具有多个可配置参数。在例子中实际上就是 ServerConf

### 2. 建造者(Builder)
这是一个接口或抽象类，定义了创建产品各个部分的抽象方法。在Go中，通常会使用具体的Builder结构体而非接口。

### 3.具体建造者(Concrete Builder)
实现Builder接口的具体类，提供构建产品各部分的具体实现，并跟踪构建过程。在Go中，通常直接使用具体的Builder结构体，省略接口定义。
在更复杂的场景中，可能会有多个具体建造者，比如：
```cgo
type TestServerConfBuilder struct {
 ServiceConfBuilder
}

type ProductServiceConfBuilder struct {
 ServiceConfBuilder
}
```


### 4. 指挥者(Director)
负责按照特定顺序调用Builder的方法来构建产品。它隐藏了产品构建的复杂性。在Go中，通常使用工厂函数代替Director角色。
```cgo
// 指挥者: 负责初始化并返回Builder
func NewServerConfigBuilder() *ServerConfigBuilder {
    return &ServerConfigBuilder{
        config: &ServerConfig{
            // 设置默认值
            protocol: "http",
            port:     8080,
            maxConn:  100,
            timeout:  30,
        },
    }
}
```

### 5.客户端(Client)
使用Builder模式创建对象的代码。客户端通过Director获取Builder，然后使用Builder创建产品。
```cgo
// 客户端: 使用Builder模式创建对象
func main() {
    // 通过Director获取Builder
    builder := NewServerConfigBuilder()
    
    // 使用Builder配置产品
    config := builder.
        SetHost("localhost").
        SetPort(9000).
        SetProtocol("https").
        SetMaxConn(1000).
        Build()
    
    // 使用最终创建的产品
    fmt.Printf("服务器配置: %+v\n", *config)
}
```
