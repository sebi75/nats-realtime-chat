import { postSignin, type SigninRequest } from "@/api/auth/postSignin";
import { useToast } from "@/components/ui/use-toast";
import { useMutation } from "@tanstack/react-query";

export const usePostSignin = (params: SigninRequest) => {
  const toast = useToast();
  return useMutation(() => postSignin(params), {
    onError: (error: string | undefined) => {
      toast.toast({
        title: "Error",
        description: error ?? "Something went wrong",
        variant: "destructive",
      });
    },
  });
};
