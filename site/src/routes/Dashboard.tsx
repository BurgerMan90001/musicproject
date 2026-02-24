//import { Link } from "react-router";
//import Image from "./Image";
//import List from "../components/List";
import { createContext } from "react";

interface ListContextType {
  animals: string[];
}
const ListContext = createContext<ListContextType>({
  animals: [],
});
/*
function Greeting({ username }: { username: string }) {
  return (
    <div className="flex align-items-center text-align-l line-height gap-16">
      <img className="size-40" src={username} />
      <div>
        <h3>Hi there,</h3>
        <h2>{username} (@morgan)</h2>
      </div>
    </div>
  );
}
function Test() {
  const animals = ["Lion", "Sheep", "Moose"];
  return (
    <ListContext value={{ animals }}>
      <main className="dashboard">
        <aside className="bg-orange-900 padding-8">
          <span>asdasdadadasdasdasdsaddasdawawdawdawd</span>
        </aside>
        <div className="color-black">
          <h1>Animals: </h1>
          <List />
          {<Image />}
          <Link to="/profile">Profile Page</Link>
        </div>
      </main>
    </ListContext>
  );
}
*/
function Dashboard() {
  return (
    <>
      <div className="" id="dashboard">
        <aside className="bg-orange-900 padding-8" id="sidebar">
          asdasdadadasdasdasdsaddasdawawdawdawd
        </aside>
        <main className="flex flex-column flex-1">
          <div className="grid-item-container gap-16 padding-16 ">
            <Item />
            <Item />
            <Item />
            <Item />
            <Item />
            <Item />
            <Item />
            <Item />
            <Item />
            <Item />
          </div>
        </main>
      </div>
    </>
  );
}
function Item() {
  return <div className="border-radius-8 height-250" id="test"></div>;
}
export default Dashboard;

export { ListContext };
