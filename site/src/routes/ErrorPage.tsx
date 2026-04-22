import { Link } from "react-router";

import Header from "../components/Header";

function ErrorPage() {
  return (
    <>
      <div className="grid-layout color-orange-700">
        <Header />
        <main className="bg-white">
          <div className="full-height flex justifiy-content-center flex-column">
            <h1>Resource not found!</h1>
            <Link to="/">Go back to the home page?</Link>
          </div>
        </main>
      </div>
    </>
  );
}

export default ErrorPage;
