import React from "react";

interface LoaderProps {
  size?: "xs" | "sm" | "md" | "lg";
  variant?: "default" | "minimal";
}

const Loader: React.FC<LoaderProps> = ({
  size = "md",
  variant = "default",
}) => {
  const sizeClasses = {
    xs: "w-4 h-4",
    sm: "w-6 h-6",
    md: "w-8 h-8",
    lg: "w-12 h-12",
  };

  if (variant === "minimal") {
    return (
      <div className="flex justify-center items-center">
        <div className={`loading-bare ${sizeClasses[size]}`}></div>
      </div>
    );
  }

  const sizeClass = `loading-spinner loading-${size}`;

  return (
    <div className="flex justify-center items-center">
      <span className={`loading ${sizeClass}`}></span>
    </div>
  );
};

export default Loader;
