import type { User } from "../types/auth.types";
import { createContext, useContext } from "react";

export interface AuthContextType {
  user?: User;
  isAuthenticated: boolean;
}

export const AuthContext = createContext<AuthContextType>({
  isAuthenticated: false,
});
export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);

  return context;
};
export default AuthContext;
