import { createContext, useContext } from "react";
import { User } from "@/types/auth";

interface AuthResponse {
  token?: string;
  error?: string;
  user?: User;
}

export interface AuthContextType {
  user: User | null;
  loading: boolean;
  error: string | null;
  login: (email: string, password: string) => Promise<AuthResponse>;
  register: (
    email: string,
    username: string,
    password: string,
  ) => Promise<AuthResponse>;
  logout: () => void;
  getAuthHeaders: () => { Authorization?: string };
  refetchUser: () => Promise<void>;
}

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
