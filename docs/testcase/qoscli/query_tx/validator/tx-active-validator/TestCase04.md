# test case of qoscli active validator*

> `qoscli active validator*` 激活验证节点

---

## 情景说明

1. 激活验证节点：首先保证节点处理inactive状态，在激活之前进行节点状态查询，之后执行操作，再次进行查询节点状态。
2. 对于已经激活的节点，根据其voting-power查看是否可以成为validator，资格不够可以追加委托绑定的tokens，验证是否可提高排名。（操作的验证节点最好是最后一名）
3. 在激活节点后如果有能力成为validator，等待一定收益周期，确认是否有得到挖矿奖励。验证委托人账户数额是否在发放周期后增加。
4. 激活验证节点的tx操作，没有基础的GAS费用消耗，但是有访问存储所消耗的GAS费用。验证操作人账户余额的扣除的GAS费用。

## 测试命令

```bash

```

## 预测结果

```bash
//全网7个验证节点 第七名绑定tokens 600000000  在执行完active后增加绑定100000000后达到200000000 状态依旧是inactive
//使用账户jlgy01增加对其的delegate，增加410000000，此版本不支持对inactive的validator进行增加委托
//使用jlgy账户在执行active时候增加委托tokens数量410000000
//等待收益发放周期720高度后，委托人账户得到挖矿奖励。
```

## 测试结果

```bash
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query validator jlgy
{"owner":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","validatorAddress":"1C97C563A804CB8B379622DA2610422807DFDCA9","validatorPubkey":{"type":"tendermint/PubKeyEd25519","value":"Hx/cQ562Bp8ZJpRBg0/r17+QvgEPxBptqjBoAeNYH2c="},"bondTokens":"100000000","description":{"moniker":"jlgy","logo":"","website":"","details":""},"commission":{"commission_rates":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.010000000000000000"},"update_time":"2019-08-29T06:16:18.033191341Z"},"status":"inactive","InactiveDesc":"Replaced","inactiveTime":"2019-08-29T06:16:18.033191341Z","inactiveHeight":"16348","bondHeight":"16348"}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query account jlgy
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","public_key":{"type":"tendermint/PubKeyEd25519","value":"9fpQTirTSpYNHw4qu6CdN456oEZeb5AM0PxIMRZre2s="},"nonce":"2"},"qos":"1799995425","qscs":null}}
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx active-validator --owner jlgy --tokens 100000000
Password to sign with 'jlgy':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"9498","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"52800","events":[{"type":"active-validator","attributes":[{"key":"dmFsaWRhdG9y","value":"YWRkcmVzczFyanR1MmNhZ3FuOWNrZHVreXRkenZ5eno5cXJhbGg5ZnAzM2NrOQ=="},{"key":"b3duZXI=","value":"YWRkcmVzczFucm0yNHZ4YXYwdTU4anB0MDR2N3Z0YXZyamp4d3c2MHg5eW03cA=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"c3Rha2U="},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFucm0yNHZ4YXYwdTU4anB0MDR2N3Z0YXZyamp4d3c2MHg5eW03cA=="}]}]},"hash":"0A066CF3C724F4ECE6F40200774A1811FA4169979CFBF4DDF24B12B93BD8F8EB","height":"16404"}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query account jlgy
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","public_key":{"type":"tendermint/PubKeyEd25519","value":"9fpQTirTSpYNHw4qu6CdN456oEZeb5AM0PxIMRZre2s="},"nonce":"3"},"qos":"1699994897","qscs":null}}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query validator jlgy
{"owner":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","validatorAddress":"1C97C563A804CB8B379622DA2610422807DFDCA9","validatorPubkey":{"type":"tendermint/PubKeyEd25519","value":"Hx/cQ562Bp8ZJpRBg0/r17+QvgEPxBptqjBoAeNYH2c="},"bondTokens":"200000000","description":{"moniker":"jlgy","logo":"","website":"","details":""},"commission":{"commission_rates":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.010000000000000000"},"update_time":"2019-08-29T06:16:18.033191341Z"},"status":"inactive","InactiveDesc":"Replaced","inactiveTime":"2019-08-29T06:21:06.822679715Z","inactiveHeight":"16404","bondHeight":"16348"}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx delegate --owner jlgy --delegator jlgy01 --tokens 410000000
Password to sign with 'jlgy01':
{"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  validator status not match. except: 0, actual:1\"}","gasWanted":"9223372036854775807","gasUsed":"2588","events":[]},"deliver_tx":{},"hash":"491B9C9606B2AA0BCA4B6EFCE2615A04D62841420B6601472B0F3F451C32F050","height":"0"}
ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  validator status not match. except: 0, actual:1"}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli tx active-validator --owner jlgy --tokens 410000000
Password to sign with 'jlgy':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"9504","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"53000","events":[{"type":"active-validator","attributes":[{"key":"dmFsaWRhdG9y","value":"YWRkcmVzczFyanR1MmNhZ3FuOWNrZHVreXRkenZ5eno5cXJhbGg5ZnAzM2NrOQ=="},{"key":"b3duZXI=","value":"YWRkcmVzczFucm0yNHZ4YXYwdTU4anB0MDR2N3Z0YXZyamp4d3c2MHg5eW03cA=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"c3Rha2U="},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFucm0yNHZ4YXYwdTU4anB0MDR2N3Z0YXZyamp4d3c2MHg5eW03cA=="}]}]},"hash":"34A0ED5C8ABCADCD7688CD7427635D33F1CF663EA5E21F72BD1FB342C1C2F6B1","height":"16539"}

[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query validator jlgy
{"owner":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","validatorAddress":"1C97C563A804CB8B379622DA2610422807DFDCA9","validatorPubkey":{"type":"tendermint/PubKeyEd25519","value":"Hx/cQ562Bp8ZJpRBg0/r17+QvgEPxBptqjBoAeNYH2c="},"bondTokens":"610000000","description":{"moniker":"jlgy","logo":"","website":"","details":""},"commission":{"commission_rates":{"rate":"0.100000000000000000","max_rate":"0.200000000000000000","max_change_rate":"0.010000000000000000"},"update_time":"2019-08-29T06:16:18.033191341Z"},"status":"active","InactiveDesc":"Replaced","inactiveTime":"2019-08-29T06:21:06.822679715Z","inactiveHeight":"16404","bondHeight":"16348"}

//高度为16594查询  ，成为验证人的高度为16539  下次收益发放的高度16539+720=17259‬
[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query account jlgy
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","public_key":{"type":"tendermint/PubKeyEd25519","value":"9fpQTirTSpYNHw4qu6CdN456oEZeb5AM0PxIMRZre2s="},"nonce":"4"},"qos":"1289994367","qscs":null}}

//高度大于17259‬时查询
‬[root@iz2zef57ni8z6ydrcloxy0z ~]# qoscli query account jlgy
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1nrm24vxav0u58jpt04v7vtavrjjxww60x9ym7p","public_key":{"type":"tendermint/PubKeyEd25519","value":"9fpQTirTSpYNHw4qu6CdN456oEZeb5AM0PxIMRZre2s="},"nonce":"4"},"qos":"1395327397","qscs":null}}

```
