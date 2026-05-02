// async function refreshTokens() {
//   fetch("http://localhost:8081/v1/auth/refresh");
// }

import { Link } from "react-router";
import api from "../../../lib/api";

async function onLogin() {
  const res = api<string>("/auth/login", {});
  // fetch("http://localhost:8081/v1/auth/login").then((res) => {
  //   console.log(res.headers.getSetCookie());
  // });
  console.log(res);
}

//function redirectToLogin() {}

function Login() {
  return (
    <>
      <main>
        <button onClick={onLogin}>Login</button>
        <Link to="/auth/reset">Forgot password</Link>
      </main>
    </>
  );
}

export default Login;
