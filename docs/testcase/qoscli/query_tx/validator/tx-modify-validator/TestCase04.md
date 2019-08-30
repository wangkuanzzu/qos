# test case of qoscli modify validator*

> `qoscli modify validator*` 激活验证节点

---

## 情景说明

1. 编辑验证节点：在编辑节点之前查询节点的信息，修改之后再次进行查询，确认信息成功修改。
2. 虽然只是修改验证节点的其他信息，但是需要验证修改前后的收益是否有变化。
3. 编辑验证节点的tx操作，有基础的GAS费用消耗（0.18QOS=180000GAS），再加上访问存储所消耗的GAS费用。验证操作人账户余额的扣除的GAS费用。

## 测试命令

```bash
    //执行modify语句前 查询验证节点信息
    qoscli query validator jlgy01 --indent

    qoscli tx modify-validator --moniker jlgy666 --owner jlgy01 --logo "http://pic32.nipic.com/20130813/3347542_160503703000_2.jpg" --website "https://github.com/test" --details "jlgy23333333333"

    //执行modify语句后 查询验证节点信息
    qoscli query validator jlgy01 --indent
```

## 预测结果

```bash
1.两次查询的节点信息有变化。
```

## 测试结果

```bash
    qoscli query validator jlgy01 --indent
    {
    "owner": "address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa",
    "validatorAddress": "6E713D1F3CCE28D820C39059E5D08D21D646FA8E",
    "validatorPubkey": {
        "type": "tendermint/PubKeyEd25519",
        "value": "exGS/yWJthwY8za4dlrPRid2I9KE4G15nlJwO/+Off8="
    },
    "bondTokens": "2000000000",
    "description": {
        "moniker": "jlgy",
        "logo": "",
        "website": "",
        "details": ""
    },
    "status": "active",
    "InactiveDesc": "",
    "inactiveTime": "0001-01-01T00:00:00Z",
    "inactiveHeight": "0",
    "bondHeight": "617422"
    }

    qoscli tx modify-validator --moniker jlgy666 --owner jlgy01 --logo "http://pic32.nipic.com/20130813/3347542_160503703000_2.jpg" --website "https://github.com/wangkuanzzu" --details "jlgy23333333333"
    Password to sign with 'jlgy01':
    {"check_tx":{"gasWanted":"100000","gasUsed":"6703"},"deliver_tx":{"gasWanted":"100000","gasUsed":"17160","tags":[{"key":"YWN0aW9u","value":"bW9kaWZ5LXZhbGlkYXRvcg=="},{"key":"b3duZXI=","value":"YWRkcmVzczFubnZkcWVmdmE4OXh3cHB6czQ2dnVza2NrcjdrbHZ6azhyNXVhYQ=="},{"key":"ZGVsZWdhdG9y","value":"YWRkcmVzczFubnZkcWVmdmE4OXh3cHB6czQ2dnVza2NrcjdrbHZ6azhyNXVhYQ=="}]},"hash":"241AC66206A955AA44AE7A2555EAA9D17320241A700D4749AF74055EEC064C57","height":"617704"}

    qoscli query validator jlgy01 --indent
    {
    "owner": "address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa",
    "validatorAddress": "6E713D1F3CCE28D820C39059E5D08D21D646FA8E",
    "validatorPubkey": {
        "type": "tendermint/PubKeyEd25519",
        "value": "exGS/yWJthwY8za4dlrPRid2I9KE4G15nlJwO/+Off8="
    },
    "bondTokens": "2000000000",
    "description": {
        "moniker": "jlgy666",
        "logo": "http://pic32.nipic.com/20130813/3347542_160503703000_2.jpg",
        "website": "https://github.com/wangkuanzzu",
        "details": "jlgy23333333333"
    },
    "status": "active",
    "InactiveDesc": "",
    "inactiveTime": "0001-01-01T00:00:00Z",
    "inactiveHeight": "0",
    "bondHeight": "617422"
    }

[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc0
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1n64h6prxz6ld5vl8d0rrzsny25nq7vgcnrhu98","public_key":{"type":"tendermint/PubKeyEd25519","value":"m4lqygnU2mG19Fpf3vj2K618G2e2WMwtxu6GANARIVY="},"nonce":"5"},"qos":"5001096596328","qscs":null}}

[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx modify-validator --moniker secondvalidator --owner acc0 --logo "http://pic32.nipic.com/20130813/3347542_160503703000_2.jpg" --website "https://github.com/" --details "update details"
Password to sign with 'acc0':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"26360","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"39100","events":[{"type":"modify-validator","attributes":[{"key":"b3duZXI=","value":"YWRkcmVzczFuNjRoNnByeHo2bGQ1dmw4ZDBycnpzbnkyNW5xN3ZnY25yaHU5OA=="},{"key":"ZGVsZWdhdG9y","value":"YWRkcmVzczFuNjRoNnByeHo2bGQ1dmw4ZDBycnpzbnkyNW5xN3ZnY25yaHU5OA=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"c3Rha2U="},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFuNjRoNnByeHo2bGQ1dmw4ZDBycnpzbnkyNW5xN3ZnY25yaHU5OA=="}]}]},"hash":"7C107679F07B09178C77EE3921186BEE14D2A5D22D27740EA334A5A7549197E0","height":"886"}

//高度为886
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc0
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1n64h6prxz6ld5vl8d0rrzsny25nq7vgcnrhu98","public_key":{"type":"tendermint/PubKeyEd25519","value":"m4lqygnU2mG19Fpf3vj2K618G2e2WMwtxu6GANARIVY="},"nonce":"6"},"qos":"5001096595937","qscs":null}}

//在高度为1080之后，预估acc0的qos=5001496595937 
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc0
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1n64h6prxz6ld5vl8d0rrzsny25nq7vgcnrhu98","public_key":{"type":"tendermint/PubKeyEd25519","value":"m4lqygnU2mG19Fpf3vj2K618G2e2WMwtxu6GANARIVY="},"nonce":"6"},"qos":"5001440714560","qscs":null}}

[vagrant@vagrant-192-168-1-200 ~]$ qoscli query inflation-phrases --indent |grep 'applied_amount'
    "applied_amount": "1471092724"
    "applied_amount": "0"
    "applied_amount": "0"
    "applied_amount": "0"
    "applied_amount": "0"
    "applied_amount": "0"
    "applied_amount": "0"

```

ps：
    //执行编辑validator的交易消耗的5001096596328-5001096595937=391
    此命令消耗的gas费用39100gas明显低于基础gas（0.18QOS=180000gas）消耗。
