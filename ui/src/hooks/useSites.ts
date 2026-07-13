import { useState, useEffect } from "react";
import apiClient from "@/lib/api";
import { Pagination, Site } from "@/types";
import { apiErrorMessage } from "@/hooks/useApi";

interface SitesResponse {
  sites: Site[];
  pagination: Pagination;
}

export function useSites() {
  const [sites, setSites] = useState<Site[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [page, setPage] = useState(1);
  const [pagination, setPagination] = useState<Pagination>({ page: 1, limit: 20, total: 0, totalPages: 0 });

  useEffect(() => {
    const fetchSites = async () => {
      try {
        setLoading(true);
        setError(null);
        const response = await apiClient.get<SitesResponse>('/sites', { params: { page, limit: 20 } });
        setSites(response.data.sites);
        setPagination(response.data.pagination);
      } catch (err: any) {
        setError(apiErrorMessage(err, "Failed to fetch sites"));
      } finally {
        setLoading(false);
      }
    };

    fetchSites();
  }, [page]);

  return { sites, pagination, setPage, loading, error };
}
