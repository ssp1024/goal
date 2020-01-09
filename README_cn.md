# goal for go programmning

## 目标

加速golang日常开发，使开发体验像python一样流畅

- 轻量：只依赖标准库，不引入第三方依赖
- 高效：选择最高效方式实现，不必担心性能问题
- 简洁：接口简单，不引入复杂实现
- 通用：实现核心功能，用 20% 的实现覆盖 80% 的需求

## 使用方式

1. 安装依赖

    ```bash
    go get github.com/sweetycode/goal
    ```

2. go文件中引入依赖

    ```go
    import "github.com/sweetycode/goal"
    ```

3. 尽情享用。。。

    ```go
    intVal := goal.Atoi("123")
    ```

## 开源协议

本开源项目遵守 [MIT 开源协议](https://opensource.org/licenses/MIT).