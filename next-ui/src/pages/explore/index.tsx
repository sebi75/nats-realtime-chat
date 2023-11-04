import { useVerifyToken } from "@/components/auth/hooks";
import { CustomHead } from "@/components/common/custom-head";
import { ChatWrapper } from "@/components/layout/chat-wrapper";

/**
 * This page will display recommendations for group chats in which a user
 * may be interested in joining based on their interests.
 */
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
