import { useContext } from "react";
import { ListContext } from "../App";

function ListItem({ animal }: { animal: string }) {
  return <li className="animal">{animal}</li>;
}
function List() {
  const { animals } = useContext(ListContext);
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
}

export default List;
