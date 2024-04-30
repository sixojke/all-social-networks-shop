import TabPanel from "@mui/joy/TabPanel";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { ReferralRecord } from "./components/ReferralRecord";
import { useContext } from "react";
import { ModalContext } from "@/shared/contexts/Modal";
import { CreateReferralModal } from "./components/CreateReferralModal";
import { useGetRefferalStatsQuery } from "@/entities/referral";
import { QueryHandler } from "@/widgets/QueryHandler";

export const ReferralTab = () => {
  const modalContext = useContext(ModalContext);
  const addReferralHandler = () => {
    modalContext?.showModal(
      <CreateReferralModal onHide={modalContext.hideModal} />
    );
  };
  const {
    data: referrals,
    isError,
    isLoading,
  } = useGetRefferalStatsQuery({ limit: 10, page: 1 });

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
      <QueryHandler
        isError={isError}
        isLoading={isLoading}
        errorLabel="Произошла ошибка при загрузке реферальных ссылок"
      />
      {!referrals?.Pagination.data && !isError && !isLoading && (
        <div className="w-full flex justify-center flex-col items-center gap-y-9 mt-24 text-3xl font-semibold text-main-dark-green">
          Реферальные ссылки не найдены
        </div>
      )}
      <div className="mt-[0.625vw] flex flex-col gap-y-[0.625vw]">
        {referrals?.Pagination?.data?.map((referral) => {
          return (
            <ReferralRecord
              visits={referral.total_visitors}
              date={referral.created_at}
              description={referral.description}
              code={referral.referral_code}
              time={"0 Ч."}
              key={referral.referral_code}
            />
          );
        })}
      </div>
    </TabPanel>
  );
};