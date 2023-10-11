import { type FunctionComponent } from "react";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { FormState, useForm } from "react-hook-form";

const signinFormSchema = z
  .object({
    email: z.string().email().optional(),
    username: z
      .string()
      .min(3, {
        message: "Username must be at least 3 characters long",
      })
      .max(20, {
        message: "Username must be at most 20 characters long",
      })
      .optional(),
    password: z
      .string()
      .min(6, {
        message: "Password must be at least 6 characters long",
      })
      .max(100, {
        message: "Password must be at most 100 characters long",
      }),
  })
  .refine((data) => {
    // We expect either email or username to be present, not both
    // not none
    const err =
      (data.email && !data.username) || (!data.email && data.username);
    err
      ? {
          message: "Either email or username must be provided, not both",
        }
      : true;
  });

export type SigninFormType = z.infer<typeof signinFormSchema>;

export const SigninForm: FunctionComponent = () => {
  const {} = useForm<SigninFormType>({
    defaultValues: {
      email: undefined,
      username: undefined,
      password: "",
    },
    mode: "onChange",
  });
  const onSubmit = (data: SigninFormType) => {
    console.log(data);
  };
  return <div>signin form page</div>;
};
