import { ENDPOINT_AUTH, ENDPOINT_VERIFY } from "@/types/endpoints";
import { restService } from "@/utils/fetcher";

export interface VerifyTokenResponse {
  userId: string;
  accountId: string;
}

export const getVerifyToken = async () => {
  const endpoint = `${ENDPOINT_AUTH}/${ENDPOINT_VERIFY}`;
  return restService.get<VerifyTokenResponse>(endpoint);
};
