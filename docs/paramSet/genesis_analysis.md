# Genesis Parameters Settings

> genesis文件中记录的是区块链创世状态，包含：共识参数、初始内置账户、铸币和挖矿、stake、qcp和qsc的初始数据、approve和distribution、governance和guardian。

## chain参数

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|genesis_time|创世时间|"genesis_time": "2019-06-28T06:19:23.196357577Z"|--|执行命令qosd init的时间|
|chain-id|链名称|"chain_id": "capricorn-3000"|"chain_id": "QOS"|--|

## 共识参数

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|max_bytes|块的最大字节，最大为1MB|"max_bytes": "1048576"|？|--|
|max_gas|块的最大限制gas|"max_gas": "-1"|？|-1：表示不受限制|
|time_iota_ms|单位是ms，也就是1s，该值用于在共识对块vote时候，计算出投票时间。|"time_iota_ms": "1000"|1000|共识状态中的锁定块或是提议块的时间加上该值，得出投票时间，如果小于当前时间votetime为计算出的时间，否则votetime为当前时间。|

## 证据参数

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|max_age|证据的有效时长,单位s|"max_age": "100000"|？|双签的记录,超过时间删除双签记录。Tendermint的单位是s,换算时间为27.8h。|

> tendermint代码示例
>
> EvidenceParams determine how we handle evidence of malfeasance.only accept new evidence more recent than this.对于渎职证据的处理方式。只接受比这个更新的证据

```go
//DefaultEvidenceParams Params returns a default EvidenceParams.
func DefaultEvidenceParams() EvidenceParams {
    return EvidenceParams
    {
        MaxAge: 100000// 27.8 hrs at 1block/s
    }
}
```

## 验证人参数

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|validator|记录验证人信息加密类型|"pub_key_types": ["ed25519"]|"pub_key_types": ["ed25519"]|--|

## 账户初始

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|accounts|记录在创世时内置的账户信息,导出时候记录当前导出条件下的网络中所有账户信息|见下示例|?|--|

> 示例:

```json
{
    "base_account": {
     "account_address": "address10ya5d8a5vy5acrqtylk6nukh6a8kza0vj54keg",
     "public_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "DV9hOS5NsEI8dkq96TUP+FHXh9OWxsnFnN6XAFEsLXE="
     },
     "nonce": "3"
    },
    "qos": "999996831",
    "qscs": [
     {
      "coin_name": "ZZU",
      "amount": "9999990000"
     }
    ]
   }
```

## 挖矿(mint)参数

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|mint:params|通胀阶段配置,根据QOS白皮书规则制定.|见下测试网络默认值|见下主网设定值|假设首块时间为2019-11-01,平均出块时间为5秒,则:第一阶段总块数`19992960`,每块通胀值为`1275448`,即每块产生的QOS约为128(1275448/10000).第二阶段总块数`25246080`,每块通胀值为`505028`,每块产生的QOS约为51(505028/10000). 第三阶段总块数`25246080`,每块通胀值为`252514`,每块产生的QOS约为25(252514/10000).第四阶段总块数`25246080`,每块通胀值为`126158`,每块产生的QOS约为13(126158/10000).|
|endtime|阶段结束的时间，一个阶段结束，即进入下一阶段|--|--|--|
|total_amount|该阶段共发行的qos数量|--|--|--|
|applied_amount|标识本阶段已经分发的QOS数量|--|--|--|
|first_block_time|第一个块的生成时间|--|无法配置|--|
|applied_qos_amount|已经发行的OQS的总数:社区池,每个通胀阶段已发行的QOS|--|根据初始化账户计算得出|--|
|待定|全网一共会发行QOS的总量,该值为一个定值.|?|?|--|

> 测试网络默认值:

```json
"mint": {
      "params": {
        "inflation_phrases": [
          {
            "endtime": "2023-01-01T00:00:00Z",
            "total_amount": "2500000000000",
            "applied_amount": "0"
          },
          {
            "endtime": "2027-01-01T00:00:00Z",
            "total_amount": "12750000000000",
            "applied_amount": "0"
          },
          {
            "endtime": "2031-01-01T00:00:00Z",
            "total_amount": "6375000000000",
            "applied_amount": "0"
          },
          {
            "endtime": "2035-01-01T00:00:00Z",
            "total_amount": "3185000000000",
            "applied_amount": "0"
          }
        ]
      },
      "first_block_time": "0",
      "applied_qos_amount": "48970000000000"
    }
```

> 主网设定值:

```json
"mint": {
      "params": {
        "inflation_phrases": [
        {  
           "end_time": "2023-01-01T00:00:00Z",
           "total_amount": "25500000000000",
           "applied_amount": "0"
          },
          {  
           "end_time": "2027-01-01T00:00:00Z",
           "total_amount": "12750000000000",
            "applied_amount": "0"
          },
          {  
          "end_time": "2031-01-01T00:00:00Z",
          "total_amount": "6375000000000",
          "applied_amount": "0"
          },
          {  
            "end_time": "2035-01-01T00:00:00Z",
            "total_amount": "3185000000000",
            "applied_amount": "0"
          }
        ]
      },
      "first_block_time": "0",
      "applied_qos_amount": "48970000000000"
    }
```

## 权益(stake)参数

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|max_validator_cnt|validator最大数目，白皮书指定为21|21|21|--|
|voting_status_len|设定的一个投票高度,用于validator的活跃性检查|1000|?|--|
|voting_status_least|设定的在投票高度内的最小的已投票块数|500|?|--|
|survival_secs|处于inactive状态的validator的生存时间,单位s,在超出该时间后会将该validator状态转化成closed|600|8小时|在inactive状态的validator,不能进行区块验证，不能提交区块，不能获得挖矿收益和交易费用，不能达成代理合约。经过$survival_secs后将自动退出，失去其验证人身份。|
|unbond_return_height|解除委托操作发起后,委托QOS不会立即返还,需要经过unbond_return_height个高度后,才能返还至委托人账户.|100|运营确定|又称为冻结期,参数名称可能会修改.改名为：unbond_frozen_height|
|redelegation_frozen_height||100|1天|又称为转委托冻结期|
|validators|记录导出高度状态为active的validators|--|--|导出某一高度区块链数据形成genesis.json会填充此部分信息|
|val_votes_info|记录导出高度状态为active的validator的votes信息|--|--|同上|
|val_votes_in_window|记录在设定的投票窗口高度内validator的投票块数|--|--|同上|
|delegators_info|记录委托人向validator的委托交易信息|--|--|同上|
|delegator_unbond_info|记录委托人向validator的解绑交易信息|--|--|同上,`但是导出为null,在链上存在解绑操作的交易.`|
|current_validators|记录导出操作时间的validators集合|--|--|同上|

> validators示例:

```json
{
     "owner": "address16cvparc8ek643ghues5xsd9yl0cjvtk63r4ppl",
     "pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "xaj42R1aRYCyS6+J7xmUxEA4K6EnTjVxsk/7t39XTDU="
     },
     "bond_tokens": "5000000000",
     "description": {
      "moniker": "瑞格钱包",
      "logo": "http://easyzone.tokenxy.cn/logo/rgqb.jpeg",
      "website": "",
      "details": ""
     },
     "status": 0,
     "inactive_code": 0,
     "inactive_time": "0001-01-01T00:00:00Z",
     "inactive_height": "0",
     "min_period": "0",
     "bond_height": "287770"
}
```

> val_votes_info示例:

```json
{
     "validator_pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "xaj42R1aRYCyS6+J7xmUxEA4K6EnTjVxsk/7t39XTDU="
     },
     "vote_info": {
      "startHeight": "287773",
      "indexOffset": "390593",
      "missedBlocksCounter": "0"
     }
}
```

> val_votes_in_window示例:

```json
{
     "validator_pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "XJEifEjxC6ik+UTOMka5V+HJVsjlhKE69CbNf6Yspas="
     },
     "index": "512",
     "vote": true
}
```

> delegators_info示例:

```json
{
     "delegator_addr": "address198qdqrd52uuewz4qguwm83dzcud3cvxjeyrq6k",
     "validator_pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "smATFP0NXD2uqbNazkiBDyMUj+GRgR3TerMahQyfiRo="
     },
     "delegate_amount": "5000000000",
     "is_compound": false
}
```

> delegator_unbond_info示例:

```json

```

> current_validators示例:

```json
{
     "owner": "address1fr2smx0n55t60927qny89nlqy6pm9c7x7w673s",
     "pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "XJEifEjxC6ik+UTOMka5V+HJVsjlhKE69CbNf6Yspas="
     },
     "bond_tokens": "52500000000",
     "description": {
      "moniker": "QOS",
      "logo": "https://avatars0.githubusercontent.com/u/41604972?s=200\u0026v=4",
      "website": "https://github.com/QOSGroup",
      "details": ""
     },
     "status": 0,
     "inactive_code": 0,
     "inactive_time": "0001-01-01T00:00:00Z",
     "inactive_height": "0",
     "min_period": "0",
     "bond_height": "0"
}
```

## qcp和qsc参数

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|qcp:ca_root_pub_key|记录qcp root ca的公钥信息|--|--|--|
|qcp:qcps|记录联盟链以及跨链交易信息|见qcps示例|--|导出某一高度区块链数据形成genesis.json会填充此部分信息|
|qsc:ca_root_pub_key|记录qsc root ca的公钥信息|--|--|--|
|qsc:qscs|记录联盟币的信息|见qscs示例|--|导出某一高度区块链数据形成genesis.json会填充此部分信息|

> qcps示例:

```json
    {
     "chain_id": "dawns-3001",
     "sequence_out": "3",
     "sequence_in": "3",
     "pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "cy+cI2m+dB1Xcln9yRqmebN77TwFWlSvRq299k6wDRM="
     },
     "txs": [
      {
       "txstd": {
        "itx": [
         {
          "type": "qbase/txs/qcpresult",
          "value": {
           "result": {
            "Code": 1,
            "Codespace": "sdk",
            "Data": null,
            "Log": "{\"codespace\":\"sdk\",\"code\":1,\"message\":\"TxStd's ITx ValidateData error:  ERROR:\\nCodespace: transfer\\nCode: 201\\nMessage: \\\"QOS、QSCs not equal in Senders and Receivers\\\"\\n\"}",
            "GasWanted": "100000",
            "GasUsed": "4144",
            "FeeAmount": "0",
            "FeeDenom": "",
            "Tags": [
             {
              "key": "cWNwLmZyb20=",
              "value": "Y2Fwcmljb3JuLTMwMDA="
             },
             {
              "key": "cWNwLnRv",
              "value": "ZGF3bnMtMzAwMQ=="
             }
            ]
           },
           "qcporiginalsequence": "3",
           "qcpextends": "heigth:157080,hash:B05FA6EAA8AC1D6E4BC469E507B3D673A599B3A3936D1AC6B77ECA0AE5F415A9",
           "info": ""
          }
         }
        ],
        "sigature": null,
        "chainid": "dawns-3001",
        "maxgas": "0"
       },
       "from": "capricorn-3000",
       "to": "dawns-3001",
       "sequence": "3",
       "sig": {
        "pubkey": null,
        "signature": null,
        "nonce": "0"
       },
       "blockheight": "676558",
       "txindex": "0",
       "isresult": true,
       "extends": ""
      }
    ]
}
```

> qscs示例:

```json
{
     "name": "ZZU",
     "chain_id": "capricorn-3000",
     "extrate": "1",
     "description": "",
     "banker": "address10ya5d8a5vy5acrqtylk6nukh6a8kza0vj54keg"
}
```

## approve参数

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|approves|记录链上发生的授权交易信息的集合|见approves示例|--|导出某一高度区块链数据形成genesis.json会填充此部分信息|

> approves示例:

```json
"approves": [
    {
     "from": "address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa",
     "to": "address10ya5d8a5vy5acrqtylk6nukh6a8kza0vj54keg",
     "qos": "1000",
     "qscs": null
    }
]
```

## 分红(distribution)参数

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|community_fee_pool|记录社区池中的QOS数量|默认为0|0|--|
|last_block_proposer|上一个区块的提议者地址|--|--|在导出某一高度状态时会填充此部分信息|
|pre_distribute_amount|`?`|默认为0|?|在导出某一高度状态时会填充此部分信息,仍然是0|
|validators_history_period|`?`|见下示例|--|在导出某一高度状态时会填充此部分信息|
|validators_current_period|导出高度的validators集合|见下示例|--|同上|
|delegators_earning_info|所有委托人获取的收益信息|见下示例|--|同上|
|delegators_income_height|所有委托人获取收益的区块高度|见下示例|--|同上|
|validator_eco_fee_pools|validator的获取的收益信息|见下示例|--|同上|
|proposer_reward_rate|对于出块的验证人,他获得额外的收益：全网单块挖矿收益 乘 该比例|0.04|0.01|--|
|community_reward_rate|每一个区块中包含的QOS，将有该设定比例的QOS归属于社区基金|0.01|0.02|--|
|validator_commission_rate|由于验证人付出了人力和物力，验证人可以从总收益中抽取一定比例的佣金，QOS网络中的验证人佣金是统一的，以该参数定义。|0.01|?|此参数后期会修改为:可由validator自行设置|
|delegator_income_period_height|创建delegate后，由该参数定义之后的每多少块为一个分配周期，在每个周期交替时为委托人分配收益/处理请求。|100|半小时|--|
|gas_per_unit_cost|每一个gas的单价,单位QOS.|10|100|gas费用=gas数量 乘 gas单价；1GSA=0.000001QOS（1/100/10000）|

> validators_history_period示例

```json
{
     "validator_pubkey": {
      "type": "tendermint/PubKeyEd25519",
      "value": "xaj42R1aRYCyS6+J7xmUxEA4K6EnTjVxsk/7t39XTDU="
     },
     "period": "4044",
     "summary": {
      "value": "0.407568896800000000"
     }
}
```

> validators_current_period示例:

```json
{
     "validator_pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "xaj42R1aRYCyS6+J7xmUxEA4K6EnTjVxsk/7t39XTDU="
     },
     "current_period_summary": {
      "fees": "488521",
      "period": "4047"
     }
}
```

> delegators_earning_info示例:

```json
{
     "validator_pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "xaj42R1aRYCyS6+J7xmUxEA4K6EnTjVxsk/7t39XTDU="
     },
     "delegator_address": "address16cvparc8ek643ghues5xsd9yl0cjvtk63r4ppl",
     "earning_start_info": {
      "previous_period": "4046",
      "bond_token": "5000000000",
      "earns_starting_height": "692370",
      "first_delegate_height": "287770",
      "historical_rewards": "22980",
      "last_income_calHeight": "692370",
      "last_income_calFees": "520960"
     }
}
```

> delegators_income_height示例:

```json
{
     "validator_pub_key": {
      "type": "tendermint/PubKeyEd25519",
      "value": "cE9/spWYL2+D8Gg7ffzJJqTDeovo0UBsI9fgnY7DhI0="
     },
     "delegator_address": "address16wxgq89am0zpyp4nw9y95lwanw7rs2uameqa6h",
     "height": "692485"
}
```

> validator_eco_fee_pools示例:

```json
    {
     "validator_address": "address1qwlsvalla3lk93sk7cd9yfj9r3t6ulrt2m5206",
     "eco_fee_pool": {
      "proposerTotalRewardFee": "161663656",
      "commissionTotalRewardFee": "38044674",
      "preDistributeTotalRewardFee": "3800281959",
      "preDistributeRemainTotalFee": "512446"
     }
    }
```

## 治理(governance)参数

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|starting_proposal_id|提议的编号,下一个提议的index|1|1|--|
|min_deposit|最小的抵押QOS数量,只有达到这个限制,提议才会进入voting阶段|100000000|?|--|
|min_proposer_deposit_rate|提议者抵押QOS数量占min_deposit的比例 > 0.334(最小抵押比例),才能提议成功|0.334|?|--|
|max_deposit_period|提议成功后,可以对该提议进行抵押QOS的时长.单位ns|7d=604800s|?|--|
|voting_period|提议达到最小质押数量,进入voting阶段,该参数表示对该提议进行投票的时长|7d=604800s|?|--|
|quorum|对提议进行vote的power占全网power的最小比例,达不到该比例,提议以不通过处理.|0.334|?|--|
|threshold|对提议进行vote:yes的power占全网的比例 大于 0.5,提议以通过处理.|0.5|?|--|
|veto|对提议进行vote:veto的power占全网power的比例 大于 0.334,提议以不通过处理.|0.334|?|--|
|penalty|对提议不进行投票的validator的惩罚比例|0|?|--|
|burn_rate|提议通过(PASS)或不通过(REJECT)都要进行销毁抵押Deposit * $burn_rate，作为治理费用，剩余的Deposit才会原路返回|0.2|?|--|
|proposals|记录网络中发生的提议集合(提议已经完成voting阶段得到结果)|见下示例|--|--|

> proposals示例:

```json
    {
     "proposal": {
      "proposal_content": {
       "type": "gov/TextProposal",
       "value": {
        "title": "myproposal",
        "description": "the first proposal",
        "deposit": "750000000"
       }
      },
      "proposal_id": "5",
      "proposal_status": "Rejected",
      "final_tally_result": {
       "yes": "2000000000",
       "abstain": "0",
       "no": "100000",
       "no_with_veto": "0"
      },
      "submit_time": "2019-08-01T09:00:38.798681467Z",
      "deposit_end_time": "2019-08-08T09:00:38.798681467Z",
      "total_deposit": "750000000",
      "voting_start_time": "2019-08-01T09:00:38.798681467Z",
      "voting_start_height": "516054",
      "voting_end_time": "2019-08-08T09:00:38.798681467Z"
     },
     "deposits": null,
     "votes": [
      {
       "voter": "address10xwx06gnrt3dlz7hfrx6a8wx3gyeghxm54rv7a",
       "proposal_id": "5",
       "option": "Yes"
      },
      {
       "voter": "address1l0wn66gh45nfta2r4vq8z54wu9hgarss298e9g",
       "proposal_id": "5",
       "option": "No"
      },
      {
       "voter": "address1nnvdqefva89xwppzs46vuskckr7klvzk8r5uaa",
       "proposal_id": "5",
       "option": "Yes"
      }
     ]
    }
```

## 系统账户(guardian)参数

|参数名称|参数释义|测试网默认值|主网设定值|备注|
|--|--|--|--|--|
|description|系统账户的描述|--|--|--|
|guardian_type|1:存在于genesis文件中<br>2:由genesis文件中guardian账户通过tx添加的|--|--|--|
|address|guardian账户的地址|--|--|--|
|creator|创建该guardian账户的地址|--|--|--|

> guardians示例:

```json
{
     "description": "genesis guardian",
     "guardian_type": 1,
     "address": "address14k430znxuf83ruta9qq8hzyrxe8x7mkya4x60a",
     "creator": "address1ah9uz0"
}
```
