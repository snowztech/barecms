import React from "react";
import Loader from "@/components/Loader";
import { useAuth } from "@/contexts/AuthContext";
import useDeleteUser from "@/hooks/useDeleteUser";

const Profile: React.FC = () => {
  const { user, loading, error } = useAuth();
  const {
    isDeleting,
    error: deleteError,
    handleDelete,
  } = useDeleteUser(user?.id as string);

  if (loading || isDeleting) {
    return (
      <div className="container-bare">
        <div className="min-h-[400px] flex items-center justify-center">
          <Loader size="lg" variant="minimal" />
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="container-bare">
        <div className="alert-bare alert-bare-error p-4 rounded-bare">
          {error}
        </div>
      </div>
    );
  }

  return (
    <div className="container-bare">
      <div className="max-w-2xl mx-auto">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-display text-3xl font-semibold text-base-content mb-2">
            Profile
          </h1>
          <p className="text-bare-600">
            Manage your account settings and preferences
          </p>
        </div>

        {/* Profile Information */}
        <div className="card-bare p-6 mb-6">
          <h2 className="text-display text-xl font-medium text-base-content mb-4">
            Account Information
          </h2>
          {user ? (
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-bare-600 mb-1">
                  Email
                </label>
                <p className="text-base-content">{user.email}</p>
              </div>
              <div>
                <label className="block text-sm font-medium text-bare-600 mb-1">
                  Username
                </label>
                <p className="text-base-content">{user.username}</p>
              </div>
            </div>
          ) : (
            <p className="text-bare-600">No user data available</p>
          )}
        </div>

        {/* Danger Zone */}
        <div className="card-bare p-6 border-error/20">
          <h2 className="text-display text-xl font-medium text-error mb-4">
            Danger Zone
          </h2>
          <p className="text-bare-600 mb-4">
            Once you delete your account, there is no going back. Please be
            certain.
          </p>
          {deleteError && (
            <div className="alert-bare alert-bare-error p-3 rounded mb-4">
              {deleteError}
            </div>
          )}
          <button
            onClick={handleDelete}
            className="btn btn-outline border-error text-error hover:bg-error hover:text-error-content transition-all duration-200"
            disabled={isDeleting}
          >
            {isDeleting ? "Deleting..." : "Delete Account"}
          </button>
        </div>
      </div>
    </div>
  );
};

export default Profile;
