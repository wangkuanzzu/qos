# Test Cases

- [缺失必须参数](./TestCase01.md)
- [参数不合法](./TestCase02.md)
- [正常查询委托收益](./TestCase03.md)
- [业务情景]暂无

## Description

>     Query distribution delegator income info.
>
>     查询分发(distribution)委托人收入信息。

## Usage

```
  qoscli query delegator-income [flags]
```

## Available Commands

>无可用命令

## Flags

| ShortCut | Flag           | Required | Input Type | Default Input             | Input Range | Description                             |
|:---------|:---------------|:---------|:-----------|:--------------------------|:------------|:----------------------------------------|
| `-h`     | `--help`       | ✖        | -          | -                         | -           | 帮助文档                                    |
| -        | `--chain-id`   | ✖        | string     | -                         | -           | Tendermint节点的链ID                        |
| -        | `--height`     | ✖        | int        | -                         | -           | (可选)要查询的块高度，省略以获取最新的可证明块                |
| -        | `--indent`     | ✖        | -          | -                         | -           | 向JSON响应添加缩进                             |
| `-n`     | `--node`       | ✖        | string     | `"tcp://localhost:26657"` | -           | 为此链提供的Tendermint RPC接口: `<host>:<port>` |
| -        | `--trust-node` | ✖        | -          | -                         | -           | 是否信任连接的完整节点（不验证其响应证据）                   |
| -        | `--delegator`  | ✖        | string     | -                         | -           | (主要参数)委托人账户本地密钥库名字或账户地址                   |
| -        | `--owner`      | ✖        | string     | -                         | -           | (主要参数)验证人节点的`owner`账户本地密钥库名字或账户地址         |

## Global Flags

| ShortCut | Flag         | Required | Input Type | Default Input | Input Range       | Description  |
|:---------|:-------------|:---------|:-----------|:--------------|:------------------|:-------------|
| `-e`     | `--encoding` | ✖        | string     | `hex`         | `hex`/`b64`/`btc` | 二进制编码        |
| -        | `--home`     | ✖        | string     | `/.qoscli`    | -                 | 配置和数据的目录     |
| `-o`     | `--output`   | ✖        | string     | `text`        | `text`/`json`     | 输出格式         |
| -        | `--trace`    | ✖        | -          | -             | -                 | 打印出错时的完整堆栈跟踪 |
