import React, { useState } from "react";
import { ApiRequestError, useApi } from "@/hooks/useApi";

interface CreateSiteModalProps {
  userId?: string;
  dialogRef: React.RefObject<HTMLDialogElement>;
  siteId?: string;
  initialName?: string;
}

const CreateSiteModal: React.FC<CreateSiteModalProps> = ({
  userId,
  dialogRef,
  siteId,
  initialName = "",
}) => {
  const [siteName, setSiteName] = useState(initialName);
  const [error, setError] = useState<string | null>(null);
  const [nameError, setNameError] = useState<string | null>(null);
  const { request, loading } = useApi();

  const handleSiteNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSiteName(e.target.value);
    setError(null);
    setNameError(null);
  };

  const closeDialog = () => {
    if (dialogRef.current) {
      dialogRef.current.close();
    }
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    if (siteName.trim() === "") {
      setError("Site name cannot be empty.");
      setNameError("is required");
      return;
    }

    try {
      await request({
        url: siteId ? `/sites/${siteId}` : "/sites",
        method: siteId ? "PUT" : "POST",
        data: { name: siteName, userId },
      });

      closeDialog();
      setTimeout(() => {
        window.location.reload();
      }, 300);
    } catch (e: unknown) {
      if (e instanceof ApiRequestError) setNameError(e.fields.name || null);
      setError(e instanceof Error ? e.message : "Failed to save site.");
    } finally {
      if (!siteId) setSiteName("");
    }
  };

  return (
    <dialog className="modal" ref={dialogRef}>
      <div className="modal-box">
        <h3 className="font-bold text-lg mb-2">{siteId ? "Edit site" : "Create new site"}</h3>
        <input
          type="text"
          placeholder="Enter site name"
          className={`input input-bordered w-full ${nameError ? "input-error" : ""}`}
          value={siteName}
          onChange={handleSiteNameChange}
          aria-invalid={Boolean(nameError)}
          aria-describedby={nameError ? "site-name-error" : undefined}
        />
        {nameError && <p id="site-name-error" role="alert" className="text-sm text-error mt-1">{nameError}</p>}
        {error && <p className="text-red-500 mt-2">{error}</p>}
        <div className="modal-action">
          <button
            disabled={loading}
            onClick={handleSubmit}
            className="btn btn-primary"
          >
            {loading ? "Saving..." : siteId ? "Save changes" : "Create"}
          </button>
          <button className="btn" onClick={closeDialog}>
            Cancel
          </button>
        </div>
      </div>
    </dialog>
  );
};

export default CreateSiteModal;
