import { restService } from "@/utils/fetcher";

export interface VerifyTokenResponse {
  userId: string;
  accountId: string;
}

export const getVerifyToken = async (endpoint: string) => {
  return restService.get<VerifyTokenResponse>(endpoint);
};
