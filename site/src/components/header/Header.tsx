import Navbar from "./Navbar";
// import SearchBar from "./SearchBar";
function Header() {
  return (
    <header className="display-block layout-header width-full">
      <div className="display-flex bg-color-body-dark align-items-center padding-block-xxs">
        <div className="display-flex align-items-center margin-left-auto gap-xxs"></div>
      </div>

      <Navbar />
    </header>
  );
}
{
  /* <SearchBar /> */
}
export default Header;
