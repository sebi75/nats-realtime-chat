/* eslint-disable @typescript-eslint/no-floating-promises */
/* eslint-disable @typescript-eslint/no-misused-promises */
"use client";
import { type FunctionComponent } from "react";
import { type z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { signinFormSchema } from "./schemas";
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
import { usePostSignin } from "./hooks";
import { useRouter } from "next/navigation";
import { useQueryClient } from "@tanstack/react-query";
import { ENDPOINT_AUTH, ENDPOINT_VERIFY } from "@/types/endpoints";

export type SigninFormType = z.infer<typeof signinFormSchema>;

export const SigninForm: FunctionComponent = () => {
  const router = useRouter();
  const queryClient = useQueryClient();
  const signinFormMethods = useForm<SigninFormType>({
    defaultValues: {
      email: undefined,
      username: undefined,
      password: "",
    },
    resolver: zodResolver(signinFormSchema),
    mode: "onChange",
  });
  const { mutate } = usePostSignin();
  const onSubmit = (data: SigninFormType) => {
    mutate(
      {
        password: data.password,
        email: data.email,
        username: data.username,
      },
      {
        onSuccess: (data) => {
          const { token } = data;
          if (token) {
            localStorage.setItem("token", token);
            queryClient.invalidateQueries({
              queryKey: [ENDPOINT_AUTH, ENDPOINT_VERIFY],
            });
          }
          router.push("/explore");
        },
      }
    );
  };

  return (
    <Form {...signinFormMethods}>
      <form
        onSubmit={signinFormMethods.handleSubmit(onSubmit)}
        className="flex flex-col items-center justify-center gap-5 space-y-3 rounded-md border p-5"
      >
        <FormField
          control={signinFormMethods.control}
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
          control={signinFormMethods.control}
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

        <Button
          type="submit"
          onClick={signinFormMethods.handleSubmit(onSubmit)}
          variant={"default"}
        >
          Signin
        </Button>
      </form>
    </Form>
  );
};
