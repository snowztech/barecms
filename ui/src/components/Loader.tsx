import React from "react";

interface LoaderProps {
  size?: "xs" | "sm" | "md" | "lg";
  variant?: "default" | "minimal" | "pulse" | "dots";
  className?: string;
}

const Loader: React.FC<LoaderProps> = ({
  size = "md",
  variant = "default",
  className = "",
}) => {
  const sizeClasses = {
    xs: "w-4 h-4",
    sm: "w-6 h-6",
    md: "w-8 h-8",
    lg: "w-12 h-12",
  };

  const dotSizes = {
    xs: "w-1 h-1",
    sm: "w-1.5 h-1.5",
    md: "w-2 h-2",
    lg: "w-3 h-3",
  };

  if (variant === "minimal") {
    return (
      <div className={`flex justify-center items-center ${className}`}>
        <div className={`loading-bare ${sizeClasses[size]}`}></div>
      </div>
    );
  }

  if (variant === "pulse") {
    return (
      <div className={`flex justify-center items-center ${className}`}>
        <div className={`loading-pulse ${sizeClasses[size]}`}></div>
      </div>
    );
  }

  if (variant === "dots") {
    return (
      <div
        className={`flex justify-center items-center space-x-1 ${className}`}
      >
        <div
          className={`loading-dot ${dotSizes[size]} animate-bounce`}
          style={{ animationDelay: "0ms" }}
        ></div>
        <div
          className={`loading-dot ${dotSizes[size]} animate-bounce`}
          style={{ animationDelay: "150ms" }}
        ></div>
        <div
          className={`loading-dot ${dotSizes[size]} animate-bounce`}
          style={{ animationDelay: "300ms" }}
        ></div>
      </div>
    );
  }

  // Default DaisyUI loader
  const sizeClass = `loading-spinner loading-${size}`;

  return (
    <div className={`flex justify-center items-center ${className}`}>
      <span className={`loading ${sizeClass} loading-primary`}></span>
    </div>
  );
};

export default Loader;
