import { useContext } from "react";
import { createContext } from "react";

function ListItem({ animal }: { animal: string }) {
  return <li className="animal">{animal}</li>;
}

interface ListContextType {
  animals: string[];
}
const ListContext = createContext<ListContextType>({
  animals: [],
});
const List = () => {
  const { animals }: { animals: string[] } = useContext(ListContext);
  if (!animals) {
    return <div>Loading...</div>;
  }
  if (animals.length === 0) {
    return <div>There are no animals</div>;
  }
  return (
    <ul>
      {animals.map((animal) => {
        return <ListItem key={animal} animal={animal}></ListItem>;
      })}
    </ul>
  );
};

export default List;
