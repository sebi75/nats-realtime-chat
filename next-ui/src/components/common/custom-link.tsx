import Link from "next/link";
import { cn } from "@/lib/utils";
import { type FunctionComponent, type PropsWithChildren } from "react";

type Props = PropsWithChildren & {
  href: string;
  className?: string;
};

export const CustomLink: FunctionComponent<Props> = ({
  href,
  className,
  children,
}) => {
  const isExternal = href.startsWith("http");
  if (isExternal) {
    return (
      <a
        href={href}
        className={cn("text-blue-500 hover:text-blue-600", className)}
        target="_blank"
        rel="noopener noreferrer"
      >
        {children}
      </a>
    );
  } else {
    return (
      <Link href={href} className={className}>
        {children}
      </Link>
    );
  }
};
