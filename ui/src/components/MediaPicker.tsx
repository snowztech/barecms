import { useCallback, useEffect, useRef, useState } from "react";
import { useApi } from "@/hooks/useApi";
import { MediaFile } from "@/types";

interface MediaPickerProps {
  siteId: string;
  value: string;
  onChange: (url: string) => void;
  required?: boolean;
  invalid?: boolean;
  ariaDescribedBy?: string;
}

const MediaPicker: React.FC<MediaPickerProps> = ({
  siteId,
  value,
  onChange,
  required,
  invalid,
  ariaDescribedBy,
}) => {
  const [files, setFiles] = useState<MediaFile[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0);
  const inputRef = useRef<HTMLInputElement>(null);
  const { request, loading } = useApi();

  const loadFiles = useCallback(async (requestedPage = 1) => {
    try {
      const response = await request({ url: `/sites/${siteId}/files`, params: { page: requestedPage, limit: 50 } });
      const images = (response.files || []).filter((file: MediaFile) =>
          file.mimeType.startsWith("image/"),
        );
      setFiles((current) => requestedPage === 1 ? images : [...current, ...images]);
      setPage(requestedPage);
      setTotalPages(response.pagination?.totalPages || 0);
      setError(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Could not load media");
    }
  }, [request, siteId]);

  useEffect(() => {
    void loadFiles(1);
  }, [loadFiles]);

  const upload = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const selected = event.target.files?.[0];
    if (!selected) return;

    const form = new FormData();
    form.append("file", selected);
    try {
      const uploaded: MediaFile = await request({
        url: `/sites/${siteId}/files`,
        method: "POST",
        data: form,
        headers: { "Content-Type": "multipart/form-data" },
      });
      setFiles((current) => [uploaded, ...current]);
      onChange(uploaded.url);
      setError(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Upload failed");
    } finally {
      event.target.value = "";
    }
  };

  return (
    <div className="space-y-3">
      <input
        ref={inputRef}
        type="file"
        accept="image/jpeg,image/png,image/gif,image/webp"
        onChange={upload}
        className="hidden"
      />
      <div className="flex flex-wrap gap-2">
        <button
          type="button"
          className="btn btn-sm btn-primary"
          disabled={loading}
          onClick={() => inputRef.current?.click()}
        >
          {loading ? "Working..." : "Upload image"}
        </button>
        {value && (
          <button type="button" className="btn btn-sm" onClick={() => onChange("")}>
            Clear
          </button>
        )}
      </div>

      {error && <p className="text-sm text-error">{error}</p>}
      {value && (
        <div className="flex items-center gap-3 rounded-bare border border-base-300 p-2">
          <img src={value} alt="Selected media" className="h-16 w-16 rounded object-cover" />
          <span className="min-w-0 truncate text-sm">Selected image</span>
        </div>
      )}

      {(files.length > 0 || page < totalPages) && (
        <div>
          <p className="mb-2 text-sm text-bare-600">Or choose an existing image</p>
          <div className="grid max-h-48 grid-cols-3 gap-2 overflow-y-auto">
            {files.map((file) => (
              <button
                type="button"
                key={file.id}
                title={file.originalName}
                aria-label={`Select ${file.originalName}`}
                onClick={() => onChange(file.url)}
                className={`overflow-hidden rounded border-2 ${value === file.url ? "border-primary" : "border-transparent"}`}
              >
                <img src={file.url} alt={file.originalName} className="aspect-square w-full object-cover" />
              </button>
            ))}
          </div>
          {page < totalPages && (
            <button type="button" className="btn btn-sm mt-2 w-full" disabled={loading} onClick={() => void loadFiles(page + 1)}>
              Load more
            </button>
          )}
        </div>
      )}

      <input
        type="hidden"
        value={value}
        required={required}
        readOnly
        aria-invalid={invalid}
        aria-describedby={ariaDescribedBy}
      />
    </div>
  );
};

export default MediaPicker;
