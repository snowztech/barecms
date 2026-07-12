export interface User {
  id: string;
  email: string;
  username: string;
  isActive?: boolean;
  createdAt?: string;
}

export interface SiteUser {
  id: string;
  user: User;
  role: 'owner' | 'editor' | 'viewer';
  joinedAt?: string;
  status: 'joined' | 'pending';
}

export interface InviteUserRequest {
  email: string;
  role: 'editor' | 'viewer';
}

export interface InviteResponse {
  id: string;
  email: string;
  role: string;
  status: string;
  invitedAt: string;
}

export const AUTH_TOKEN_KEY = "auth_token";

export const ROLES = {
  OWNER: 'owner' as const,
  EDITOR: 'editor' as const,
  VIEWER: 'viewer' as const,
};

export type Role = typeof ROLES[keyof typeof ROLES];
