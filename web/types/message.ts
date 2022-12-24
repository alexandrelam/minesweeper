export enum MessageType {
  UPDATE_BOARD = "UPDATE_BOARD",
  CONNECTED_USERS = "CONNECTED_USERS",
  USER_MOUSE = "USER_MOUSE",
  CREATE_GAME = "CREATE_GAME",
  GAME_LOST = "GAME_LOST",
  GAME_WON = "GAME_WON",
}

export type Message<T = any> = {
  type: string;
  data: T;
};
