import GoogleIcon from "../../../components/svg/GoogleIcon";
import { API_URL } from "../../../config/env";
import SignupButton from "./Button";


function Signup() {
  window.location.href = `${API_URL}/v1/auth/songsled/login`;
  return (
    <>
      <div className="display-flex justifiy-content-center ">
        <main className="display-flex test  justifiy-content-center bg-color-body-dark padding-xxl">
          <div>
            <h1 className="border-bottom margin-block-xs padding-block-xs">
              Signup
            </h1>
            <section className="display-flex flex-column gap-xs">
              <SignupButton
                path="/v1/auth/google/login"
                name="Signup with Google"
                icon={<GoogleIcon />}
              />
              <span className="font-size-md font-weight-semilight color-text-subtle">
                Or signup with email
              </span>
              {/* <SignupForm /> */}
              <SignupButton
                // className="button border font-size-md font-weight-semibold padding-xxs"
                path="/v1/auth/songsled/login"
                name="Signup with email"
              />
            </section>
          </div>
        </main>
      </div>
    </>
  );
}

export default Signup;
