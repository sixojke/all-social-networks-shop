import { Button } from "@/shared/components/ui/Buttons/Button";
import { ButtonIcon } from "@/shared/components/ui/Buttons/ButtonIcon";
import { FC } from "react";

type Props = {
  link: string;
  description: string;
  time: string;
  date: string;
};

export const ReferralRecord: FC<Props> = () => {
  return (
    <>
      <div className="w-full h-[1.875vw] flex items-center justify-between bg-white px-[0.625vw] py-[0.625vw] rounded-[0.313vw] border-[#ABABAB] border-[0.016vw]">
        <div className="text-[0.833vw] font-medium">Тестовая категория</div>
        <div className="flex gap-x-[0.417vw]">
          <ButtonIcon
            className="!py-[0.052vw] !px-[0.417vw] rounded-[0.26vw] !text-[0.938vw]"
            buttonType="add"
          />
          <Button className="!text-[0.521vw] !rounded-[0.313vw] !py-[0.26vw] !px-[0.677vw]">
            Удалить
          </Button>
          <Button className="!text-[0.521vw] !rounded-[0.313vw] !py-[0.26vw] !px-[0.677vw]">
            Изменить
          </Button>
        </div>
      </div>
    </>
  );
};
