import { SignupForm } from "@/components/auth/SignupForm";
import { CustomHead } from "@/components/common/custom-head";
import { type GetServerSideProps, type GetServerSidePropsContext } from "next";

export default function SignupPage() {
  return (
    <>
      <CustomHead title="Signup" />
      <div className="flex min-h-screen items-center justify-center">
        <SignupForm />
      </div>
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
