# Test Cases

- [缺失必须参数](./TestCase01.md)
- [参数不合法](./TestCase02.md)
- [正常抵押提议](./TestCase03.md)
- [业务情景](./TestCase04.md)

## Description
>     deposit.
>
>     向提案抵押存款。

## Usage

```
  qoscli tx deposit [flags]
```

## Available Commands

>无可用命令

## Flags

| ShortCut | Flag                | Input Type | Default Input             | Input Range | Description                                                |
|:---------|:--------------------|:-----------|:--------------------------|:------------|:-----------------------------------------------------------|
| `-h`     | `--help`            | -          | -                         | -           | (可选)帮助文档                                                   |

| ShortCut | Flag                | Required | Input Type | Default Input             | Input Range | Description                                            |
|:---------|:--------------------|:---------|:-----------|:--------------------------|:------------|:-------------------------------------------------------|
| `-h`     | `--help`            | ✖        | -          | -                         | -           | 帮助文档                                                   |
| -        | `--async`           | ✖        | -          | -                         | -           | 是否异步广播交易                                               |
| -        | `--chain-id`        | ✖        | string     | -                         | -           | Tendermint节点的链ID                                       |
| -        | `--indent`          | ✖        | -          | -                         | -           | 向JSON响应添加缩进                                            |
| -        | `--max-gas`         | ✖        | int        | `100000`                  | -           | 每个Tx设置的气体限制值                                           |
| `-n`     | `--node`            | ✖        | string     | `"tcp://localhost:26657"` | -           | 为此链提供的Tendermint RPC接口: `<host>:<port>`                |
| -        | `--nonce`           | ✖        | int        | -                         | -           | 要签署Tx的帐户nonce                                          |
| -        | `--nonce-node`      | ✖        | string     | -                         | -           | 用于其他链查询账户nonce的Tendermint RPC接口: `tcp://<host>:<port>` |
| -        | `--qcp`             | ✖        | -          | -                         | -           | 是否启用QCP模式(qcp mode), 发送QCP Tx                          |
| -        | `--qcp-blockheight` | ✖        | int        | -                         | -           | QCP模式Flag标志: 原始Tx块高度，块高度必须大于0                          |
| -        | `--qcp-extends`     | ✖        | string     | -                         | -           | QCP模式Flag标志: QCP Tx扩展信息                                |
| -        | `--qcp-from`        | ✖        | string     | -                         | -           | QCP模式Flag标志: QCP Tx源链ID                                |
| -        | `--qcp-seq`         | ✖        | int        | -                         | -           | QCP模式Flag标志: QCP顺序                                     |
| -        | `--qcp-signer`      | ✖        | string     | -                         | -           | QCP模式Flag标志: QCP Tx签名者key名称                            |
| -        | `--qcp-txindex`     | ✖        | int        | -                         | -           | QCP模式Flag标志: 原始Tx索引                                    |
| -        | `--trust-node`      | ✖        | -          | -                         | -           | 是否信任连接的完整节点（不验证其响应证据）                                  |
| -        | `--amount`          | ✖        | uint       | -                         | -           | (主要参数)要抵押存款的QOS百分比                                           |
| -        | `--depositor`       | ✖        | uint       | -                         | -           | (主要参数)抵押人。                                                   |
| -        | `--proposal-id`     | ✖        | uint       | -                         | -           | (主要参数)提案ID                                                   |

## Global Flags

| ShortCut | Flag         | Required | Input Type | Default Input | Input Range       | Description  |
|:---------|:-------------|:---------|:-----------|:--------------|:------------------|:-------------|
| `-e`     | `--encoding` | ✖        | string     | `hex`         | `hex`/`b64`/`btc` | 二进制编码        |
| -        | `--home`     | ✖        | string     | `/.qoscli`    | -                 | 配置和数据的目录     |
| `-o`     | `--output`   | ✖        | string     | `text`        | `text`/`json`     | 输出格式         |
| -        | `--trace`    | ✖        | -          | -             | -                 | 打印出错时的完整堆栈跟踪 |