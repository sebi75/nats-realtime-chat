import { Card, CardContent } from "@/components/ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { type FunctionComponent } from "react";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import { MoreHorizontal } from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useRouter } from "next/router";

type Props = {
  className?: string;
  avatarUrl?: string;
  username: string;
};

export const UserCard: FunctionComponent<Props> = ({
  avatarUrl,
  className,
  username,
}) => {
  const router = useRouter();
  const avatarFallback = username.slice(0, 2).toUpperCase();

  const handleSignout = async () => {
    try {
      const token = localStorage.getItem("token");
      if (!token) {
        await router.push("/auth/login");
      } else {
        localStorage.removeItem("token");
        await router.push("/auth/login");
      }
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <Card className={cn("", className)}>
      <CardContent className="flex items-center justify-between gap-3 p-3">
        {/* card leftside */}
        <div className="flex flex-row items-center gap-3">
          <Avatar>
            <AvatarImage src={avatarUrl} alt={username} />
            <AvatarFallback>{avatarFallback ?? "AV"}</AvatarFallback>
          </Avatar>
          <p className="text-sm text-gray-500">{`@${username}`}</p>
        </div>
        {/* card rightside / menu options */}
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant={"ghost"}>
              <MoreHorizontal size={21} />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-56">
            <DropdownMenuLabel>My Account</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuGroup>
              <DropdownMenuItem>Profile</DropdownMenuItem>
              <DropdownMenuItem>Settings</DropdownMenuItem>
            </DropdownMenuGroup>
            <DropdownMenuSeparator />
            <DropdownMenuItem onClick={handleSignout}>
              Sign out
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </CardContent>
    </Card>
  );
};
