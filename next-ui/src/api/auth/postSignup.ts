import { restService } from "@/utils/fetcher";
import { type Account } from "../account";
import { type User } from "../user";

export type SignupRequest = {
  email: string;
  password: string;
  username: string;
};

export interface SignupResponse extends User {
  account: Account;
}

export const postSignup = async (endpoint: string, payload: SignupRequest) => {
  return restService.post<SignupResponse, SignupRequest>(endpoint, payload);
};
