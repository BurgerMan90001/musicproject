import Header from "../../components/header/Header";
import api from "../../lib/api";
import type { User } from "../../types/auth.types";

function Profile() {
  const res = api<string>("/users", {
    method: "GET",
  });
  console.log(res);

  const user: User = {
    id: "",
    username: "Paul Casigay",
    email: "paulcasigay@gmail.com",
    provider: "",
  };

  var hero = null;
  if (!user.hero) {
    hero = (
      <section className="bg-color-primary width-full hero">asdasdasd</section>
    );
  }

  if (!user.avatar)
    return (
      <div className="height-full display-block">
        <Header />
        <main className="margin-inline-xxxl display-block ">
          {hero}
          <section className="padding-xxl bg-color-body-dark margin-inline-md">
            <div className=" font-weight-semibold">
              <span className="">{user.username}</span>
            </div>
          </section>
        </main>
      </div>
    );
}

export default Profile;
