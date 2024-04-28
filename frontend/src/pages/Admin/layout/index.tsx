import { ProtectedPageWrapper } from "@/hocs/ProtectedPageWrapper";
import { Navbar } from "../components/Navbar";
import { Panel } from "../components/Panel";
import Tabs from "@mui/joy/Tabs";

export const AdminLayout = () => {
  // After backend fix, replace React empty Fragment to ProtectedPageWrapper
  return (
    <>
      <div className="flex w-full overflow-x-hidden h-screen overflow-y-hidden absolute top-0 z-50">
        <Tabs defaultValue={"appSettings"} orientation="vertical">
          <Navbar />
          <div className="w-screen h-screen bg-[#F4F4F4]">
            <Panel />
          </div>
        </Tabs>
      </div>
    </>
  );
};
