import magnify from "../assets/images/magnify.svg";
import bell from "../assets/images/bell.svg";


function MediaButton({ label }: { label: string }) {
  return (
    <button className="bold font-large padding-16 transition-bg-color">
      {label}
    </button>
  );
}
function Form() {
  return (
    <form method="get" className="flex">
      <div className="flex bg-orange-900 border-radius-8 align-items-center">
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
    <header className="header bg-orange-700 box-shadow gap-8 padding-16-32 color-white">
      <Form />

      <div className="margin-l-auto flex">
        <div className="flex gap-4 margin-auto" id="top-media-buttons">
          <MediaButton label={"New"} />
          <MediaButton label={"Upload"} />
          <MediaButton label={"Share"} />
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
