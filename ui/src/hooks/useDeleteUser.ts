import { useState } from "react";
import apiClient from "@/lib/api";

const useDeleteUser = (userId: string) => {
  const [isDeleting, setIsDeleting] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const handleDelete = async () => {
    if (window.confirm("Are you sure you want to delete your account?")) {
      setIsDeleting(true);
      setError(null);
      try {
        const res = await apiClient.delete(`/user/${userId}`);
        if (res.status === 200) {
          localStorage.removeItem("token");
          window.location.href = "/login";
        } else {
          setError("Failed to delete resource");
        }
      } catch (error) {
        console.error(error);
        setError("Failed to delete resource");
      } finally {
        setIsDeleting(false);
      }
    }
  };

  return { isDeleting, error, handleDelete };
};

export default useDeleteUser;
