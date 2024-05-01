import { ProtectedPageWrapper } from "@/hocs/ProtectedPageWrapper";
import { Navbar } from "../components/Navbar";
import { Panel } from "../components/Panel";
import Tabs from "@mui/joy/Tabs";

export const AdminLayout = () => {
  // TODO After backend fix, replace React empty Fragment to ProtectedPageWrapper
  return (
    <ProtectedPageWrapper>
      <div className="flex w-full overflow-x-hidden h-screen overflow-y-hidden absolute top-0 z-50">
        <Tabs defaultValue={"profile"} orientation="vertical">
          <Navbar />
          <div className="w-[calc(100vw-11.302vw)] h-screen bg-[#F4F4F4] pt-[2.5vw] pl-[1.042vw] pr-[2.604vw]">
            <Panel />
          </div>
        </Tabs>
      </div>
    </ProtectedPageWrapper>
  );
};
