import useDelete from "@/hooks/useDelete";
import { Trash2 } from "lucide-react";
import Loader from "@/components/Loader";
import React from "react";
import { FieldType } from "@/types/fields";

interface EntryData {
  value: any;
  type: FieldType;
}

interface EntryCardProps {
  siteId: string;
  collectionId: string;
  entryId: string;
  data: Record<string, EntryData>;
}

const EntryCard: React.FC<EntryCardProps> = ({
  collectionId,
  siteId,
  entryId,
  data,
}) => {
  const { isDeleting, error, handleDelete } = useDelete(
    `/entries/${entryId}`,
    `/sites/${siteId}/collections/${collectionId}`
  );

  if (isDeleting) {
    return (
      <div className="card-bare p-6 flex items-center justify-center min-h-[120px]">
        <Loader size="md" variant="minimal" />
      </div>
    );
  }

  return (
    <div className="card-bare p-6 hover-lift group relative">
      {error && (
        <div className="alert-bare alert-bare-error p-2 text-xs mb-4">
          {error}
        </div>
      )}

      <div className="space-y-3 mb-4">
        {Object.entries(data).map(([key, { value, type }]) => (
          <div key={key} className="space-y-1">
            <label className="block text-xs font-medium text-bare-600 uppercase tracking-wide">
              {key}
            </label>
            {type === FieldType.STRING && (
              <p className="text-sm text-base-content break-words">{value}</p>
            )}
            {type === FieldType.TEXT && (
              <div className="text-sm text-base-content break-words">
                {value && value.length > 150 ? (
                  <details className="cursor-pointer">
                    <summary className="hover:text-primary transition-colors">
                      {value.substring(0, 150)}...
                    </summary>
                    <div className="mt-2 text-sm">{value}</div>
                  </details>
                ) : (
                  <p className="text-sm text-base-content break-words">
                    {value}
                  </p>
                )}
              </div>
            )}
            {type === FieldType.IMAGE && (
              <img
                src={value}
                alt={value}
                className="w-full h-24 object-contain rounded border border-bare-200"
              />
            )}
            {type === FieldType.DATE && (
              <p className="text-mono text-sm text-base-content">
                {new Date(value).toLocaleDateString()}
              </p>
            )}
            {type === FieldType.NUMBER && (
              <p className="text-mono text-sm text-base-content">{value}</p>
            )}
            {type === FieldType.BOOLEAN && (
              <span
                className={`badge-bare ${value === "true" || value === true ? "text-success bg-teal-500/10" : "text-error bg-rose-500/10"}`}
              >
                {value === "true" || value === true ? "Yes" : "No"}
              </span>
            )}
            {type === FieldType.URL && (
              <a
                className="text-sm text-primary hover:text-primary-focus transition-colors break-all"
                href={value}
                target="_blank"
                rel="noopener noreferrer"
              >
                {value}
              </a>
            )}
          </div>
        ))}
      </div>

      <div className="flex justify-end border-t border-bare-200 pt-3 mt-3">
        <button
          onClick={handleDelete}
          className="p-2 text-bare-400 hover:text-error transition-colors rounded"
          aria-label="Delete entry"
        >
          <Trash2 size={16} />
        </button>
      </div>
    </div>
  );
};

export default EntryCard;
