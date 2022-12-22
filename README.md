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
board[][]

```

## mouse update

```json
{
  "id": "id",
  "name": "name",
  "mouseX": 0,
  "mouseY": 0
}
```

This data is send by a user and received by all other users
