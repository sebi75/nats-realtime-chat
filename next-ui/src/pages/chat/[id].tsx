import { CustomHead } from "@/components/common/custom-head";
import { ChatWrapper } from "@/components/layout/chat-wrapper";

/**
 * This page will display an individual chat.
 * This is the main entry for any type of chat (group or individual).
 */
export default function Chat() {
  return (
    <>
      <CustomHead title="Chat | ..." />
      <ChatWrapper>
        <div>individual chat</div>
      </ChatWrapper>
    </>
  );
}
