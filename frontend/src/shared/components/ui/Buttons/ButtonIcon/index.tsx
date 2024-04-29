import { ComponentProps, FC } from "react";
import { Button } from "../Button";
import DeleteIcon from "@/assets/icons/delete.svg";

type Props = {
  buttonType: "add" | "delete";
} & ComponentProps<"button">;

export const ButtonIcon: FC<Props> = ({ buttonType, ...props }) => {
  const getContent = () => {
    switch (buttonType) {
      case "add":
        return "+";
      case "delete":
        return <DeleteIcon />;
    }
  };
  return (
    <Button
      className="font-bold !text-[0.938vw] !px-[0.625vw] !py-[0.208vw] !rounded-[0.417vw]"
      {...props}
    >
      {getContent()}
    </Button>
  );
};
