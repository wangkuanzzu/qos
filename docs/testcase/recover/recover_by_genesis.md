# recover_by_genesis

> 演习验证：版本升级，经过社区投票后，所有节点导入某一高度的genesis（文件较大），存储文件从零开始
> <https://github.com/QOSGroup/qos/issues/272>

## 数据准备（genesis）

    1. 数据大小分等级测试。例如：2M/1G/更大
    2. 数据导出从0到指定高度。例如：qosd export  --height 100000 --for-zero-height

    ps：由于测试网络中账户、交易等数据较少，导出数据大小不足满足测试情景。所以要创造测试数据，比如批量生成账户，参与交易。

## 测试步骤

    在启动全节点时候，初始化账户信息（满足量级：万）。
    1. 节点init后，使用qosd add-genesis-accounts命令增加10000个初始账户，然后启动。启动命令：qosd start --pruning nothing 用于记录每个高度的状态保存下来。
    2. 保证初始账户的所有qos合计为49亿。

    在节点运行中，停止qosd进程，开始导出从0到指定高度数据。
    1. 停掉节点qosd进程
    2. 导出从height=0开始至某一高度的genesis文件，不加导出路径的话，默认文件路径为：$HOME/.qosd/genesis-100000-****.json
        > qosd export --for-zero-height --height 100000     预测结果：启动成功
        > qosd export --height 100000                       预测结果：启动失败
    3. 执行qosd unsafe-reset-all 完成以下操作：清空区块链数据库，地址表，重置priv_validator.json至创世状态
    4. 将原始的genesis文件首先备份，然后使用导出的genesis文件(注意名称修改为genesis.json)将其替换.
        > 其实原来的genesis文件已经无用了，直接替换也是OK的。
        > genesis文件路径：$HOME/.qosd/config/genesis.json
    5. 执行qosd start
        查看启动情况。

## 测试结果

    1、导出genesis文件中大约包含10000个账户，此时文件大约为2M。
    2、使用两种导出方式得到的genesis文件，替换原来的文件，测试重启：执行qosd start启动都正常。

## 其他

    1、对于导出genesis文件要测试10G大小无法实现。原因json的文本文件对于2G大小的文件都已经很难打开了。
    2、10G大小指的应该是data目录数据，而不是针对genesis文件。
