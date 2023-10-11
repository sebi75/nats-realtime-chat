import { ENDPOINT_AUTH, ENDPOINT_LOGIN } from "@/types/endpoints";
import { restService } from "@/utils/fetcher";

export type SigninRequest = {
  email?: string;
  password: string;
  username?: string;
};

export interface SigninResponse {
  token?: string;
}

export const postSignin = async (payload: SigninRequest) => {
  const endpoint = `${ENDPOINT_AUTH}/${ENDPOINT_LOGIN}`;
  return restService.post<SigninResponse, SigninRequest>(endpoint, payload);
};
