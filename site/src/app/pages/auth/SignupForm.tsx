import { useState } from "react";
import type { ApiError } from "../../../lib/error";
import { useNavigate } from "react-router";
import { useAuthStore } from "../../../hooks/auth";
import type { User } from "../../../types/auth.types";

const SignupForm = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const navigate = useNavigate();
  const auth = useAuthStore();

  const action = async () => {
    const res = await auth.signup(email, password);

    if (!res) {
      return;
    }
    const json = await res.json();
    if (res.ok) {
      const user = json as unknown as User;
      auth.setUser(user);
      auth.setUser({
        id: "",
        email: email,
        username: "",
      });

      navigate("/");
    }

    const error = json as unknown as ApiError;

    setError(error.message);
  };

  return (
    <>
      <form action={action} className="display-flex flex-column gap-xxs">
        {/* <input
          type="username"
          aria-label="Your Username"
          className="bg-color-body-darker padding-xxs"
          placeholder="Your Username"
          onChange={(e) => setUsername(e.target.value)}
        ></input> */}
        <input
          type="email"
          aria-label="Your Email"
          className="bg-color-body-darker padding-xxs"
          placeholder="Your Email"
          onChange={(e) => setEmail(e.target.value)}
          required
        ></input>

        <input
          type="password"
          aria-label="Your Password"
          className="bg-color-body-darker padding-xxs"
          placeholder="Your Password"
          onChange={(e) => setPassword(e.target.value)}
          required
        ></input>

        <span className="color-text-danger">{error}</span>
      </form>
    </>
  );
};

export default SignupForm;
