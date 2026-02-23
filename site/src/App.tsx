import "./assets/css/App.css";
import { createContext } from "react";
import { Outlet } from "react-router";
import Header from "./components/Header";
import ErrorPage from "./routes/ErrorPage";
import Dashboard from "./routes/Dashboard";
interface ListContextType {
  animals: string[];
}
const ListContext = createContext<ListContextType>({
  animals: [],
});

function App() {
  let outlet = <Outlet />;
  // if (test == null) {
  //   test = <Dashboard/>;
  // }

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
