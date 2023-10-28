import {
  postSignin,
  type SigninRequest,
  type SigninResponse,
} from "@/api/auth/postSignin";
import { useToast } from "@/components/ui/use-toast";
import { useMutation } from "@tanstack/react-query";

export const usePostSignin = () => {
  const toast = useToast();
  return useMutation<SigninResponse, Error, SigninRequest>(
    (payload) => postSignin(payload),
    {
      onError: (error) => {
        console.log(error);
        toast.toast({
          title: "Error",
          description: error.message ?? "Something went wrong",
          variant: "destructive",
        });
      },
    }
  );
};
