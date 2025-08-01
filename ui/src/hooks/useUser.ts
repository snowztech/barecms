import { useState, useEffect } from "react";
import apiClient from "@/lib/api";
import { User } from "@/types";
import { AUTH_TOKEN_KEY } from "@/types/auth";

export const useUser = () => {
  const [user, setUser] = useState<User | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const fetchUser = async () => {
    setLoading(true);
    setError(null);
    try {
      const token = localStorage.getItem(AUTH_TOKEN_KEY);
      if (!token) {
        setLoading(false);
        setUser(null);
        return;
      }
      const response = await apiClient.get("/user");
      console.log("use user", response.data);
      setUser(response.data.user);
    } catch (err: any) {
      setError(
        err.response?.data?.error || "An error occurred. Please try again.",
      );
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUser();
  }, []);

  return {
    user,
    loading,
    error,
    fetchUser,
  };
};
