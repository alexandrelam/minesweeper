import { useContext } from "react";
import { WebSocketContext } from "../provider/WebSocketProvider";
import { Tile as TileType, TileState } from "../types/tile";

type Props = {
  tile: TileType;
  row: number;
  column: number;
};

export function Tile({ tile, row, column }: Props) {
  const { websocket } = useContext(WebSocketContext);

  if (!websocket) {
    return null;
  }

  function display() {
    switch (tile.state) {
      case TileState.HIDDEN:
        return " ";
      case TileState.FLAGGED:
        return "🚩";
      case TileState.REVEALED:
        if (tile.isBomb) return "💣";
        if (tile.value === 0) return " ";

        return tile.value;
    }
  }

  function handleLeftClick() {
    if (!websocket) return;

    switch (tile.state) {
      case TileState.HIDDEN:
        websocket.send(
          JSON.stringify({
            action: "DIG",
            row,
            column,
          })
        );
    }
  }

  function handleRightClick(event: React.MouseEvent<HTMLDivElement>) {
    event.preventDefault();

    if (!websocket) return;

    switch (tile.state) {
      case TileState.HIDDEN:
        websocket.send(
          JSON.stringify({
            action: "FLAG",
            row,
            column,
          })
        );
      case TileState.FLAGGED:
        websocket.send(
          JSON.stringify({
            action: "UNFLAG",
            row,
            column,
          })
        );
    }
  }

  return (
    <div
      onClick={handleLeftClick}
      onContextMenu={handleRightClick}
      className="flex justify-center items-center 
                 w-10 h-10 bg-blue-100 border border-blue-300
                 rounded text-blue-500 text-lg font-bold"
    >
      {display()}
    </div>
  );
}
