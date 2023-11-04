import { cn } from "@/lib/utils";
import { type PropsWithChildren, type FunctionComponent } from "react";
import { UserCard } from "./components/user-card";
import { Label } from "@/components/ui/label";

type Props = PropsWithChildren & {
  className?: string;
};

export const ChatLeftSidebar: FunctionComponent<Props> = ({ className }) => {
  // sidebar with the list of chats, search, and user profile down below
  return (
    <div
      className={cn(
        "flex max-h-screen min-h-screen max-w-sm flex-col justify-between border-r p-4",
        className
      )}
    >
      <div className="flex w-full flex-col items-center gap-5">
        <Label className="text-2xl">Chats</Label>
      </div>
      {/* user card */}
      <UserCard
        username="mock-username"
        avatarUrl="https://pbs.twimg.com/profile_images/1350895249678348292/RS1Aa0iK_400x400.jpg"
      />
    </div>
  );
};
