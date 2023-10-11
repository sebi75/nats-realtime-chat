import { SigninForm } from "@/components/auth/SigninForm";
import { GetServerSideProps, GetServerSidePropsContext } from "next";

export default function SigninPage() {
  return (
    <>
      <SigninForm />
    </>
  );
}

// do verification here whether the user is already logged in or not
export const getServerSideProps: GetServerSideProps = async ({
  req,
}: GetServerSidePropsContext) => {
  await new Promise((resolve) => setTimeout(resolve, 1000));
  return {
    props: {},
  };
};
