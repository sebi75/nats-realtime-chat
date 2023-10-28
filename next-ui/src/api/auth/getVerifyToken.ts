import { ENDPOINT_AUTH, ENDPOINT_VERIFY } from "@/types/endpoints";
import { restService } from "@/utils/fetcher";
import { type User } from "../user";
import { type Account } from "../account";

export interface VerifyResponse extends User {
  account: Account;
}

export const getVerifyToken = async () => {
  const endpoint = `${ENDPOINT_AUTH}/${ENDPOINT_VERIFY}`;
  return restService.get<VerifyResponse>(endpoint);
};
