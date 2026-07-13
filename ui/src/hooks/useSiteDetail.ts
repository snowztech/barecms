import { useState, useEffect } from "react";
import apiClient from "@/lib/api";
import { Site, Collection, Pagination } from "@/types";
import { apiErrorMessage } from "@/hooks/useApi";

interface SiteDetailData {
  site: Site;
  collections: Collection[];
  pagination: Pagination;
}

export function useSiteDetail(siteId: string | undefined) {
  const [data, setData] = useState<SiteDetailData | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [page, setPage] = useState(1);

  useEffect(() => {
    if (!siteId) {
      setError("Site ID is required");
      setLoading(false);
      return;
    }

    const fetchSiteDetail = async () => {
      try {
        setLoading(true);
        setError(null);

        // Fetch both site and collections in parallel
        const [siteResponse, collectionsResponse] = await Promise.all([
          apiClient.get(`/sites/${siteId}`),
          apiClient.get(`/sites/${siteId}/collections`, { params: { page, limit: 20 } })
        ]);

        setData({
          site: siteResponse.data.site,
          collections: collectionsResponse.data.collections,
          pagination: collectionsResponse.data.pagination,
        });
      } catch (err: any) {
        setError(apiErrorMessage(err, "Failed to fetch site details"));
      } finally {
        setLoading(false);
      }
    };

    fetchSiteDetail();
  }, [siteId, page]);

  return {
    site: data?.site || null,
    collections: data?.collections || [],
    pagination: data?.pagination || { page: 1, limit: 20, total: 0, totalPages: 0 },
    setPage,
    loading,
    error
  };
}
