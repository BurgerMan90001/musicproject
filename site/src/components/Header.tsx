import magnify from "../assets/images/magnify.svg";
import bell from "../assets/images/bell.svg";
import { Link } from "react-router";
function MediaButton({ label }: { label: string }) {
  return (
    <Link to="/upload" className="bold font-large padding-16 ">
      {label}
    </Link>
  );
}
function SearchBar() {
  return (
    <form method="get" className="flex">
      <div className="flex bg-black2 align-items-center ">
        <button className="size-40 padding-8" type="submit">
          <img src={magnify} />
        </button>
        <input
          type="text"
          className="border-radius-8 font-size-15 font-weight-500"
          name="query"
          placeholder="Search something"
        />
      </div>
    </form>
  );
}
function Header() {
  const companyName = "Company";
  return (
    <header className="header bg-black box-shadow gap-8 padding-16-32 color-white">
      <SearchBar />

      <div className="margin-l-auto flex">
        <div className="flex gap-4 margin-auto">
          <MediaButton label={"Upload"} />
        </div>
      </div>

      <div className="flex gap-8 align-items-center">
        <button className="size-40 padding-8" title="Notifications">
          <img src={bell} />
        </button>
        <h3>{companyName}</h3>
      </div>
    </header>
  );
}

export default Header;
