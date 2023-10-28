import { restService } from "@/utils/fetcher";
import { type Account } from "../account";
import { type User } from "../user";
import { ENDPOINT_AUTH, ENDPOINT_SIGNUP } from "@/types/endpoints";

export type SignupRequest = {
  email: string;
  password: string;
  username: string;
};

export interface SignupResponse extends User {
  account: Account;
}

export const postSignup = async (payload: SignupRequest) => {
  const endpoint = `${ENDPOINT_AUTH}/${ENDPOINT_SIGNUP}`;
  return restService.post<SignupResponse, SignupRequest>(endpoint, payload, {
    authorized: false,
  });
};
