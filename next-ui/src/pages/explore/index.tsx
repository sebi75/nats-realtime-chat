import { useVerifyToken } from "@/components/auth/hooks";
import { CustomHead } from "@/components/common/CustomHead";

export default function Explore() {
  const { data } = useVerifyToken();
  console.log(data);
  return (
    <>
      <CustomHead title="Explore" />
      <div>
        Hello from explore page. Here you will be able to find realtime chats
        which meet your interests
      </div>
    </>
  );
}
