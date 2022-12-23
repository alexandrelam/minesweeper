export enum MessageType {
  UPDATE_BOARD = "UPDATE_BOARD",
  CONNECTED_USERS = "CONNECTED_USERS",
  USER_MOUSE = "USER_MOUSE",
}

export type Message<T = any> = {
  type: string;
  data: T;
};
