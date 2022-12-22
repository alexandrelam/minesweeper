# message

## request

```json
{
  "action": "CREATE_GAME"
}
```

```json
{
  "action": "DIG", // FLAG, UNFLAG
  "row": "0",
  "column": "0"
}
```

action: CREATE_GAME, FLAG, UNFLAG, DIG

## response

```
byte: board[][]

```
