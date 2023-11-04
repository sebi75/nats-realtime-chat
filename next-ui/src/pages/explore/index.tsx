import { useVerifyToken } from "@/components/auth/hooks";
import { CustomHead } from "@/components/common/CustomHead";
import { ChatWrapper } from "@/components/layout/ChatWrapper";

export default function Explore() {
  const { data } = useVerifyToken();
  console.log(data);
  return (
    <>
      <CustomHead title="Explore" />
      <ChatWrapper>
        <div className="flex flex-col items-center justify-center">
          <h1 className="text-4xl font-bold">Explore</h1>
          <p className="text-gray-500">Explore page</p>
        </div>
      </ChatWrapper>
    </>
  );
}
