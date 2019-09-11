# The compare of recover by data and genesis

|dimension|data|genesis|remarks|
|--|--|--|--|
|chain-id|no update：chain-id保存在genesis文件中，无法修改|update：使用导出的genesis文件可修改chain-id名称||
|history tx|query：可以查询历史交易|no query：数据清空，历史交易无法查询，账户状态可查询||
|genesis|origin：无须替换genesis| new :替换genesis文件，文件来源qosd export，文件导出来源节点启动需增加参数--pruning nothing|参数表示无修剪保留数据信息。|
|height|接着数据高度继续增加| 高度从1开始||
|data|首先qosd unsafe-reset-all清空数据，使用共识的data数据进行替换|使用qosd unsafe-reset-all清空数据||
|difficulty|复杂|简单|复杂:替换data目录数据|
|||||
