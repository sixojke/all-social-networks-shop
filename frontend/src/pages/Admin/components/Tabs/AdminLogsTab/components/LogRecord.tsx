import { Button } from "@/shared/components/ui/Buttons/Button";
import { FC } from "react";
import { format } from "date-fns";

type Props = {
  adminName: string;
  adminAction: string;
  actionDate: string;
};

export const LogRecord: FC<Props> = ({
  actionDate,
  adminAction,
  adminName,
}) => {
  return (
    <>
      <div className="w-full min-h-[1.875vw] flex items-center justify-between bg-white px-[0.625vw] py-[0.625vw] rounded-[0.313vw] border-[#ABABAB] border-[0.016vw]">
        <div>
          <div className="text-[0.677vw] font-medium">{adminName}</div>
          <div className="text-[#999DA6] text-[0.521vw]">{adminAction}</div>
        </div>
        <div className="flex gap-x-[0.417vw]">
          <Button className="!text-[0.521vw] !rounded-[0.313vw] !py-[0.26vw] !px-[0.677vw]">
            {format(actionDate, "HH:mm")}
          </Button>
          <Button className="!text-[0.521vw] !rounded-[0.313vw] !py-[0.26vw] !px-[0.677vw]">
            {format(actionDate, "dd.MM.yyyy")}
          </Button>
        </div>
      </div>
    </>
  );
};
