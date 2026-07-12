import { useCallback, useState } from "react";
import apiClient from "@/lib/api";
import { AxiosRequestConfig } from "axios";

export const useApi = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const request = useCallback(async (config: AxiosRequestConfig) => {
    setLoading(true);
    setError(null);

    try {
      const response = await apiClient(config);
      return response.data;
    } catch (err: any) {
      const apiError = err.response?.data?.error;
      const errorMessage =
        (typeof apiError === "string" ? apiError : apiError?.message) ||
        err.response?.data?.message ||
        err.message;
      setError(errorMessage);
      throw new Error(errorMessage);
    } finally {
      setLoading(false);
    }
  }, []);

  return { request, loading, error };
};
