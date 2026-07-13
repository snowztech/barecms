import React, { useEffect, useState } from "react";
import { ApiRequestError, useApi } from "@/hooks/useApi";
import { Field, FieldType, VALID_FIELD_TYPES } from "@/types/fields";

interface CreateCollectionModalProps {
  siteId: string;
  dialogRef: React.RefObject<HTMLDialogElement>;
  collectionId?: string;
  initialName?: string;
  initialFields?: Field[];
}

const CreateCollectionModal: React.FC<CreateCollectionModalProps> = ({
  siteId,
  dialogRef,
  collectionId,
  initialName = "",
  initialFields = [],
}) => {
  const [collectionName, setCollectionName] = useState(initialName);
  const [fields, setFields] = useState<Field[]>(initialFields);
  const [newFieldName, setNewFieldName] = useState("");
  const [newFieldType, setNewFieldType] = useState<FieldType>(FieldType.STRING);
  const [newFieldOptional, setNewFieldOptional] = useState<boolean>(false)
  const [error, setError] = useState<string | null>(null);
  const [fieldsError, setFieldsError] = useState<string | null>(null);
  const [serverErrors, setServerErrors] = useState<Record<string, string>>({});
  const { request, loading } = useApi();

  useEffect(() => {
    setCollectionName(initialName);
    setFields(initialFields);
  }, [initialFields, initialName]);

  const handleCollectionNameChange = (
    e: React.ChangeEvent<HTMLInputElement>
  ) => {
    setCollectionName(e.target.value);
    setError(null);
    setServerErrors((current) => ({ ...current, name: "" }));
  };

  const handleNewFieldNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewFieldName(e.target.value);
  };

  const handleNewOptionalChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setNewFieldOptional(e.target.checked)
  }

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
        optional: newFieldOptional
      },
    ]);
    setNewFieldName("");
    setNewFieldType(FieldType.STRING);
    setNewFieldOptional(false);
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
      setServerErrors({ name: "is required" });
      return;
    }

    if (fields.length === 0) {
      setFieldsError("You must add at least one field to the collection.");
      setServerErrors({ fields: "must contain at least one field" });
      return;
    }

    setError(null);
    setFieldsError(null);
    setServerErrors({});
    try {
      await request({
        url: collectionId ? `/collections/${collectionId}` : "/collections",
        method: collectionId ? "PUT" : "POST",
        data: {
          name: collectionName,
          fields,
          siteId,
        },
      });

      closeDialog();
      setTimeout(() => {
        window.location.reload();
      }, 300);
    } catch (e: unknown) {
      if (e instanceof ApiRequestError) setServerErrors(e.fields);
      setError(e instanceof Error ? e.message : "Failed to save collection.");
    } finally {
      if (!collectionId) {
        setCollectionName("");
        setFields([]);
      }
    }
  };

  return (
    <dialog className="modal" ref={dialogRef}>
      <div className="modal-box">
        <h3 className="font-bold text-lg mb-3">{collectionId ? "Edit collection" : "Create new collection"}</h3>
        <input
          type="text"
          placeholder="Enter collection name"
          className={`input input-bordered w-full ${serverErrors.name ? "input-error" : ""}`}
          value={collectionName}
          onChange={handleCollectionNameChange}
          aria-invalid={Boolean(serverErrors.name)}
          aria-describedby={serverErrors.name ? "collection-name-error" : undefined}
        />
        {serverErrors.name && <p id="collection-name-error" role="alert" className="text-sm text-error mt-1">{serverErrors.name}</p>}
        {error && <p className="text-red-500 mt-2">{error}</p>}
        <div className="mt-4">
          <h4 className="font-semibold mb-2">Fields</h4>
          {fields.map((field, index) => (
            <div key={index} className="flex items-center mb-2">
              <p className="mr-2">
                {field.name} ({field.type}{field.optional && ', optional'})
              </p>
              {(serverErrors[`fields.${index}.name`] || serverErrors[`fields.${index}.type`]) && (
                <p role="alert" className="text-sm text-error mr-2">
                  {serverErrors[`fields.${index}.name`] || serverErrors[`fields.${index}.type`]}
                </p>
              )}
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
          <label className="label justify-start gap-2">
            <input
              id="field-optional"
              type="checkbox"
              className="checkbox"
              checked={newFieldOptional}
              onChange={handleNewOptionalChange}
            />
            Optional
          </label>
        </div>
        {fieldsError && <p className="text-red-500 mt-2">{fieldsError}</p>}
        {serverErrors.fields && <p role="alert" className="text-sm text-error mt-2">{serverErrors.fields}</p>}
        <div className="modal-action">
          <button
            disabled={loading}
            onClick={handleSubmit}
            className="btn btn-primary"
          >
            {loading ? "Saving..." : collectionId ? "Save changes" : "Create"}
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
