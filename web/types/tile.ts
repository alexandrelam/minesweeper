export enum TileState {
  HIDDEN = 0,
  REVEALED = 1,
  FLAGGED = 2,
}

export type Tile = {
  state: TileState;
  value: number;
  isBomb: boolean;
};
