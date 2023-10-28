import { SigninForm } from "@/components/auth/SigninForm";
import { CustomHead } from "@/components/common/CustomHead";
import { type GetServerSideProps, type GetServerSidePropsContext } from "next";

export default function SigninPage() {
  return (
    <>
      <CustomHead title="Signin" />
      <div className="flex min-h-screen items-center justify-center">
        <SigninForm />
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
