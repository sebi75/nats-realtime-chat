export enum HTTPMethod {
  GET = "GET",
  POST = "POST",
  PUT = "PUT",
  DELETE = "DELETE",
}

export const fetcher = async <T>(
  path: string,
  method: HTTPMethod,
  authorized = false,
  params?: Record<string, string>
): Promise<T> => {
  const baseURL = getServerURL();
  const url = `${baseURL}/${path}`;

  const token = localStorage.getItem("token");

  const headers = new Headers();
  headers.append("Content-Type", "application/json");
  headers.append("Accept", "application/json");
  if (token && authorized) {
    headers.append("Authorization", `Bearer ${token}`);
  }

  try {
    const response = await fetch(url, {
      headers: headers,
      method: method,
      body: JSON.stringify(params),
    });

    if (!response.ok) {
      throw new Error(response.statusText);
    }

    const data = (await response.json()) as T;
    return data;
  } catch (error) {
    if (error instanceof Error) {
      throw new Error(error.message);
    }
    throw new Error("Something went wrong");
  }
};

export const getServerURL = () => {
  return "http://localhost:8080";
};

export const restService = {
  get: async (path: string) => {
    return fetcher(path, HTTPMethod.GET);
  },
};
