import { EventMessage } from "../types/message";
import { useState, createContext, SetStateAction } from "react";

type HistoryContextType = {
  history: EventMessage[];
  setHistory: React.Dispatch<SetStateAction<EventMessage[]>>;
};

const HistoryContext = createContext<HistoryContextType>({
  history: [],
  setHistory: () => {},
});

function HistoryProvider(props: { children: React.ReactNode }) {
  const [history, setHistory] = useState<EventMessage[]>([]);

  return (
    <HistoryContext.Provider value={{ history, setHistory }}>
      {props.children}
    </HistoryContext.Provider>
  );
}

export { HistoryContext, HistoryProvider };
