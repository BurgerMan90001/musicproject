import "./assets/css/App.css";
import { createContext } from "react";
import { Outlet } from "react-router";
import Header from "./components/Header";
interface ListContextType {
  animals: string[];
}
const ListContext = createContext<ListContextType>({
  animals: [],
});

function App() {
  const outlet = <Outlet />;
  
  return (
    <div className="grid-layout">
      <Header />
      {outlet}
      <footer className="footer"></footer>
    </div>
  );
}
export default App;

export { ListContext };
