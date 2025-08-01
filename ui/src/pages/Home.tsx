import React, { useRef } from "react";
import CreateSiteModal from "@/components/modals/CreateSiteModal";
import { useAuth } from "@/contexts/AuthContext";
import { useSites } from "@/hooks/useSites";
import Loader from "@/components/Loader";

const HomePage: React.FC = () => {
  const { user, loading: userLoading } = useAuth();
  const { sites, loading: sitesLoading, error } = useSites();
  const modalRef = useRef<HTMLDialogElement>(null);

  const openModal = () => {
    if (modalRef.current) {
      modalRef.current.showModal();
    }
  };

  if (userLoading || sitesLoading) {
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

  if (!user) {
    return (
      <div className="container-bare">
        <div className="alert-bare alert-bare-warning p-4 rounded-bare">
          Please log in to continue
        </div>
      </div>
    );
  }

  return (
    <div className="container-bare">
      {/* Header Section */}
      <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 mb-8">
        <div>
          <h1 className="text-display text-3xl font-semibold text-base-content mb-2">
            My Sites
          </h1>
          <p className="text-bare-600 text-base">
            Manage your content collections and sites
          </p>
        </div>
        <button
          className="btn btn-primary px-6 py-3 font-medium shadow-bare hover:shadow-bare-lg transition-all duration-200"
          onClick={openModal}
        >
          + New Site
        </button>
      </div>

      {/* Sites Grid */}
      {sites.length === 0 ? (
        <div className="text-center py-16">
          <div className="max-w-md mx-auto">
            <h3 className="text-display text-xl font-medium text-base-content mb-4">
              No sites yet
            </h3>
            <p className="text-bare-600 mb-6">
              Create your first site to start managing content collections.
            </p>
            <button
              className="btn btn-primary px-6 py-3 font-medium"
              onClick={openModal}
            >
              Create First Site
            </button>
          </div>
        </div>
      ) : (
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          {sites.map((site) => (
            <div
              key={site.id}
              className="card-bare p-6 hover-lift group cursor-pointer"
            >
              <div className="flex flex-col h-full">
                <h3 className="text-display text-lg font-semibold text-base-content mb-2 group-hover:text-primary transition-colors">
                  {site.name}
                </h3>
                <p className="text-mono text-sm text-bare-600 mb-4 flex-1">
                  {site.slug}
                </p>
                <div className="flex justify-end">
                  <a
                    className="text-sm font-medium cursor-pointer text-primary hover:text-primary-focus transition-colors inline-flex items-center gap-1"
                    href={`/sites/${site.id}`}
                  >
                    Configure →
                  </a>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}

      <CreateSiteModal userId={user.id} dialogRef={modalRef} />
    </div>
  );
};

export default HomePage;
