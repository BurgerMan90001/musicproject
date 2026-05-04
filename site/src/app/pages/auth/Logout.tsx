import { useNavigate } from "react-router";
import { useAuthStore } from "../../../hooks/auth";

const Logout = () => {
  const auth = useAuthStore();
  const navigate = useNavigate();

  cookieStore.delete("accessToken");
  cookieStore.delete("refreshToken");
  auth.logout();

  navigate("/");
  return <div></div>;
};

export default Logout;
