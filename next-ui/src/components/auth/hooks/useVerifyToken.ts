import { useQuery } from "@tanstack/react-query";
import { ENDPOINT_AUTH, ENDPOINT_VERIFY } from "@/types/endpoints";
import { getVerifyToken } from "@/api/auth";

export const useVerifyToken = () => {
  return useQuery([ENDPOINT_AUTH, ENDPOINT_VERIFY], () => getVerifyToken(), {
    cacheTime: 0,
  });
};
