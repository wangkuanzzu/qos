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
