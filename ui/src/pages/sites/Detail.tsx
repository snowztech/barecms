import React, { useRef } from "react";
import { useParams } from "react-router-dom";
import CreateCollectionModal from "@/components/modals/CreateCollectionModal";
import { useSiteDetail } from "@/hooks/useSiteDetail";
import Loader from "@/components/Loader";
import useDelete from "@/hooks/useDelete";
import ViewSiteDataModal from "@/components/modals/ViewSiteDataModal";

const SiteDetailsPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const { site, collections, loading, error } = useSiteDetail(id);
  const { isDeleting, handleDelete } = useDelete(`/sites/${id || ""}`, "/");

  const collectionModalRef = useRef<HTMLDialogElement>(null);
  const viewDataModalRef = useRef<HTMLDialogElement>(null);

  const openCollectionModal = () => {
    if (collectionModalRef.current) {
      collectionModalRef.current.showModal();
    }
  };

  const openDataModal = () => {
    if (viewDataModalRef.current) {
      viewDataModalRef.current.showModal();
    }
  };

  if (loading) {
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
        <div className="alert-bare alert-bare-error p-4 rounded-bare mb-4">
          {error}
        </div>
        <a href="/" className="btn btn-primary">
          ← Back to Sites
        </a>
      </div>
    );
  }

  if (!site) {
    return (
      <div className="container-bare">
        <div className="alert-bare alert-bare-warning p-4 rounded-bare mb-4">
          Site not found
        </div>
        <a href="/" className="btn btn-primary">
          ← Back to Sites
        </a>
      </div>
    );
  }

  return (
    <div className="container-bare">
      {/* Breadcrumbs */}
      <nav className="mb-6">
        <div className="breadcrumbs text-sm">
          <ul className="flex items-center space-x-2 text-bare-600">
            <li>
              <a href="/" className="hover:text-primary transition-colors">
                My Sites
              </a>
            </li>
            <li className="before:content-['/'] before:mx-2">
              <span className="text-base-content font-medium">{site.name}</span>
            </li>
          </ul>
        </div>
      </nav>

      {/* Header Section */}
      <div className="flex flex-col lg:flex-row justify-between items-start lg:items-center gap-4 mb-8">
        <div className="flex-1">
          <h1 className="text-display text-3xl font-semibold text-base-content mb-2">
            {site.name}
          </h1>
          <div className="space-y-1">
            <p className="text-mono text-sm text-bare-600">
              <span className="font-medium">ID:</span> {site.id}
            </p>
            <p className="text-mono text-sm text-bare-600">
              <span className="font-medium">Slug:</span> {site.slug}
            </p>
          </div>
        </div>

        <div className="flex flex-wrap gap-3">
          <button className="btn btn-primary px-4 py-2" onClick={openDataModal}>
            View Site Data
          </button>
          <div className="dropdown dropdown-end">
            <div tabIndex={0} role="button" className="btn btn-bare px-4 py-2">
              Settings
            </div>
            <ul
              tabIndex={0}
              className="dropdown-content dropdown-bare menu bg-base-100 rounded-bare z-[1] p-2 w-40 shadow-bare-lg mt-2"
            >
              <li>
                <button
                  onClick={handleDelete}
                  className="text-sm text-error hover:bg-error/10 transition-colors w-full text-left"
                  disabled={isDeleting}
                >
                  {isDeleting ? "Deleting..." : "Delete Site"}
                </button>
              </li>
            </ul>
          </div>
        </div>
      </div>

      {/* Collections Section */}
      <div className="space-y-6">
        <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
          <div>
            <h2 className="text-display text-2xl font-semibold text-base-content mb-2">
              Collections
            </h2>
            <p className="text-bare-600">
              Organize your content into collections
            </p>
          </div>
          <button
            className="btn btn-primary px-6 py-3 font-medium"
            onClick={openCollectionModal}
          >
            + New Collection
          </button>
        </div>

        {collections.length === 0 ? (
          <div className="text-center py-16">
            <div className="max-w-md mx-auto">
              <h3 className="text-display text-xl font-medium text-base-content mb-4">
                No collections yet
              </h3>
              <p className="text-bare-600 mb-6">
                Create your first collection to start organizing content.
              </p>
              <button
                className="btn btn-primary px-6 py-3 font-medium"
                onClick={openCollectionModal}
              >
                Create First Collection
              </button>
            </div>
          </div>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {collections.map((collection) => (
              <div
                key={collection.id}
                className="card-bare p-6 hover-lift group cursor-pointer"
              >
                <div className="flex flex-col h-full">
                  <h3 className="text-display text-lg font-semibold text-base-content mb-2 group-hover:text-primary transition-colors">
                    {collection.name}
                  </h3>
                  <p className="text-mono text-sm text-bare-600 mb-4 flex-1">
                    {collection.slug}
                  </p>
                  <div className="flex justify-end">
                    <a
                      className="text-sm font-medium text-primary hover:text-primary-focus transition-colors inline-flex items-center gap-1"
                      href={`/sites/${site.id}/collections/${collection.id}`}
                    >
                      View →
                    </a>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>

      <CreateCollectionModal siteId={site.id} dialogRef={collectionModalRef} />
      <ViewSiteDataModal siteSlug={site.slug} dialogRef={viewDataModalRef} />
    </div>
  );
};

export default SiteDetailsPage;
