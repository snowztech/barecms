import { Field, FieldType } from "@/types/fields";
import { useEffect, useState } from "react";
import { ApiRequestError, useApi } from "@/hooks/useApi";
import MediaPicker from "@/components/MediaPicker";

interface CreateEntryModalProps {
  collectionId: string;
  siteId: string;
  fields: Field[];
  dialogRef: React.RefObject<HTMLDialogElement>;
  entryId?: string;
  initialData?: Record<string, { value: any; type: string }>;
}

const CreateEntryModal: React.FC<CreateEntryModalProps> = ({
  collectionId,
  siteId,
  fields,
  dialogRef,
  entryId,
  initialData,
}) => {
  const [formState, setFormState] = useState<Record<string, any>>({});
  const [error, setError] = useState<string | null>(null);
  const [fieldErrors, setFieldErrors] = useState<Record<string, string>>({});
  const { request, loading } = useApi();

  const closeDialog = () => {
    if (dialogRef.current) {
      dialogRef.current.close();
    }
  };

  useEffect(() => {
    // Initialize form state with empty values for each field
    const initialFormState: Record<string, any> = {};
    fields.forEach((field) => {
      if (field.type === FieldType.BOOLEAN) {
        initialFormState[field.name] = initialData?.[field.name]?.value ?? "false";
      } else {
        initialFormState[field.name] = initialData?.[field.name]?.value ?? "";
      }
    });
    setFormState(initialFormState);
  }, [fields, initialData]);

  const handleInputChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement
    >,
  ) => {
    const { name, value } = e.target;
    setFormState((prevState) => ({ ...prevState, [name]: value }));
    setFieldErrors((current) => ({ ...current, [name]: "" }));
  };

  const handleSubmit = async () => {
    const hasRequiredFields = fields.some(
      (field) => !field.optional && formState[field.name].trim() === "",
    );

    if (hasRequiredFields) {
      setFieldErrors(Object.fromEntries(fields.filter((field) => !field.optional && !formState[field.name]).map((field) => [field.name, "is required"])));
      setError("Please correct the highlighted fields.");
      return;
    }

    setError(null);
    setFieldErrors({});
    try {
      // Prepare data with types
      const dataWithTypes = fields.reduce(
        (acc, field) => {
          let value = formState[field.name];
          if (field.optional && value.trim() === "") value = null;
          acc[field.name] = {
            value,
            type: field.type,
          };
          return acc;
        },
        {} as Record<string, { value: any; type: string }>,
      );

      await request({
        url: entryId ? `/entries/${entryId}` : "/entries",
        method: entryId ? "PUT" : "POST",
        data: entryId ? { data: dataWithTypes } : { collectionId, data: dataWithTypes },
      });

      closeDialog();
      setTimeout(() => {
        window.location.reload();
      }, 300);
    } catch (e: unknown) {
      if (e instanceof ApiRequestError) setFieldErrors(e.fields);
      setError(e instanceof Error ? e.message : "Failed to save entry.");
    }
  };

  const accessibilityProps = (field: Field) => ({
    "aria-invalid": Boolean(fieldErrors[field.name]),
    "aria-describedby": fieldErrors[field.name] ? `${field.name}-error` : undefined,
  });

  const renderFieldInput = (field: Field) => {
    switch (field.type) {
      case FieldType.URL:
        return (
          <input
            type="url"
            id={field.name}
            name={field.name}
            value={formState[field.name]}
            onChange={handleInputChange}
            className="input input-bordered w-full"
            required={!field.optional}
            minLength={field.minLength}
            maxLength={field.maxLength}
            {...accessibilityProps(field)}
          />
        );
      case FieldType.STRING:
        return (
          <input
            type="text"
            id={field.name}
            name={field.name}
            value={formState[field.name]}
            onChange={handleInputChange}
            className="input input-bordered w-full"
            required={!field.optional}
            minLength={field.minLength}
            maxLength={field.maxLength}
            {...accessibilityProps(field)}
          />
        );
      case FieldType.TEXT:
        return (
          <textarea
            id={field.name}
            name={field.name}
            value={formState[field.name]}
            onChange={handleInputChange}
            className="textarea textarea-bordered w-full"
            required={!field.optional}
            minLength={field.minLength}
            maxLength={field.maxLength}
            {...accessibilityProps(field)}
          />
        );
      case FieldType.NUMBER:
        return (
          <input
            type="number"
            id={field.name}
            name={field.name}
            value={formState[field.name]}
            onChange={handleInputChange}
            className="input input-bordered w-full"
            required={!field.optional}
            min={field.min}
            max={field.max}
            {...accessibilityProps(field)}
          />
        );
      case FieldType.BOOLEAN:
        return (
          <select
            id={field.name}
            name={field.name}
            value={formState[field.name]}
            onChange={handleInputChange}
            className="select select-bordered w-full"
            required={!field.optional}
            {...accessibilityProps(field)}
          >
            <option disabled>
              Select an option
            </option>
            <option value="true">True</option>
            <option value="false">False</option>
          </select>
        );
      case FieldType.DATE:
        return (
          <input
            type="date"
            id={field.name}
            name={field.name}
            value={formState[field.name]}
            onChange={handleInputChange}
            className="input input-bordered w-full"
            required={!field.optional}
            {...accessibilityProps(field)}
          />
        );
      case FieldType.IMAGE:
        return (
          <MediaPicker
            siteId={siteId}
            value={formState[field.name]}
            onChange={(url) => {
              setFormState((previous) => ({ ...previous, [field.name]: url }));
              setFieldErrors((current) => ({ ...current, [field.name]: "" }));
            }}
            required={!field.optional}
            invalid={Boolean(fieldErrors[field.name])}
            ariaDescribedBy={fieldErrors[field.name] ? `${field.name}-error` : undefined}
          />
        );
      default:
        return (
          <input
            type="text"
            id={field.name}
            name={field.name}
            value={formState[field.name]}
            onChange={handleInputChange}
            className="input input-bordered w-full"
            required={!field.optional}
            {...accessibilityProps(field)}
          />
        );
    }
  };

  return (
    <dialog className="modal" ref={dialogRef}>
      <div className="modal-box">
        <h3 className="font-bold text-lg mb-3">{entryId ? "Edit entry" : "Create new entry"}</h3>
        {error && <p className="text-red-500 mb-4">{error}</p>}
        <div className="mt-4">
          {fields.map((field) => (
            <div
              key={field.name}
              className={`mb-4 rounded ${fieldErrors[field.name] ? "ring-1 ring-error p-2" : ""}`}
            >
              <label
                className="block text-sm font-medium text-gray-700 mb-2"
                htmlFor={field.name}
              >
                {field.name} ({field.type}{field.optional && ', optional'})
              </label>
              {renderFieldInput(field)}
              {fieldErrors[field.name] && (
                <p id={`${field.name}-error`} role="alert" className="text-sm text-error mt-1">
                  {fieldErrors[field.name]}
                </p>
              )}
            </div>
          ))}
        </div>
        <div className="modal-action">
          <button
            disabled={loading}
            onClick={handleSubmit}
            className="btn btn-primary"
          >
            {loading ? "Saving..." : entryId ? "Save changes" : "Create"}
          </button>
          <button className="btn" onClick={closeDialog}>
            Cancel
          </button>
        </div>
      </div>
    </dialog>
  );
};
export default CreateEntryModal;
