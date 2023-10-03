import { useEffect, useState } from "react";
import "./App.css";

// Import service definition that you want to connect to.
import { ToDoService } from "./gen/todo/v1/todo_connect";
import { useClient } from "./useClient";
import { ToDoItem } from "./gen/todo/v1/todo_pb";

function App() {
  const [inputValue, setInputValue] = useState("");
  const toDoClient = useClient(ToDoService);
  const [toDos, setToDos] = useState<ToDoItem[]>([]);

  useEffect(() => {
    toDoClient.getToDos({}).then((toDos) => {
      setToDos(toDos.toDos);
    });
  }, [toDoClient]);

  const handleChangeToDoCompleted = (id: string) => {
    console.log(`changing id ${id} to completed`);
    toDoClient.completeToDo({ id }).then((toDos) => {
      setToDos(toDos.toDos);
    });
  };

  return (
    <>
      <ol>
        {toDos.map((toDo) => (
          <li
            key={toDo.id}
            style={{ textDecoration: toDo.completed ? "line-through" : "" }}
          >
            <input
              type="checkbox"
              checked={toDo.completed}
              onChange={() => handleChangeToDoCompleted(toDo.id)}
            />
            {toDo.toDo}
          </li>
        ))}
      </ol>
      <form
        onSubmit={async (e) => {
          e.preventDefault();
          // Clear inputValue since the user has submitted.
          setInputValue("");
          // Store the inputValue in the chain of messages and
          // mark this message as coming from "me"
          const response = await toDoClient.createToDo({
            toDo: inputValue,
          });
          setToDos(response.toDos);
        }}
      >
        <input
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
        />
        <button type="submit">Send</button>
      </form>
    </>
  );
}

export default App;
