import React, { useRef } from "react";
import { useParams } from "react-router-dom";
import EntryCard from "@/components/cards/EntryCard";
import Loader from "@/components/Loader";
import CreateEntryModal from "@/components/modals/CreateEntryModal";
import { useCollectionDetail } from "@/hooks/useCollectionDetail";
import useDelete from "@/hooks/useDelete";

const CollectionDetailsPage: React.FC = () => {
  const { id, siteId } = useParams<{ id: string; siteId: string }>();
  const { collection, entries, site, loading, error } = useCollectionDetail(
    id,
    siteId,
  );

  const { isDeleting, handleDelete } = useDelete(
    `/collections/${id || ""}`,
    `/sites/${siteId || ""}`,
  );

  const entryModalRef = useRef<HTMLDialogElement>(null);

  const openEntryModal = () => {
    if (entryModalRef.current) {
      entryModalRef.current.showModal();
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
        <a href={`/sites/${siteId}`} className="btn btn-primary">
          ← Back to Site
        </a>
      </div>
    );
  }

  if (!collection) {
    return (
      <div className="container-bare">
        <div className="alert-bare alert-bare-warning p-4 rounded-bare mb-4">
          Collection not found
        </div>
        <a href={`/sites/${siteId}`} className="btn btn-primary">
          ← Back to Site
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
              <a
                href={`/sites/${siteId}`}
                className="hover:text-primary transition-colors"
              >
                {site?.name || "Site"}
              </a>
            </li>
            <li className="before:content-['/'] before:mx-2">
              <span className="text-base-content font-medium">
                {collection.name}
              </span>
            </li>
          </ul>
        </div>
      </nav>

      {/* Header Section */}
      <div className="flex flex-col lg:flex-row justify-between items-start lg:items-center gap-4 mb-8">
        <div className="flex-1">
          <h1 className="text-display text-3xl font-semibold text-base-content mb-2">
            {collection.name}
          </h1>
          <div className="space-y-1">
            <p className="text-mono text-sm text-bare-600">
              <span className="font-medium">ID:</span> {collection.id}
            </p>
            <p className="text-mono text-sm text-bare-600">
              <span className="font-medium">Slug:</span> {collection.slug}
            </p>
          </div>
        </div>

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
                {isDeleting ? "Deleting..." : "Delete Collection"}
              </button>
            </li>
          </ul>
        </div>
      </div>

      {/* Fields Section */}
      {collection.fields && collection.fields.length > 0 && (
        <div className="mb-8">
          <h2 className="text-display text-xl font-semibold text-base-content mb-4">
            Fields
          </h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
            {collection.fields.map((field) => (
              <div key={field.name} className="card-bare p-4">
                <h3 className="text-display text-base font-medium text-base-content mb-1">
                  {field.name}
                </h3>
                <p className="badge-bare">{field.type}{field.optional && ', optional'}</p>
              </div>
            ))}
          </div>
        </div>
      )}

      {/* Entries Section */}
      <div className="space-y-6">
        <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
          <div>
            <h2 className="text-display text-2xl font-semibold text-base-content mb-2">
              Entries
            </h2>
            <p className="text-bare-600">Content entries for this collection</p>
          </div>
          <button
            className="btn btn-primary px-6 py-3 font-medium"
            onClick={openEntryModal}
          >
            + New Entry
          </button>
        </div>

        {entries.length > 0 ? (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {entries.map((entry) => (
              <EntryCard
                key={entry.id}
                siteId={siteId as string}
                collectionId={collection.id}
                entryId={entry.id}
                data={entry.data}
              />
            ))}
          </div>
        ) : (
          <div className="text-center py-16">
            <div className="max-w-md mx-auto">
              <h3 className="text-display text-xl font-medium text-base-content mb-4">
                No entries yet
              </h3>
              <p className="text-bare-600 mb-6">
                Create your first entry to add content to this collection.
              </p>
              <button
                className="btn btn-primary px-6 py-3 font-medium"
                onClick={openEntryModal}
              >
                Create First Entry
              </button>
            </div>
          </div>
        )}
      </div>

      <CreateEntryModal
        dialogRef={entryModalRef}
        collectionId={collection.id}
        fields={collection.fields}
      />
    </div>
  );
};

export default CollectionDetailsPage;
