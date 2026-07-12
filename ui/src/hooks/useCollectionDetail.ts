import { useState, useEffect } from "react";
import apiClient from "@/lib/api";
import { Collection, Entry, Pagination, Site } from "@/types";

interface CollectionDetailData {
  collection: Collection;
  entries: Entry[];
  site: Site;
  pagination: Pagination;
}

export function useCollectionDetail(
  collectionId: string | undefined,
  siteId: string | undefined
) {
  const [data, setData] = useState<CollectionDetailData | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [page, setPage] = useState(1);

  useEffect(() => {
    if (!collectionId || !siteId) {
      setError("Collection ID and Site ID are required");
      setLoading(false);
      return;
    }

    const fetchCollectionDetail = async () => {
      try {
        setLoading(true);
        setError(null);

        // Fetch collection, entries, and site info in parallel
        const [collectionResponse, entriesResponse, siteResponse] = await Promise.all([
          apiClient.get(`/collections/${collectionId}`),
          apiClient.get(`/collections/${collectionId}/entries`, { params: { page, limit: 20 } }),
          apiClient.get(`/sites/${siteId}`)
        ]);

        setData({
          collection: collectionResponse.data,
          entries: entriesResponse.data.entries,
          pagination: entriesResponse.data.pagination,
          site: siteResponse.data.site
        });
      } catch (err: any) {
        setError(err.response?.data?.error || 'Failed to fetch collection details');
      } finally {
        setLoading(false);
      }
    };

    fetchCollectionDetail();
  }, [collectionId, siteId, page]);

  return {
    collection: data?.collection || null,
    entries: data?.entries || [],
    pagination: data?.pagination || { page: 1, limit: 20, total: 0, totalPages: 0 },
    setPage,
    site: data?.site || null,
    loading,
    error
  };
}
