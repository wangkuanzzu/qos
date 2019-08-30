# test case of qoscli tx issue-qsc`

> `qoscli tx issue-qsc` 增发联盟币

---

## 情景说明

1. 初始化联盟币后，进行发币操作，验证是否发币成功。
2. 发币成功后，进行币数量增发，验证币的总量变化是否正确。
3. 发币的tx操作，有基础的gas费用消耗（0.18QOS=180000GAS），再加上访问存储所消耗的GAS费用。

## 测试命令

```bash

```

## 预测结果

```text

```

## 测试结果

```bash
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx issue-qsc --banker zzubanker --qsc-name ZZU --amount 100000000 --indent
Password to sign with 'zzubanker':
{
  "check_tx": {
    "gasWanted": "9223372036854775807",
    "gasUsed": "24397",
    "events": []
  },
  "deliver_tx": {
    "code": 1,
    "log": "{\"codespace\":\"sdk\",\"code\":1,\"message\":\"address1hsurfl3kt334qwj45y3d2hsdwt6ufsfpzphuze no enough coins to pay the gas after this tx done\"}",
    "gasWanted": "9223372036854775807",
    "gasUsed": "35100",
    "events": [],
    "codespace": "sdk"
  },
  "hash": "9CC5C8BE570FC581013082810518C583F0183A41658E7421514B37CF04AB5464",
  "height": "35633"
}
ERROR: {"codespace":"sdk","code":1,"message":"address1hsurfl3kt334qwj45y3d2hsdwt6ufsfpzphuze no enough coins to pay the gas after this tx done"}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx issue-qsc --banker zzubanker --qsc-name ZZU --amount 100000000 --indent
Password to sign with 'zzubanker':
null
ERROR: broadcast_tx_commit: response error: RPC error -32603 - Internal error: Timed out waiting for tx to be included in a block
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx issue-qsc --banker zzubanker --qsc-name ZZU --amount 100000000 --indent
Password to sign with 'zzubanker':
null
ERROR: broadcast_tx_commit: response error: RPC error -32603 - Internal error: Error on broadcastTxCommit: Tx already exists in cache
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query account zzubanker
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1hsurfl3kt334qwj45y3d2hsdwt6ufsfpzphuze","public_key":{"type":"tendermint/PubKeyEd25519","value":"d4ZELB/NC2hnvJBhC+Z8AOwpEOvBxqGO+60hfjCSzOg="},"nonce":"1"},"qos":"1000000000","qscs":null}}
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx issue-qsc --banker zzubanker --qsc-name ZZU --amount 1000000000 --indent
Password to sign with 'zzubanker':
{
  "check_tx": {
    "code": 1,
    "log": "{\"codespace\":\"sdk\",\"code\":1,\"message\":\"invalid nonce. expect: 3, got: 2\"}",
    "gasWanted": "9223372036854775807",
    "gasUsed": "20387",
    "events": []
  },
  "deliver_tx": {},
  "hash": "DC5E86D941D269535F8A2A1986B0520D04D159D174D35389037C88D69A60B847",
  "height": "0"
}
ERROR: {"codespace":"sdk","code":1,"message":"invalid nonce. expect: 3, got: 2"}

```

ps：
    1、执行发币操作之前，banker账户没有qos，交易会进入缓存中，下次执行该命令时候报交易已经存在缓存。
