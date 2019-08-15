# Description

```text
  特权账户查询，以列表形式查询所有guardians。
```

## Input

```bash
    qoscli query guardians --indent
```

## Output

```bash
    qoscli query guardians --indent
    [
      {
        "description":"this is the abc guardian",
        "guardian_type":1,
        "address":"address10xwx06gnrt3dlz7hfrx6a8wx3gyeghxm54rv7a",
        "creator":"address1ah9uz0"
      },
      {
        "description":"set adas to be a guardian",
        "guardian_type":2,
        "address":"address1l6juaqy9fk0dps0fn5dcg4fpy36zmryp8my4ux","creator":"address10xwx06gnrt3dlz7hfrx6a8wx3gyeghxm54rv7a"
      }
    ]
```
