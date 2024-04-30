import { useDeleteReferralLinkMutation } from "@/entities/referral";
import { ButtonIcon } from "@/shared/components/ui/Buttons/ButtonIcon";
import { FC } from "react";

type Props = {
  code: string;
};

export const DeleteReferralButton: FC<Props> = ({ code }) => {
  const [deleteReferral] = useDeleteReferralLinkMutation();
  const deleteHandler = () => {
    deleteReferral({ code: code });
  };
  return (
    <ButtonIcon
      onClick={deleteHandler}
      className="!py-[0.417vw] !px-[0.469vw] !rounded-[0.26vw] !text-[0.938vw]"
      buttonType="delete"
    />
  );
};
