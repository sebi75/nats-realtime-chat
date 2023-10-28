/* eslint-disable @typescript-eslint/no-floating-promises */
/* eslint-disable @typescript-eslint/no-misused-promises */
"use client";
import { type FunctionComponent } from "react";
import { type z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { signupFormSchema } from "./schemas";
import {
  Form,
  FormControl,
  FormField,
  FormLabel,
  FormMessage,
  FormItem,
} from "@/components/ui/form";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import { usePostSignup } from "./hooks";
import { useRouter } from "next/navigation";
import { CustomLink } from "../common/CustomLink";
import { useToast } from "../ui/use-toast";

export type SignupFormType = z.infer<typeof signupFormSchema>;

export const SignupForm: FunctionComponent = () => {
  const router = useRouter();
  const { toast } = useToast();
  const signupFormMethods = useForm<SignupFormType>({
    defaultValues: {
      email: undefined,
      username: undefined,
      password: "",
    },
    resolver: zodResolver(signupFormSchema),
    mode: "onChange",
  });
  const { mutate: signup } = usePostSignup();
  const onSubmit = (data: SignupFormType) => {
    signup(
      {
        password: data.password,
        email: data.email,
        username: data.username,
      },
      {
        onSuccess: () => {
          toast({
            title: "Signup success",
            description: "Please signin to continue",
            variant: "default",
          });
          router.push("/auth/signin");
        },
      }
    );
  };

  return (
    <Form {...signupFormMethods}>
      <form
        onSubmit={signupFormMethods.handleSubmit(onSubmit)}
        className="flex flex-col items-center justify-center gap-5 space-y-3 rounded-md border p-5"
      >
        <FormField
          control={signupFormMethods.control}
          name="email"
          render={({ field }) => {
            return (
              <FormItem>
                <FormLabel htmlFor="email">Email</FormLabel>
                <FormControl>
                  <Input {...field} id="email" placeholder="Enter email" />
                </FormControl>
                <FormMessage />
              </FormItem>
            );
          }}
        />
        <FormField
          control={signupFormMethods.control}
          name="username"
          render={({ field }) => {
            return (
              <FormItem>
                <FormLabel htmlFor="email">Username</FormLabel>
                <FormControl>
                  <Input
                    {...field}
                    id="username"
                    placeholder="Enter Username"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            );
          }}
        />
        <FormField
          control={signupFormMethods.control}
          name="password"
          render={({ field }) => {
            return (
              <FormItem>
                <FormLabel htmlFor="password">Password</FormLabel>
                <FormControl>
                  <Input
                    type="password"
                    {...field}
                    id="password"
                    placeholder="Enter password"
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            );
          }}
        />
        <CustomLink
          href="/auth/signin"
          className="my-2 underline underline-offset-2"
        >
          Already have an account? Signin
        </CustomLink>
        <Button
          type="submit"
          onClick={signupFormMethods.handleSubmit(onSubmit)}
          variant={"default"}
        >
          Signup
        </Button>
      </form>
    </Form>
  );
};
