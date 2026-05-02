import { create } from "zustand";
import type { User } from "../types/auth.types";

interface AuthState {
  user: User | null;
  authenticated: () => boolean;
  login: (provider: string) => void;
  logout: () => void;
  // handleCallback: (code: string, state: string) => Promise<void>;
  refresh: () => void;
}
export const useAuthStore = create<AuthState>()((set, get) => ({
  user: null,
  login: (provider: string) => {
    switch (provider) {
    }
  },
  logout: () => {
    set(() => ({
      user: undefined,
    }));
    cookieStore.delete("");
  },
  refresh: () => {
    // cookieStore.set()
  },

  authenticated: () => {
    return get().user === undefined;
  },
}));
