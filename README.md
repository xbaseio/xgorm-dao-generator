# xgorm-dao-generator

基于 Go struct 模型生成 GORM DAO。生成器会创建两层代码：

- `internal`：字段、查询和 CRUD 实现。
- 外层 DAO：数据库实例、表名和业务扩展入口。

## 安装

```bash
go install github.com/xbaseio/xgorm-dao-generator@latest
```

## 基本用法

```bash
xgorm-dao-generator \
  -model-dir=. \
  -model-names=User,Wallet \
  -dao-dir=../dao/ \
  -sub-pkg-enable=true \
  -mysql-pkg-path=/internal/component/mysql/mysql-default
```

常用参数：

| 参数 | 默认值 | 说明 |
| --- | --- | --- |
| `-model-dir` | 空 | 模型所在目录，必填 |
| `-model-names` | 空 | 要生成的模型名，多个名称用逗号分隔，必填；可用 `Model:table_name` 覆盖表名 |
| `-dao-dir` | 空 | DAO 输出目录，必填 |
| `-model-pkg-path` | 自动推导 | 模型包导入路径 |
| `-model-pkg-alias` | 自动推导 | 模型包导入别名，例如 `modelpkg` |
| `-dao-pkg-path` | 自动推导 | DAO 包导入路径 |
| `-sub-pkg-enable` | `false` | 每个模型使用独立 DAO 子包 |
| `-sub-pkg-style` | `kebab` | 子包目录命名风格：`kebab`、`underscore`、`lower`、`camel`、`pascal` |
| `-file-style` | `underscore` | 生成文件命名风格：`kebab`、`underscore`、`lower`、`camel`、`pascal` |
| `-mysql-pkg-path` | 空 | 生成外层 DAO 的数据库组件路径 |
| `-instance-func` | `Instance` | 外层 DAO 使用的数据库实例函数：`Instance`、`InstanceWrite`、`InstanceRead` |
| `-include-anonymous` | `false` | 是否展开匿名嵌入 struct 的字段 |

`-model-dir`、`-dao-dir` 支持相对路径，均相对于命令执行目录解析。包路径参数只影响生成代码中的 import，不会创建或迁移数据库。

## 匿名嵌入字段

默认情况下，匿名嵌入字段会被跳过。例如：

```go
type User struct {
    ID int
    ModelTime
}
```

默认只生成 `ID`。如果模型确实需要把 `ModelTime` 的字段也生成到 DAO Columns，显式开启：

```bash
xgorm-dao-generator -model-dir=. -model-names=User -dao-dir=../dao/ -include-anonymous=true
```

admin 项目通常将匿名字段作为 ORM 行为或公共模型复用，因此建议保持默认关闭。

## go:generate

```go
//go:generate xgorm-dao-generator -model-dir=. -model-names=User -dao-dir=../dao/ -sub-pkg-enable=true
```

执行：

```bash
go generate ./...
```

生成的 `internal` 文件是机器生成代码，不要手工修改；业务扩展应写在外层 DAO 文件中。
