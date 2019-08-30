# test case of qoscli tx cancel-approve

> `qoscli tx cancel-approve` 取消授权

---

## 情景说明

1. 取消授权，通过查询在取消之前的授权信息和取消之后的授权信息，进行校验，确认是否取消授权成功。
2. 取消授权后，被授权的账户无法使用授权人的账户余额进行交易操作。
3. 取消授权的tx操作，无基础的gas费用消耗，但是有访问存储消耗的gas。

## 测试命令

```bash
    qoscli query approve --from acc1 --to acc2
    qoscli query account acc1

    qoscli tx cancel-approve --from acc1 --to acc2

    qoscli query approve --from acc1 --to acc2
    qoscli query account acc1

    qoscli tx use-approve --coins 500000qos --from acc1 --to acc2
```

## 预测结果

```bash
1.取消后，授权信息查询不到。
2.执行取消交易，账户acc1扣除gas费用。
```

## 测试结果

```bash
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query approve --from acc1 --to acc2
{"from":"address1sr5g3czlpqv8tst82gsczdmws58eerm3qymy84","to":"address18axmtq0t32r3hwxhk2e7g80hpchgqtlrehqkl6","qos":"6000000","qscs":null}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc1
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1sr5g3czlpqv8tst82gsczdmws58eerm3qymy84","public_key":{"type":"tendermint/PubKeyEd25519","value":"F7dzWBxFHyoL4VY9SQ9Bzz09w/ZmIMYvSbub9PcNfpU="},"nonce":"2"},"qos":"3999999999758","qscs":null}}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx cancel-approve --from acc1 --to acc2
Password to sign with 'acc1':
{"check_tx":{"gasWanted":"9223372036854775807","gasUsed":"6931","events":[]},"deliver_tx":{"gasWanted":"9223372036854775807","gasUsed":"7900","events":[{"type":"cancel-approve","attributes":[{"key":"YXBwcm92ZS1mcm9t","value":"YWRkcmVzczFzcjVnM2N6bHBxdjh0c3Q4MmdzY3pkbXdzNThlZXJtM3F5bXk4NA=="},{"key":"YXBwcm92ZS10bw==","value":"YWRkcmVzczE4YXhtdHEwdDMycjNod3hoazJlN2c4MGhwY2hncXRscmVocWtsNg=="}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"YXBwcm92ZQ=="},{"key":"Z2FzLnBheWVy","value":"YWRkcmVzczFzcjVnM2N6bHBxdjh0c3Q4MmdzY3pkbXdzNThlZXJtM3F5bXk4NA=="}]}]},"hash":"43FA43301C5A9D2BF3E1341EB25FB7B63BD8EB96951F59AC83F47C7F4D9DF373","height":"3709"}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query approve --from acc1 --to acc2
{"from":"address1sr5g3czlpqv8tst82gsczdmws58eerm3qymy84","to":"address18axmtq0t32r3hwxhk2e7g80hpchgqtlrehqkl6","qos":"6000000","qscs":null}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query approve --from acc1 --to acc2
ERROR: approve does not exist
[vagrant@vagrant-192-168-1-200 ~]$ qoscli query account acc1
{"type":"qos/types/QOSAccount","value":{"base_account":{"account_address":"address1sr5g3czlpqv8tst82gsczdmws58eerm3qymy84","public_key":{"type":"tendermint/PubKeyEd25519","value":"F7dzWBxFHyoL4VY9SQ9Bzz09w/ZmIMYvSbub9PcNfpU="},"nonce":"3"},"qos":"3999999999679","qscs":null}}
[vagrant@vagrant-192-168-1-200 ~]$ qoscli tx use-approve --coins 500000qos --from acc1 --to acc2
Password to sign with 'acc2':
{"check_tx":{"code":1,"log":"{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: approve\\nCode: 104\\nMessage: \\\"approve not exists\\\"\\n\"}","gasWanted":"9223372036854775807","gasUsed":"1000","events":[]},"deliver_tx":{},"hash":"6FB78182DC3964F31EDD25B2576682006FA0D1A0EDB7293228D340F3A98E5D9E","height":"0"}
ERROR: {"codespace":"sdk","code":1,"message":"TxStd's ITx ValidateData error:  ERROR:\nCodespace: approve\nCode: 104\nMessage: \"approve not exists\"\n"}

```
