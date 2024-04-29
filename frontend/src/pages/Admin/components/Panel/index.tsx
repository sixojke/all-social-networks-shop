import { Breadcrumbs } from "@/shared/components/ui/Breadcrumbs";
import { AppSettingsTab } from "../Tabs/AppSettingsTab";
import { ReferralTab } from "../Tabs/ReferralTab";

export const Panel = () => {
  return (
    <div className="bg-[#F4F4F4] w-full">
      <Breadcrumbs
        breadcrumbs={[
          { id: 1, label: "Test" },
          { id: 2, label: "Test 2" },
        ]}
      />
      <AppSettingsTab />
      <ReferralTab />
    </div>
  );
};
