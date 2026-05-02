import Navbar from "./Navbar";
import SearchBar from "./SearchBar";
function Header() {
  return (
    <header className="display-block layout-header width-full">
      <div className="display-flex bg-color-body-dark align-items-center padding-block-xxs">
        <SearchBar />

        <div className="display-flex align-items-center margin-left-auto gap-xxs">
          {/* <button className="button-clear" title="Notifications">
            <img src={bell} className="icon font-size-xxl" />
          </button> */}

          {/* <p className="bold">{companyName}</p> */}
        </div>
      </div>

      <Navbar />
    </header>
  );
}

export default Header;
