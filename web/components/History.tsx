import { useContext } from "react";
import { HistoryContext } from "../provider/HistoryProvider";
import { ActionType } from "../types/action";
import { EventMessage } from "../types/message";

type HistoryItemProps = {
  event: EventMessage;
  setHoverTile: React.Dispatch<React.SetStateAction<[number, number]>>;
};

function HistoryItem({ event, setHoverTile }: HistoryItemProps) {
  function getIcon() {
    switch (event.action) {
      case ActionType.DIG:
        return "⛏️";
      case ActionType.FLAG:
        return "🚩";
      case ActionType.UNFLAG:
        return "🏳️";
      case ActionType.GAME_LOST:
        return "💥";
      case ActionType.GAME_WON:
        return "🎉";
    }
  }

  const hour = new Date(event.date).getHours();
  const minutes = new Date(event.date).getMinutes();
  const seconds = new Date(event.date).getSeconds();

  return (
    <li
      className="px-0.5 py-1 rounded cursor-default hover:bg-indigo-200"
      onMouseEnter={() => setHoverTile([event.column, event.row])}
      onMouseLeave={() => setHoverTile([-1, -1])}
    >
      {getIcon()} {event.authorName}{" "}
      <span className="font-medium">
        {event.column} {event.row}
      </span>
      {"   "}
      {hour}:{minutes}:{seconds}
    </li>
  );
}

type Props = {
  setHoverTile: React.Dispatch<React.SetStateAction<[number, number]>>;
};

export function History({ setHoverTile }: Props) {
  const { history } = useContext(HistoryContext);

  const sortedHistory = history.sort((a, b) => {
    return new Date(b.date).getTime() - new Date(a.date).getTime();
  });

  return (
    <div className="bg-indigo-100 p-4 pr-20 rounded text-indigo-500 h-96 overflow-y-scroll">
      <h2 className="font-bold text-lg mb-4">History</h2>
      <ul>
        {sortedHistory.map((event, index) => (
          <HistoryItem key={index} event={event} setHoverTile={setHoverTile} />
        ))}
      </ul>
    </div>
  );
}
