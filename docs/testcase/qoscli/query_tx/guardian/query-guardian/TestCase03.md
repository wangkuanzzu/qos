# Description

```text
  特权账户查询，查询单个guardian，可以通过指定keys_name 或是 account_address。
```

## Input

```bash
    //本地密钥库中存在账户abc ，可以使用key_name查询
    qoscli query guardian abc

    //本地密钥库中不存在账户abc，只能使用account_address查询
    qoscli query guardian address10xwx06gnrt3dlz7hfrx6a8wx3gyeghxm54rv7a
```

## Output

```bash
    qoscli query guardian abc
    {"description":"this is the abc guardian","guardian_type":1,"address":"address10xwx06gnrt3dlz7hfrx6a8wx3gyeghxm54rv7a","creator":"address1ah9uz0"}

    qoscli query guardian address10xwx06gnrt3dlz7hfrx6a8wx3gyeghxm54rv7a
    {"description":"this is the abc guardian","guardian_type":1,"address":"address10xwx06gnrt3dlz7hfrx6a8wx3gyeghxm54rv7a","creator":"address1ah9uz0"}
```
