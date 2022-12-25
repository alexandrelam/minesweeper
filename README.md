# :zap: Blazingly fast CO-OP Minesweeper :bomb:

![Capture d’écran 2022-12-25 011617](https://user-images.githubusercontent.com/25727549/209458735-7431f146-47da-4a90-ace0-4664db7df498.png)

![Capture d’écran 2022-12-25 011706](https://user-images.githubusercontent.com/25727549/209458754-86648c95-c657-4e9b-ba12-5487d516d98d.png)

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

```json
{
  "type": "UPDATE_BOARD",
  "data": "board[][]"
}

{
  "type": "CONNECTED_USERS",
  "data": "users[]"
}

{
  "type": "USER_MOUSE",
  "data": {
    "id": "id",
    "name": "name",
    "mouseX": 0,
    "mouseY": 0
  }
}
```

a update always contains a type and a data field

## Todo

- [ ] fix game create bomb
- [ ] random board
