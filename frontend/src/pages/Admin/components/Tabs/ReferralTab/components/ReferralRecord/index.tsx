import { Button } from "@/shared/components/ui/Buttons/Button";
import { ButtonIcon } from "@/shared/components/ui/Buttons/ButtonIcon";
import { getHostname } from "@/utils/getHostname";
import { FC } from "react";
import { format } from "date-fns";
import { DeleteReferralButton } from "../DeleteReferralButton";

type Props = {
  code: string;
  description: string;
  time: string;
  date: string;
  visits: number;
};

export const ReferralRecord: FC<Props> = ({
  code,
  date,
  description,
  visits,
  time,
}) => {
  const url = `${getHostname([{ name: "code", value: code }])}`;
  return (
    <>
      <div className="w-full min-h-[1.875vw] flex items-center justify-between bg-white px-[0.625vw] py-[0.625vw] rounded-[0.313vw] border-[#ABABAB] border-[0.016vw]">
        <div>
          <div className="text-[0.677vw] font-medium">{url}</div>
          <div className="text-[#999DA6] text-[0.521vw]">{description}</div>
        </div>
        <div className="flex gap-x-[0.417vw]">
          <Button className="!py-[0.052vw] !px-[0.417vw] !rounded-[0.24vw] !text-[0.521vw]">
            {visits} Переходов
          </Button>
          <Button className="!py-[0.052vw] !px-[0.417vw] !rounded-[0.24vw] !text-[0.521vw]">
            {time}
          </Button>
          <Button className="!py-[0.052vw] !px-[0.417vw] !rounded-[0.24vw] !text-[0.521vw]">
            {format(date, "dd.MM.yyyy")}
          </Button>
          <DeleteReferralButton code={code} />
        </div>
      </div>
    </>
  );
};
