import { Breadcrumbs } from "@/shared/components/ui/Breadcrumbs";
import { AppSettingsTab } from "../Tabs/AppSettingsTab";

export const Panel = () => {
  return (
    <div className="bg-[#F4F4F4] w-full pl-7 pt-12">
      <Breadcrumbs
        breadcrumbs={[
          { id: 1, label: "Test" },
          { id: 2, label: "Test 2" },
        ]}
      />
      <AppSettingsTab />
    </div>
  );
};
