import TabPanel from "@mui/joy/TabPanel";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { ReferralRecord } from "./components/ReferralRecord";
import { useContext } from "react";
import { ModalContext } from "@/shared/contexts/Modal";
import { CreateReferralModal } from "./components/CreateReferralModal";

export const ReferralTab = () => {
  const modalContext = useContext(ModalContext);
  const addReferralHandler = () => {
    modalContext?.showModal(<CreateReferralModal />);
  };
  return (
    <TabPanel
      sx={{
        [`&.MuiTabPanel-root`]: {
          padding: 0,
        },
      }}
      value="referrals"
    >
      <div className="flex items-center justify-between">
        <p className="text-[1.302vw] font-bold">Все ссылки</p>
        <Button
          className="font-bold !text-[0.938vw] !px-[0.625vw] !py-[0.208vw] !rounded-[0.417vw]"
          onClick={addReferralHandler}
        >
          +
        </Button>
      </div>
      <div className="mt-[0.625vw]">{/* <ReferralRecord /> */}</div>
    </TabPanel>
  );
};
