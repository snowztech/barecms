import React, { useState } from "react";
import { useApi } from "@/hooks/useApi";

interface CreateSiteModalProps {
  userId: string;
  dialogRef: React.RefObject<HTMLDialogElement>;
}

const CreateSiteModal: React.FC<CreateSiteModalProps> = ({
  userId,
  dialogRef,
}) => {
  const [siteName, setSiteName] = useState("");
  const [error, setError] = useState<string | null>(null);
  const { request, loading } = useApi();

  const handleSiteNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSiteName(e.target.value);
    setError(null);
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
      return;
    }

    try {
      await request({
        url: "/sites",
        method: "POST",
        data: { name: siteName, userId },
      });

      console.log("Site created successfully");
      closeDialog();
      setTimeout(() => {
        window.location.reload();
      }, 300);
    } catch (e: any) {
      console.error(e);
      setError(e.message || "Failed to create site. Please try again.");
    } finally {
      setSiteName("");
    }
  };

  return (
    <dialog className="modal" ref={dialogRef}>
      <div className="modal-box">
        <h3 className="font-bold text-lg mb-2">Create new site</h3>
        <input
          type="text"
          placeholder="Enter site name"
          className="input input-bordered w-full"
          value={siteName}
          onChange={handleSiteNameChange}
        />
        {error && <p className="text-red-500 mt-2">{error}</p>}
        <div className="modal-action">
          <button
            disabled={loading}
            onClick={handleSubmit}
            className="btn btn-primary"
          >
            {loading ? "Creating..." : "Create"}
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
