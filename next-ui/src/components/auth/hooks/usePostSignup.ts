import {
  type SignupRequest,
  postSignup,
  type SignupResponse,
} from "@/api/auth/postSignup";
import { useToast } from "@/components/ui/use-toast";
import { useMutation } from "@tanstack/react-query";

export const usePostSignup = () => {
  const toast = useToast();
  return useMutation<SignupResponse, Error, SignupRequest>(
    (payload) => postSignup(payload),
    {
      onError: (error: Error) => {
        toast.toast({
          title: "Error",
          description: error.message ?? "Something went wrong",
          variant: "destructive",
        });
      },
    }
  );
};
