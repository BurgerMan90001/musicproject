import Header from "../../components/header/Header";
import { apiUrl } from "../../config/env";
import type { User } from "../../types/auth.types";

function Profile() {
  fetch(`${apiUrl}/users`, {
    method: "GET",
  }).then((res) => {
    if (!res.ok) {
      throw new Error();
    }
    res.json();
  });

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
