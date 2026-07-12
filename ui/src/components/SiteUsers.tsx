import React, { useState } from 'react';
import { SiteUser, InviteUserRequest, ROLES, Role } from '@/types/auth';
import { useSiteUsers } from '@/hooks/useSiteUsers';

interface SiteUsersProps {
  siteId: string;
  currentUserRole: Role;
}

const SiteUsers: React.FC<SiteUsersProps> = ({ siteId, currentUserRole }) => {
  const { siteUsers, loading, error, inviteUser, removeUser, updateUserRole } = useSiteUsers(siteId);
  const [showInviteForm, setShowInviteForm] = useState(false);
  const [inviteForm, setInviteForm] = useState<InviteUserRequest>({
    email: '',
    role: 'editor',
  });
  const [submitting, setSubmitting] = useState(false);

  const canManageUsers = currentUserRole === ROLES.OWNER;

  const handleInviteUser = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!canManageUsers) return;

    setSubmitting(true);
    try {
      await inviteUser(inviteForm);
      setInviteForm({ email: '', role: 'editor' });
      setShowInviteForm(false);
    } catch (err: any) {
      alert(`Failed to invite user: ${err.message}`);
    } finally {
      setSubmitting(false);
    }
  };

  const handleRemoveUser = async (userId: string, username: string) => {
    if (!canManageUsers) return;

    if (confirm(`Are you sure you want to remove ${username} from this site?`)) {
      try {
        await removeUser(userId);
      } catch (err: any) {
        alert(`Failed to remove user: ${err.message}`);
      }
    }
  };

  const handleRoleChange = async (userId: string, newRole: string) => {
    if (!canManageUsers) return;

    try {
      await updateUserRole(userId, newRole);
    } catch (err: any) {
      alert(`Failed to update role: ${err.message}`);
    }
  };

  const getRoleBadgeColor = (role: string) => {
    switch (role) {
      case 'owner': return 'bg-purple-100 text-purple-800';
      case 'editor': return 'bg-blue-100 text-blue-800';
      case 'viewer': return 'bg-gray-100 text-gray-800';
      default: return 'bg-gray-100 text-gray-800';
    }
  };

  if (loading) return <div className="p-4">Loading users...</div>;
  if (error) return <div className="p-4 text-red-600">Error: {error}</div>;

  return (
    <div className="space-y-6">
      <div className="flex justify-between items-center">
        <h3 className="text-lg font-semibold">Site Members</h3>
        {canManageUsers && (
          <button
            onClick={() => setShowInviteForm(true)}
            className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
          >
            Invite User
          </button>
        )}
      </div>

      {/* Invite Form */}
      {showInviteForm && canManageUsers && (
        <div className="bg-gray-50 p-4 rounded-lg">
          <h4 className="font-medium mb-3">Invite New User</h4>
          <form onSubmit={handleInviteUser} className="space-y-3">
            <div>
              <label className="block text-sm font-medium mb-1">Email</label>
              <input
                type="email"
                value={inviteForm.email}
                onChange={(e) => setInviteForm({ ...inviteForm, email: e.target.value })}
                className="w-full px-3 py-2 border rounded"
                required
              />
            </div>
            <div>
              <label className="block text-sm font-medium mb-1">Role</label>
              <select
                value={inviteForm.role}
                onChange={(e) => setInviteForm({ ...inviteForm, role: e.target.value as 'editor' | 'viewer' })}
                className="w-full px-3 py-2 border rounded"
              >
                <option value="editor">Editor</option>
                <option value="viewer">Viewer</option>
              </select>
            </div>
            <div className="flex gap-2">
              <button
                type="submit"
                disabled={submitting}
                className="px-4 py-2 bg-green-600 text-white rounded hover:bg-green-700 disabled:opacity-50"
              >
                {submitting ? 'Inviting...' : 'Invite'}
              </button>
              <button
                type="button"
                onClick={() => setShowInviteForm(false)}
                className="px-4 py-2 bg-gray-300 text-gray-700 rounded hover:bg-gray-400"
              >
                Cancel
              </button>
            </div>
          </form>
        </div>
      )}

      {/* Users List */}
      <div className="bg-white border rounded-lg">
        <div className="p-4 border-b">
          <h4 className="font-medium">Current Members ({siteUsers.length})</h4>
        </div>
        <div className="divide-y">
          {siteUsers.map((siteUser) => (
            <div key={siteUser.id} className="p-4 flex items-center justify-between">
              <div className="flex items-center gap-3">
                <div className="w-8 h-8 bg-gray-200 rounded-full flex items-center justify-center">
                  {siteUser.user.username.charAt(0).toUpperCase()}
                </div>
                <div>
                  <div className="font-medium">{siteUser.user.username}</div>
                  <div className="text-sm text-gray-600">{siteUser.user.email}</div>
                </div>
                <span className={`px-2 py-1 rounded text-xs font-medium ${getRoleBadgeColor(siteUser.role)}`}>
                  {siteUser.role}
                </span>
                {siteUser.status === 'pending' && (
                  <span className="px-2 py-1 rounded text-xs font-medium bg-yellow-100 text-yellow-800">
                    Pending
                  </span>
                )}
              </div>

              {canManageUsers && siteUser.role !== 'owner' && (
                <div className="flex items-center gap-2">
                  <select
                    value={siteUser.role}
                    onChange={(e) => handleRoleChange(siteUser.id, e.target.value)}
                    className="px-2 py-1 text-sm border rounded"
                  >
                    <option value="editor">Editor</option>
                    <option value="viewer">Viewer</option>
                  </select>
                  <button
                    onClick={() => handleRemoveUser(siteUser.id, siteUser.user.username)}
                    className="px-3 py-1 text-sm text-red-600 hover:bg-red-50 rounded"
                  >
                    Remove
                  </button>
                </div>
              )}
            </div>
          ))}
        </div>
      </div>

      {/* Role Explanations */}
      <div className="bg-gray-50 p-4 rounded-lg">
        <h4 className="font-medium mb-2">Role Permissions</h4>
        <div className="text-sm space-y-1">
          <div><strong>Owner:</strong> Full control, can invite/remove users, manage site</div>
          <div><strong>Editor:</strong> Can create, edit, and delete content</div>
          <div><strong>Viewer:</strong> Read-only access to content</div>
        </div>
      </div>
    </div>
  );
};

export default SiteUsers;