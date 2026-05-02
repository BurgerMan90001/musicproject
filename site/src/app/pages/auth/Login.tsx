// async function refreshTokens() {
//   fetch("http://localhost:8081/v1/auth/refresh");
// }

import { Link } from "react-router";
import Header from "../../../components/header/Header";

async function onLogin() {
  fetch("http://localhost:8081/v1/auth/login").then((res) => {
    console.log(res.headers.getSetCookie());
  });
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
