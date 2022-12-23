enum TileState {
  Hidden = 0,
  Visible = 1,
  Flagged = 2,
}

export type Tile = {
  state: TileState;
  value: number;
  isBomb: boolean;
};
