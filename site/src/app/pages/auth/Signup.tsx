import GoogleIcon from "../../../components/icons/GoogleIcon";
import { apiUrl } from "../../../config/env";
import { HTTPError } from "../../../lib/error";
import type { FormInput } from "../../../types/form.types";
import OauthButton from "./OauthButton";
import GoogleSignupButton from "./OauthButton";

function Signup() {
  const SignupForm = () => {
    const FormInput = (input: FormInput) => {
      return (
        <input
          name={input.name}
          type={input.type}
          className="bg-color-body-darker border font-size-md padding-inline-xs padding-block-xxs"
          placeholder={input.placeholder}
          required
        ></input>
      );
    };
    const emailSignup = (formData: FormData) => {
      fetch(`${apiUrl}/auth/signup`, {
        method: "POST",
        body: JSON.stringify({
          email: formData.get("email"),
          password: formData.get("password"),
        }),
      }).then((res) => {
        if (!res.ok) {
          HTTPError(res);
        }
        res.json();
      });
    };

    return (
      <>
        <span className="font-size-md font-weight-semilight color-text-subtle">
          Or signup with email
        </span>
        <form action={emailSignup} className="display-flex flex-column gap-xxs">
          <FormInput name="email" type="email" placeholder="Your Email" />
          <FormInput
            name="password"
            type="password"
            placeholder="Your Password"
          />

          <button
            className="button border font-size-md font-weight-semibold padding-xxs"
            type="submit"
          >
            Signup with email
          </button>
        </form>
      </>
    );
  };

  return (
    <>
      <div className="display-flex height-full justifiy-content-center ">
        <main className="display-flex flex-column justifiy-content-center bg-color-body-dark padding-xxl">
          <h1 className="border-bottom margin-block-xs padding-block-xs">
            Signup
          </h1>
          <section className="display-flex flex-column gap-xs">
            <OauthButton name="Signup with Google" icon={<GoogleIcon/>}/>
            <SignupForm />
          </section>
        </main>
      </div>
    </>
  );
}

export default Signup;
