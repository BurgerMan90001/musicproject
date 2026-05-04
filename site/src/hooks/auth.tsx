import { create } from "zustand";
import type { User } from "../types/auth.types";
import { devtools, persist } from "zustand/middleware";
import fetchApi from "../lib/api";

interface AuthProps {
  user?: User;
}
interface AuthState extends AuthProps {
  // refreshKey: string;
  // accessKey: string;
  // login: (user: User) => void;
  signup: (email: string, password: string) => Promise<Response | undefined>;
  logout: () => void;
  setUser: (user: User) => void;
  // handleCallback: (code: string, state: string) => Promise<void>;
  refresh: () => void;
}
export const useAuthStore = create<AuthState>()(
  devtools(
    persist(
      (set) => ({
        // accessKey: "",
        // refreshKey: "",
        setUser: (user: User) => set({ user }),
        signup: async (email: string, password: string) => {
          try {
            const res = await fetchApi("/v1/auth/signup", {
              method: "POST",
              body: JSON.stringify({
                email: email,
                password: password,
              }),
              credentials: "include",
            });
            return res;
          } catch (e) {
            console.log(e);
          }
        },
        login: (user: User) => {
          set({ user: user });
        },
        logout: () => {
          set(() => ({
            user: undefined,
          }));

          cookieStore.delete("accessToken");
          cookieStore.delete("refreshToken");
        },
        refresh: () => {
          // cookieStore.set()
        },

        // authenticated: () => {
        //   return get().user !== null;
        // },
      }),
      { name: "auth-store" },
    ),
  ),
);
