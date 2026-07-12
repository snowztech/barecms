import { useCallback, useEffect, useRef, useState } from "react";
import { useApi } from "@/hooks/useApi";
import { MediaFile } from "@/types";

interface MediaPickerProps {
  siteId: string;
  value: string;
  onChange: (url: string) => void;
  required?: boolean;
}

const MediaPicker: React.FC<MediaPickerProps> = ({
  siteId,
  value,
  onChange,
  required,
}) => {
  const [files, setFiles] = useState<MediaFile[]>([]);
  const [error, setError] = useState<string | null>(null);
  const inputRef = useRef<HTMLInputElement>(null);
  const { request, loading } = useApi();

  const loadFiles = useCallback(async () => {
    try {
      const response = await request({ url: `/sites/${siteId}/files` });
      setFiles(
        (response.files || []).filter((file: MediaFile) =>
          file.mimeType.startsWith("image/"),
        ),
      );
      setError(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Could not load media");
    }
  }, [request, siteId]);

  useEffect(() => {
    void loadFiles();
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

      {files.length > 0 && (
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
        </div>
      )}

      <input type="hidden" value={value} required={required} readOnly />
    </div>
  );
};

export default MediaPicker;
