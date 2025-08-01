import React, { useEffect, useState, ReactNode } from "react";
import axios from "axios";
import apiClient from "@/lib/api";
import { AUTH_TOKEN_KEY, User } from "@/types/auth";
import { AuthContext, AuthContextType } from "@/contexts/AuthContext";

interface AuthResponse {
  token?: string;
  error?: string;
  user?: User;
}

interface AuthProviderProps {
  children: ReactNode;
}

export const AuthProvider: React.FC<AuthProviderProps> = ({ children }) => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchUser = async () => {
    const token = localStorage.getItem(AUTH_TOKEN_KEY);
    if (!token) {
      setUser(null);
      return;
    }

    try {
      const response = await apiClient.get("/user");
      setUser(response.data.user);
    } catch (err: any) {
      console.error("Failed to fetch user:", err);
      if (err.response?.status === 401) {
        localStorage.removeItem(AUTH_TOKEN_KEY);
        setUser(null);
      }
    }
  };

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
      setUser(user);
      setLoading(false);
      return { token, user };
    } catch (error: any) {
      setLoading(false);
      const errorMessage =
        error.response?.data?.error || "An error occurred. Please try again.";
      setError(errorMessage);
      return { error: errorMessage };
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
      const { token, user } = response.data;
      if (token) {
        localStorage.setItem(AUTH_TOKEN_KEY, token);
        setUser(user);
      }
      setLoading(false);
      return { token, user };
    } catch (err: any) {
      setLoading(false);
      const errorMessage =
        err.response?.data?.error || "An error occurred. Please try again.";
      setError(errorMessage);
      return { error: errorMessage };
    }
  };

  const logout = () => {
    localStorage.removeItem(AUTH_TOKEN_KEY);
    setUser(null);
    window.location.href = "/login";
  };

  const getAuthHeaders = () => {
    const token = localStorage.getItem(AUTH_TOKEN_KEY);
    return token ? { Authorization: `Bearer ${token}` } : {};
  };

  const refetchUser = async () => {
    await fetchUser();
  };

  useEffect(() => {
    fetchUser();
  }, []);

  const value: AuthContextType = {
    user,
    loading,
    error,
    login,
    register,
    logout,
    getAuthHeaders,
    refetchUser,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};
