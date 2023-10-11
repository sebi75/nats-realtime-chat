import { type SignupRequest, postSignup } from "@/api/auth/postSignup";
import { useToast } from "@/components/ui/use-toast";
import { useMutation } from "@tanstack/react-query";

export const usePostSignup = (params: SignupRequest) => {
  const toast = useToast();
  return useMutation(() => postSignup(params), {
    onError: (error: string | undefined) => {
      toast.toast({
        title: "Error",
        description: error ?? "Something went wrong",
        variant: "destructive",
      });
    },
  });
};
