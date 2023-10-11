import { restService } from "@/utils/fetcher";

export type SigninRequest = {
  email?: string;
  password: string;
  username?: string;
};

export interface SigninResponse {
  token?: string;
}

export const postSignin = async (endpoint: string, payload: SigninRequest) => {
  return restService.post<SigninResponse, SigninRequest>(endpoint, payload);
};
