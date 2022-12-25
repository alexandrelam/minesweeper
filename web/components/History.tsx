import { useContext } from "react";
import { HistoryContext } from "../provider/HistoryProvider";
import { ActionType } from "../types/action";
import { EventMessage } from "../types/message";

type HistoryItemProps = {
  event: EventMessage;
};

function HistoryItem({ event }: HistoryItemProps) {
  function getIcon() {
    switch (event.action) {
      case ActionType.DIG:
        return "⛏️";
      case ActionType.FLAG:
        return "🚩";
      case ActionType.UNFLAG:
        return "🏳️";
    }
  }

  const hour = new Date(event.date).getHours();
  const minutes = new Date(event.date).getMinutes();
  const seconds = new Date(event.date).getSeconds();

  return (
    <li>
      {getIcon()} {event.authorName}{" "}
      <span className="font-medium">
        {event.column} {event.row}
      </span>
      {"   "}
      {hour}:{minutes}:{seconds}
    </li>
  );
}

export function History() {
  const { history } = useContext(HistoryContext);

  const sortedHistory = history.sort((a, b) => {
    return new Date(b.date).getTime() - new Date(a.date).getTime();
  });

  return (
    <div className="bg-indigo-100 p-4 pr-20 rounded text-indigo-500 h-96 overflow-y-scroll">
      <h2 className="font-bold text-lg mb-4">History</h2>
      <ul>
        {sortedHistory.map((event, index) => (
          <HistoryItem key={index} event={event} />
        ))}
      </ul>
    </div>
  );
}
