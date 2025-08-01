import React from "react";

interface SkeletonProps {
  className?: string;
  variant?: "text" | "rectangular" | "circular";
  width?: string | number;
  height?: string | number;
  lines?: number;
}

const Skeleton: React.FC<SkeletonProps> = ({
  className = "",
  variant = "text",
  width,
  height,
  lines = 1,
}) => {
  const baseClasses = "animate-pulse bg-bare-200 rounded";
  
  const variantClasses = {
    text: "h-4",
    rectangular: "h-32",
    circular: "rounded-full",
  };

  const skeletonStyle = {
    width: width || "100%",
    height: height || undefined,
  };

  if (variant === "text" && lines > 1) {
    return (
      <div className={`space-y-2 ${className}`}>
        {Array.from({ length: lines }).map((_, index) => (
          <div
            key={index}
            className={`${baseClasses} ${variantClasses.text}`}
            style={{
              ...skeletonStyle,
              width: index === lines - 1 ? "75%" : "100%",
            }}
          />
        ))}
      </div>
    );
  }

  return (
    <div
      className={`${baseClasses} ${variantClasses[variant]} ${className}`}
      style={skeletonStyle}
    />
  );
};

export default Skeleton;