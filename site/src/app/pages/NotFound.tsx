import { Link } from "react-router";

import Header from "../../components/header/Header";
function ErrorPage() {
  return (
    <div className="display-flex flex-column layout-single height-full">
      <Header />
      <main className="margin-auto">
        <h1>Resource not found!</h1>
        <Link to="/" className="button-clear">
          Go back to the home page?
        </Link>
      </main>
    </div>
  );
}

export default ErrorPage;
