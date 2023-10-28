import { env } from "@/env.mjs";

export type RequestOptions = RequestInit & {
  authorized: boolean;
};

export type ApiErrorResponse = {
  message: string;
  errors?: Record<string, string>;
};

export const fetcher = async <T>(
  endpoint: string,
  options: RequestOptions
): Promise<T> => {
  const { authorized, ...rest } = options;
  const url = `${env.NEXT_PUBLIC_API_URL}/${endpoint}`;

  try {
    const headers = new Headers();
    headers.set("Content-Type", "application/json");
    if (authorized) {
      const token = localStorage.getItem("token");
      if (!token) {
        throw new Error("Unauthorized");
      }
      headers.set("Authorization", `Bearer ${token}`);
    }

    for (const [key, value] of Object.entries(rest?.headers || {})) {
      headers.set(key, (value as string) ?? "");
    }
    console.log({ rest, headers });
    const result = await fetch(url, {
      ...rest,
      headers,
    });

    if (!result.ok) {
      const json = (await result.json()) as ApiErrorResponse;
      throw new Error(json.message);
    } else {
      return (await result.json()) as T;
    }
  } catch (error) {
    console.error(error);
    throw new Error("Something went wrong!");
  }
};

export const restService = {
  get: <T>(
    endpoint: string,
    options: RequestOptions = { method: "GET", authorized: true }
  ) => {
    return fetcher<T>(endpoint, options);
  },
  post: <T, U>(
    endpoint: string,
    payload: U,
    options: RequestOptions = { method: "POST", authorized: true }
  ) => {
    return fetcher<T>(endpoint, {
      body: JSON.stringify(payload),
      method: "POST",
      ...options,
    });
  },
  put: <T, U>(
    endpoint: string,
    payload: U,
    options: RequestOptions = { method: "PUT", authorized: true }
  ) => {
    return fetcher<T>(endpoint, { ...options, body: JSON.stringify(payload) });
  },
};
