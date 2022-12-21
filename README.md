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
  "column": "0",
  "row": "0"
}
```

```json
{
  "action": "JOIN",
  "name": "alexandre"
}
```

action: CREATE_GAME, FLAG, UNFLAG, DIG, JOIN

## response

```
byte: board[][]

```
