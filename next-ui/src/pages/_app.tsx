import { type AppType } from "next/app";
import "@/styles/globals.css";
import { useState } from "react";
import {
  QueryClient,
  QueryClientProvider,
  // Hydrate,
} from "@tanstack/react-query";

const MyApp: AppType = ({ Component, pageProps: { ...pageProps } }) => {
  const [queryClient] = useState(
    () =>
      new QueryClient({
        defaultOptions: {
          queries: {
            refetchOnWindowFocus: false,
            retry: 1,
          },
        },
      })
  );
  return (
    <QueryClientProvider client={queryClient}>
      {/* <Hydrate state={pageProps.dehydratedState}> */}
      <Component {...pageProps} />
      {/* </Hydrate> */}
    </QueryClientProvider>
  );
};

export default MyApp;
