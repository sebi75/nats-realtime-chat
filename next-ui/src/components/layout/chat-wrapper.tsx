import { cn } from "@/lib/utils";
import { type PropsWithChildren, type FunctionComponent } from "react";
import { ChatLeftSidebar } from "./chat-left-sidebar/chat-left-sidebar";

type Props = PropsWithChildren & {
  className?: string;
};

export const ChatWrapper: FunctionComponent<Props> = ({
  children,
  className,
}) => {
  return (
    <div className="flex flex-row">
      <ChatLeftSidebar />
      <div className={cn("p-5", className)}>{children}</div>
    </div>
  );
};
