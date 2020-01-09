# goal for go programmning

加速golang日常开发，使开发体验像python一样流畅

[英文](/README.md)

## 原则

- **轻量**：只依赖标准库，不引入第三方依赖
- **高效**：选择高效方式实现，没有任何额外消耗
- **简洁**：接口简单，不引入复杂实现
- **通用**：实现核心功能，用 20% 的代码满足 80% 的场景

## 使用方式

1. 安装依赖

    ```bash
    go get github.com/sweetycode/goal
    ```

2. 项目中引入依赖

    ```go
    import "github.com/sweetycode/goal"
    ```

3. 尽情享用。。。

    ```go
    intVal := goal.Atoi("123")
    ```

## 开源协议

本开源项目遵守 [MIT 开源协议](https://opensource.org/licenses/MIT).