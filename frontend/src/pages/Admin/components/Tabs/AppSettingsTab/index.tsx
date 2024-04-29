import TabPanel from "@mui/joy/TabPanel";
import { CategoryRecord } from "./components/CategoryRecord";

export const AppSettingsTab = () => {
  return (
    <TabPanel value="appSettings">
      <CategoryRecord />
    </TabPanel>
  );
};
