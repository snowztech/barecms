import { AUTH_TOKEN_KEY } from "@/types/auth";
import React from "react";
import { Navigate, Outlet } from "react-router-dom";

const PrivateRoute: React.FC = () => {
  const token = localStorage.getItem(AUTH_TOKEN_KEY);
  return token ? <Outlet /> : <Navigate to="/login" />;
};

export default PrivateRoute;
