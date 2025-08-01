import { useState } from "react";
import axios from "axios";
import { AUTH_TOKEN_KEY, User } from "@/types/auth";

interface AuthResponse {
  token?: string;
  error?: string;
  user?: User;
}

export const useAuth = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const login = async (
    email: string,
    password: string,
  ): Promise<AuthResponse> => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.post("/api/auth/login", { email, password });
      const { token, user } = response.data;
      localStorage.setItem(AUTH_TOKEN_KEY, token);
      setLoading(false);
      return { token, user };
    } catch (error: any) {
      setLoading(false);
      setError(
        error.response?.data?.error || "An error occurred. Please try again.",
      );
      return {
        error:
          error.response?.data?.error || "An error occurred. Please try again.",
      };
    }
  };

  const register = async (
    email: string,
    username: string,
    password: string,
  ): Promise<AuthResponse> => {
    setLoading(true);
    setError(null);
    try {
      const response = await axios.post("/api/auth/register", {
        email,
        username,
        password,
      });
      setLoading(false);

      const { token, user } = response.data;
      if (token) {
        localStorage.setItem(AUTH_TOKEN_KEY, token);
        return { token, user };
      }

      return { user };
    } catch (err: any) {
      setLoading(false);
      setError(
        err.response?.data?.error || "An error occurred. Please try again.",
      );
      return {
        error:
          err.response?.data?.error || "An error occurred. Please try again.",
      };
    }
  };

  const logout = () => {
    localStorage.removeItem(AUTH_TOKEN_KEY);
    window.location.href = "/login";
  };

  const getAuthHeaders = () => {
    const token = localStorage.getItem(AUTH_TOKEN_KEY);
    return token ? { Authorization: `Bearer ${token}` } : {};
  };

  return {
    loading,
    error,
    login,
    register,
    logout,
    getAuthHeaders,
  };
};
