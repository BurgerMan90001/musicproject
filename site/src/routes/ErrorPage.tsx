import { Link } from "react-router";

import Header from "../components/Header";

function ErrorPage() {
  return (
    <>
      <div className="grid-layout color-orange-700">
        <Header />
        <main className="bg-white">
          <div className="full-height flex justifiy-content-center flex-column">
            <h1>Oh no, this route doesn't exist!</h1>
            <Link to="/">
              You can go back to the home page by clicking here, though!
            </Link>
          </div>
        </main>
      </div>
    </>
  );
}

export default ErrorPage;
