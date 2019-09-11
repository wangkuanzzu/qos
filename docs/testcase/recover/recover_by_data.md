# recover_by_data

> 演习验证：版本升级，经过社区投票后，所有节点导入某一高度的genesis（文件较大），存储文件从零开始
> <https://github.com/QOSGroup/qos/issues/272>

## 数据准备（data）

data目录说明：

    ```bash
        drwxr-xr-x 2 vagrant vagrant  97 Sep  2 17:08 application.db
        drwxr-xr-x 2 vagrant vagrant  80 Sep  2 15:24 blockstore.db
        drwx------ 2 vagrant vagrant  16 Sep  2 15:24 cs.wal
        drwxr-xr-x 2 vagrant vagrant  80 Sep  2 15:24 evidence.db
        -rw------- 1 vagrant vagrant 399 Sep  2 17:27 priv_validator_state.json
        drwxr-xr-x 2 vagrant vagrant  80 Sep  2 15:24 state.db
        drwxr-xr-x 2 vagrant vagrant  80 Sep  2 15:24 tx_index.db
    ```

|FileOrDir|Explain|Remarks|
|--|--|--|
|application.db|记录应用程序的数据||
|blockstore.db|记录每一个区块的信息||
|cs.wal|日志记录||
|evidence.db|存储证据信息||
|priv_validator_state.json|记录验证人签名状态信息|只有是验证人内容才会更新|
|state.db|记录存储最新的世界状态||
|tx_index.db|存储交易索引信息||

    1. data数据大小分等级测试。例如：1G/10G

    ps：由于测试网络中账户、交易等数据较少，导出数据大小不足满足测试情景。所以要创造测试数据，比如批量生成账户，批量生成交易，以达到需要的数据大小等级。

## 测试步骤

    准备任务：在QOS网络中启动多个节点（以3个为例：A,B,C），正常运行。
    恢复场景：系统需要升级（bug修复等），数据没有问题。
    开始测试：
        1.选择其中一个节点A。停止qosd进程，停止打块，高度不再增加。
        2.将该节点A目录：~/.qosd/data/* 备份压缩为文件(除priv_validator_state.json外)：qos_data_bak.tar
        3.只有一个节点（votingpower小于全网1/3）退出共识，不会影响qos全网后续打块，所以节点B和C的数据会比A备份时候的大，高度也会大。之后余下两个节点收到升级通知，同样停止qosd进程，全网停止运行，等待重新启动。
        4.等待节点A系统升级完成并启动，首先追赶至节点B和C的高度，然后开始打新块。
        5.将节点A备份的数据qos_data_bak.tar交给节点B和C。
        6.节点B和C完成以下操作：执行qosd unsafe-reset-all ,将收到的备份文件放置到目录：.qosd/data/ 
        7.最后执行qosd start

## 测试结果

    1、数据拷贝后，在启动qosd，可以看到从停止那一刻高度开始继续打块共识。
    2、需要超过全网votingpower的2/3节点启动后，开始打块，不够的情况下无法打块。

## 获取指定高度以下的data数据

> 一、区块信息，交易信息等，不仅仅是指定高度当时的状态

    步骤如下：
    1. 指定一个信任的全节点，从该节点获取指定高度的data数据。假设高度为H
    2. 重新启动一个全节点，同步信任的全节点数据，使用的是修改过源代码后编译的qosd程序（如何修改请看第二部分）。
    3. 开始同步，由于使用的修改后的qosd程序，在同步到我们指定的高度H后，便会停止，此时我们便获取到想要的data数据。

> 二、源代码修改

    1. 使用目前主网使用的qos版本代码进行修改。
    2. 找到文件：\github.com\QOSGroup\qos\app\app.go
    3. 修改方法：BeginBlocker
    4. 在return之前增加代码：if ctx.BlockHeight() > H { panic(ctx) }
    5. 重新编译qosd程序。

> 三、代码修改如下

```go
    func (app *QOSApp) BeginBlocker(ctx context.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {

        //测试同步至指定高度100，获取指定高度之下的data数据
        if ctx.BlockHeight() > 100 {
            panic(ctx)
        }

        return app.mm.BeginBlock(ctx, req)
    }
```
