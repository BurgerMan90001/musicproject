import GoogleIcon from "../../../components/svg/GoogleIcon";

import OauthButton from "./OauthButton";
import GoogleSignupButton from "./OauthButton";
import SignupForm from "./SignupForm";

function Signup() {
  return (
    <>
      <div className="display-flex justifiy-content-center ">
        <main className="display-flex test  justifiy-content-center bg-color-body-dark padding-xxl">
          <div>
            <h1 className="border-bottom margin-block-xs padding-block-xs">
              Signup
            </h1>
            <section className="display-flex flex-column gap-xs">
              <OauthButton name="Signup with Google" icon={<GoogleIcon />} />
              <SignupForm />
            </section>
          </div>
        </main>
      </div>
    </>
  );
}

export default Signup;
