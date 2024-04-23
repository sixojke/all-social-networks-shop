import { ProtectedComponent } from "@/hocs/ProtectedComponent";
import { Navbar } from "../components/Navbar";
import { Panel } from "../components/Panel";

export const AdminLayout = () => {
  return (
    <ProtectedComponent>
      <div className="flex w-screen h-screen fixed overflow-y-hidden mt-6">
        <Navbar />
        <Panel />
      </div>
    </ProtectedComponent>
  );
};
