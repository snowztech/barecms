import { Field, FieldType } from "@/types/fields";
import { useEffect, useState } from "react";
import { useApi } from "@/hooks/useApi";

interface CreateEntryModalProps {
  collectionId: string;
  fields: Field[];
  dialogRef: React.RefObject<HTMLDialogElement>;
}

const CreateEntryModal: React.FC<CreateEntryModalProps> = ({
  collectionId,
  fields,
  dialogRef,
}) => {
  const [formState, setFormState] = useState<Record<string, any>>({});
  const [error, setError] = useState<string | null>(null);
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
        initialFormState[field.name] = "false";
      } else {
        initialFormState[field.name] = "";
      }
    });
    setFormState(initialFormState);
  }, [fields]);

  const handleInputChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement
    >,
  ) => {
    const { name, value } = e.target;
    setFormState((prevState) => ({ ...prevState, [name]: value }));
  };

  const handleSubmit = async () => {
    console.log("fields", fields);

    const hasEmptyFields = fields.some(
      (field) => formState[field.name].trim() === "",
    );

    console.log("hasEmptyFields", hasEmptyFields);

    if (hasEmptyFields) {
      setError("All fields are required.");
      return;
    }

    setError(null);
    try {
      // Prepare data with types
      const dataWithTypes = fields.reduce(
        (acc, field) => {
          acc[field.name] = {
            value: formState[field.name],
            type: field.type,
          };
          return acc;
        },
        {} as Record<string, { value: any; type: string }>,
      );

      await request({
        url: "/entries",
        method: "POST",
        data: { collectionId, data: dataWithTypes },
      });

      console.log("Entry created successfully");
      closeDialog();
      setTimeout(() => {
        window.location.reload();
      }, 300);
    } catch (e: any) {
      console.error(e);
      setError(e.message || "Failed to create entry. Please try again.");
    }
  };

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
            required
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
            required
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
            required
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
            required
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
            required
          >
            <option disabled selected>
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
            required
          />
        );
      case FieldType.IMAGE:
        return (
          <input
            type="url"
            id={field.name}
            name={field.name}
            value={formState[field.name]}
            onChange={handleInputChange}
            className="input input-bordered w-full"
            required
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
            required
          />
        );
    }
  };

  return (
    <dialog className="modal" ref={dialogRef}>
      <div className="modal-box">
        <h3 className="font-bold text-lg mb-3">Create new entry</h3>
        {error && <p className="text-red-500 mb-4">{error}</p>}
        <div className="mt-4">
          {fields.map((field) => (
            <div key={field.name} className="mb-4">
              <label
                className="block text-sm font-medium text-gray-700 mb-2"
                htmlFor={field.name}
              >
                {field.name} ({field.type})
              </label>
              {renderFieldInput(field)}
            </div>
          ))}
        </div>
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
export default CreateEntryModal;
