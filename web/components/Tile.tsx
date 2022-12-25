import { memo, useContext } from "react";
import { WebSocketContext } from "../provider/WebSocketProvider";
import { Tile as TileType, TileState } from "../types/tile";

type Props = {
  tile: TileType;
  row: number;
  column: number;
  isHighlighted?: boolean;
};

function Tile({ tile, row, column, isHighlighted }: Props) {
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

  function backgroundColor() {
    if (tile.state === TileState.HIDDEN) return "bg-slate-400";
    return "bg-slate-200";
  }

  function textColor() {
    switch (tile.value) {
      case 0:
        return "text-black";
      case 1:
        return "text-blue-500";
      case 2:
        return "text-green-500";
      case 3:
        return "text-red-500";
      case 4:
        return "text-purple-500";
      case 5:
        return "text-yellow-500";
      case 6:
        return "text-orange-500";
      case 7:
        return "text-pink-500";
      case 8:
        return "text-gray-500";
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
        break;
      case TileState.REVEALED:
        websocket.send(
          JSON.stringify({
            action: "DIG",
            row,
            column,
          })
        );
        break;
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
        break;
      case TileState.FLAGGED:
        websocket.send(
          JSON.stringify({
            action: "UNFLAG",
            row,
            column,
          })
        );
        break;
    }
  }

  return (
    <div
      onClick={handleLeftClick}
      onContextMenu={handleRightClick}
      className={`flex justify-center items-center 
                 w-10 h-10 ${backgroundColor()} border ${
        isHighlighted ? "border-yellow-500 border-2" : "border-slate-200"
      }
                 rounded ${textColor()} text-lg font-bold `}
    >
      {display()}
    </div>
  );
}

export const MemoizedTile = memo(Tile, (prevProps, nextProps) => {
  return (
    prevProps.tile.state === nextProps.tile.state &&
    prevProps.isHighlighted === nextProps.isHighlighted
  );
});
