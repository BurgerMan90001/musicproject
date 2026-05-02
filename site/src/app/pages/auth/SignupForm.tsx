import { useState } from "react";
import type { FormInput } from "../../../types/form.types";
import api from "../../../lib/api";

const SignupForm = () => {
  const [error, setError] = useState("");
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
  const emailSignup = async (formData: FormData) => {
    const res = await api<string>("/auth/signup");

    console.log(res);
    // fetch(`${apiUrl}/auth/signup`, {
    //   method: "POST",
    //   body: JSON.stringify({
    //     email: formData.get("email"),
    //     password: formData.get("password"),
    //   }),
    // })
    //   .then((res) => res.json())
    //   .then(data) => d);
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

export default SignupForm;
