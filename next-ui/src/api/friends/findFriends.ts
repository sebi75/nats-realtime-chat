import { FRIENDS } from "@/types/endpoints";
import { type Query } from "@/types/query";
import { restService } from "@/utils/fetcher";
import { type FriendWithUser } from "./domain";

export type FindFriendsResponse = Query<FriendWithUser[]>;

export const findFriends = async () => {
  const endpoint = `${FRIENDS}`;
  return restService.get(endpoint);
};
