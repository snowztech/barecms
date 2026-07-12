import { useState, useEffect } from 'react';
import { SiteUser, InviteUserRequest, InviteResponse } from '@/types/auth';
import { useApi } from '@/hooks/useApi';

export const useSiteUsers = (siteId: string) => {
  const [siteUsers, setSiteUsers] = useState<SiteUser[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const { apiCall } = useApi();

  const fetchSiteUsers = async () => {
    if (!siteId) return;

    setLoading(true);
    setError(null);

    try {
      const response = await apiCall(`/sites/${siteId}/users`);
      setSiteUsers(response.users || []);
    } catch (err: any) {
      setError(err.message || 'Failed to fetch site users');
    } finally {
      setLoading(false);
    }
  };

  const inviteUser = async (inviteData: InviteUserRequest): Promise<InviteResponse> => {
    try {
      const response = await apiCall(`/sites/${siteId}/users/invite`, {
        method: 'POST',
        body: JSON.stringify(inviteData),
      });

      // Refresh the users list
      await fetchSiteUsers();

      return response.invitation;
    } catch (err: any) {
      throw new Error(err.message || 'Failed to invite user');
    }
  };

  const removeUser = async (userId: string): Promise<void> => {
    try {
      await apiCall(`/sites/${siteId}/users/${userId}`, {
        method: 'DELETE',
      });

      // Refresh the users list
      await fetchSiteUsers();
    } catch (err: any) {
      throw new Error(err.message || 'Failed to remove user');
    }
  };

  const updateUserRole = async (userId: string, role: string): Promise<void> => {
    try {
      await apiCall(`/sites/${siteId}/users/${userId}/role`, {
        method: 'PUT',
        body: JSON.stringify({ role }),
      });

      // Refresh the users list
      await fetchSiteUsers();
    } catch (err: any) {
      throw new Error(err.message || 'Failed to update user role');
    }
  };

  useEffect(() => {
    fetchSiteUsers();
  }, [siteId]);

  return {
    siteUsers,
    loading,
    error,
    inviteUser,
    removeUser,
    updateUserRole,
    refetch: fetchSiteUsers,
  };
};