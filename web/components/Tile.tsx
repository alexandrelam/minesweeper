import { Tile as TileType } from "../types/tile";

type Props = {
  tile: TileType;
};
export function Tile({ tile }: Props) {
  return (
    <div
      className="flex justify-center items-center 
                 w-10 h-10 bg-blue-100 border border-blue-300
                 rounded text-blue-500 text-lg font-bold"
    >
      {tile.value}
    </div>
  );
}
