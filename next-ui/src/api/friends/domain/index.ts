import { type User } from "@/api/user";

export interface Friend {
  id: string;
  requester_id: string;
  addressee_id: string;
  created_at: string;
  status: string;
}

export interface FriendWithUser extends Friend {
  user: User;
}
