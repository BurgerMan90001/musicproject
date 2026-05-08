import { API_URL } from "../../../config/env";
//function redirectToLogin() {}

function Login() {
  // const [email, setEmail] = useState("");
  // const [password, setPassword] = useState("");

  // const onLogin = () => {
  //   window.location.href = `${API_URL}"/v1/auth/login"`;
  // const res = await fetchApi("/v1/auth/login", {
  //   method: "POST",
  //   body: JSON.stringify({
  //     email: email,
  //     password: password,
  //   }),
  // });

  // console.log(res);
  // };
  window.location.href = `${API_URL}/v1/auth/songsled/login`;
  return <div></div>;
  // return (
  //   <div className="display-flex justifiy-content-center ">
  //     <main className="display-flex test  justifiy-content-center bg-color-body-dark padding-xxl">
  //       <div>
  //         <h1 className="border-bottom margin-block-xs padding-block-xs">
  //           Login
  //         </h1>
  //         <section className="display-flex flex-column gap-xs">
  //           <Button
  //             path="/auth/google/login"
  //             name="Login with Google"
  //             icon={<GoogleIcon />}
  //           />
  //           <SignupForm />
  //           {/* <button onClick={window.location.href}></button> */}
  //           {/* <Link to="/auth/reset">Forgot password?</Link> */}
  //         </section>
  //       </div>
  //     </main>
  //   </div>
  // );
}

export default Login;
