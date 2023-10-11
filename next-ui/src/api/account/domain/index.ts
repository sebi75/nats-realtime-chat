export interface Account {
  id: string;
  createdAt: Date;
  updatedAt: Date | null;
  email: string;
  lastLogin: Date | null;
  emailVerified: boolean;
}
