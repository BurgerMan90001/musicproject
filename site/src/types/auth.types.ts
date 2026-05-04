export interface Login {
  email: string;
  password: string;
}
export interface User {
  id: string;
  email: string;
  username: string;
  avatar?: string;
  hero?: string;
}

export interface TokenResponse {
  access_token: string;
  refresh_token: string;
}

// export interface AuthContextType {
//   login: (provider: string) => Promise<void>;
//   logout: () => Promise<void>;
//   handleCallback: (code: string, state: string) => Promise<void>;
//   refreshAccessToken: () => Promise<void>;
// }
