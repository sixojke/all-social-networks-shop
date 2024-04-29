import TabPanel from "@mui/joy/TabPanel";
import { CategoryRecord } from "./components/CategoryRecord";
import { ButtonIcon } from "@/shared/components/ui/Buttons/ButtonIcon";

export const AppSettingsTab = () => {
  return (
    <TabPanel
      sx={{
        [`&.MuiTabPanel-root`]: {
          padding: 0,
        },
      }}
      value="appSettings"
    >
      <div className="flex items-center justify-between">
        <p className="text-[1.302vw] font-bold">Категории</p>
        <ButtonIcon buttonType="add" />
      </div>
      <div className="mt-[0.625vw]">
        <CategoryRecord />
      </div>
    </TabPanel>
  );
};
