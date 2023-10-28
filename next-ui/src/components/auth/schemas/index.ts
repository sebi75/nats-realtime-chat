import { z } from "zod";

export const signinFormSchema = z.object({
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
});
// .refine((data) => {
//   const err =
//     (data.email && !data.username) || (!data.email && data.username);
//   err
//     ? {
//         message: "Either email or username must be provided, not both",
//       }
//     : true;
// });

export const signupFormSchema = z.object({
  email: z.string().email({
    message: "Please enter a valid email address",
  }),
  username: z
    .string()
    .min(3, {
      message: "Username must be at least 3 characters long",
    })
    .max(20, {
      message: "Username must be at most 20 characters long",
    }),
  password: z
    .string()
    .min(6, {
      message: "Password must be at least 6 characters long",
    })
    .max(100, {
      message: "Password must be at most 100 characters long",
    }),
});
