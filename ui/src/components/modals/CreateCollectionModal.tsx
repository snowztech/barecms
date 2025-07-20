import React, { useState } from "react";
import { useApi } from "@/hooks/useApi";
import { Field, FieldType, VALID_FIELD_TYPES } from "@/types/fields";

interface CreateCollectionModalProps {
  siteId: string;
  dialogRef: React.RefObject<HTMLDialogElement>;
}

const CreateCollectionModal: React.FC<CreateCollectionModalProps> = ({
  siteId,
  dialogRef,
}) => {
  const [collectionName, setCollectionName] = useState("");
  const [fields, setFields] = useState<Field[]>([]);
  const [newFieldName, setNewFieldName] = useState("");
  const [newFieldType, setNewFieldType] = useState<FieldType>(FieldType.STRING);
  const [error, setError] = useState<string | null>(null);
  const [fieldsError, setFieldsError] = useState<string | null>(null);
  const { request, loading } = useApi();

  const handleCollectionNameChange = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    setCollectionName(e.target.value);
    setError(null);
  };

  const handleNewFieldNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewFieldName(e.target.value);
  };

  const addField = () => {
    if (newFieldName.trim() === "") {
      setFieldsError("Field name cannot be empty.");
      return;
    }

    if (
      fields.find((field) => field.name === newFieldName.trim().toLowerCase())
    ) {
      setFieldsError("Field name must be unique.");
      return;
    }

    setFields([
      ...fields,
      {
        name: newFieldName.trim().toLowerCase(),
        type: newFieldType,
      },
    ]);
    setNewFieldName("");
    setNewFieldType(FieldType.STRING);
  };

  const removeField = (index: number) => {
    setFields(fields.filter((_, i) => i !== index));
  };

  const closeDialog = () => {
    if (dialogRef.current) {
      dialogRef.current.close();
    }
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    if (collectionName.trim() === "") {
      setError("Collection name cannot be empty.");
      return;
    }

    if (fields.length === 0) {
      setFieldsError("You must add at least one field to the collection.");
      return;
    }

    setError(null);
    setFieldsError(null);
    try {
      await request({
        url: "/collections",
        method: "POST",
        data: {
          name: collectionName,
          fields,
          siteId,
        },
      });

      console.log("Collection created successfully");
      closeDialog();
      setTimeout(() => {
        window.location.reload();
      }, 300);
    } catch (e: any) {
      console.error(e);
      setError(e.message || "Failed to create collection. Please try again.");
    } finally {
      setCollectionName("");
      setFields([]);
    }
  };

  return (
    <dialog className="modal" ref={dialogRef}>
      <div className="modal-box">
        <h3 className="font-bold text-lg mb-3">Create new collection</h3>
        <input
          type="text"
          placeholder="Enter collection name"
          className="input input-bordered w-full"
          value={collectionName}
          onChange={handleCollectionNameChange}
        />
        {error && <p className="text-red-500 mt-2">{error}</p>}
        <div className="mt-4">
          <h4 className="font-semibold mb-2">Fields</h4>
          {fields.map((field, index) => (
            <div key={index} className="flex items-center mb-2">
              <p className="mr-2">
                {field.name} ({field.type})
              </p>
              <button
                className="btn btn-sm btn-error btn-outline"
                onClick={() => removeField(index)}
              >
                Remove
              </button>
            </div>
          ))}
          <div className="flex mb-2">
            <input
              type="text"
              placeholder="Field name"
              className="input input-bordered w-full mr-2"
              value={newFieldName}
              onChange={handleNewFieldNameChange}
            />
            <select
              className="select select-bordered"
              value={newFieldType}
              onChange={(e) => setNewFieldType(e.target.value as FieldType)}
            >
              {VALID_FIELD_TYPES.map((type: FieldType) => (
                <option key={type} value={type}>
                  {type.charAt(0).toUpperCase() + type.slice(1)}
                </option>
              ))}
            </select>
            <button className="btn btn-primary ml-2" onClick={addField}>
              Add Field
            </button>
          </div>
        </div>
        {fieldsError && <p className="text-red-500 mt-2">{fieldsError}</p>}
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

export default CreateCollectionModal;
