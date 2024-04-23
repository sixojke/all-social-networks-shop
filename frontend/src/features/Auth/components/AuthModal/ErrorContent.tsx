import { Button } from "@/shared/components/ui/Buttons/Button";
import { FC } from "react";

type Props = {
  setContentType: () => void;
};

export const ErrorContent: FC<Props> = ({ setContentType }) => {
  return (
    <div className="w-full mt-32 flex flex-col gap-y-5 items-center justify-center">
      <Button onClick={setContentType}>Повторить попытку</Button>
    </div>
  );
};
