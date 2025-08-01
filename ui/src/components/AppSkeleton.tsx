import React from "react";
import Skeleton from "./Skeleton";

const AppSkeleton: React.FC = () => {
  return (
    <div className="min-h-screen bg-base-100 text-base-content flex flex-col">
      {/* Header Skeleton */}
      <header className="border-b border-bare-200">
        <div className="container-bare flex justify-between items-center py-6">
          <div className="flex items-center gap-2">
            <Skeleton variant="rectangular" width="28px" height="28px" />
            <Skeleton width="120px" height="32px" />
          </div>
          <div className="flex items-center gap-4">
            <Skeleton variant="circular" width="40px" height="40px" />
            <Skeleton width="100px" height="20px" />
          </div>
        </div>
      </header>

      {/* Main Content Skeleton */}
      <main className="flex-1 py-8">
        <div className="container-bare">
          <div className="max-w-2xl mx-auto">
            <div className="mb-8">
              <Skeleton className="mb-2" width="200px" height="36px" />
              <Skeleton width="300px" />
            </div>
            <div className="card-bare p-6 mb-6">
              <Skeleton className="mb-4" width="180px" height="24px" />
              <div className="space-y-4">
                <Skeleton lines={2} />
                <Skeleton lines={2} />
              </div>
            </div>
          </div>
        </div>
      </main>

      {/* Footer Skeleton */}
      <footer className="border-t border-bare-200 py-4">
        <div className="container-bare text-center">
          <Skeleton width="150px" height="16px" className="mx-auto" />
        </div>
      </footer>
    </div>
  );
};

export default AppSkeleton;