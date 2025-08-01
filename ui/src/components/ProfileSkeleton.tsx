import React from "react";
import Skeleton from "./Skeleton";

const ProfileSkeleton: React.FC = () => {
  return (
    <div className="container-bare">
      <div className="max-w-2xl mx-auto">
        {/* Header Skeleton */}
        <div className="mb-8">
          <Skeleton className="mb-2" width="200px" height="36px" />
          <Skeleton width="300px" />
        </div>

        {/* Profile Information Skeleton */}
        <div className="card-bare p-6 mb-6">
          <Skeleton className="mb-4" width="180px" height="24px" />
          <div className="space-y-4">
            <div>
              <Skeleton className="mb-1" width="60px" height="14px" />
              <Skeleton width="250px" />
            </div>
            <div>
              <Skeleton className="mb-1" width="80px" height="14px" />
              <Skeleton width="180px" />
            </div>
          </div>
        </div>

        {/* Danger Zone Skeleton */}
        <div className="card-bare p-6 border-error/20">
          <Skeleton className="mb-4" width="120px" height="24px" />
          <Skeleton className="mb-4" lines={2} />
          <Skeleton width="140px" height="40px" variant="rectangular" />
        </div>
      </div>
    </div>
  );
};

export default ProfileSkeleton;