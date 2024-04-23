import { Button } from "@/shared/components/ui/Buttons/Button";
import { Loader } from "@/shared/components/ui/Loader";
import { useRouter } from "next/navigation";
import { FC } from "react";

type Props = {
  isLoading: boolean;
  isError: boolean;
  errorLabel: string;
};

export const QueryHandler: FC<Props> = ({ errorLabel, isError, isLoading }) => {
  const router = useRouter();
  const reloadPage = () => {
    router.refresh();
  };
  return (
    <>
      {isLoading && (
        <div className="w-full flex justify-center mt-24">
          <Loader />
        </div>
      )}
      {isError && (
        <div className="w-full flex justify-center flex-col items-center gap-y-9 mt-24 text-3xl font-bold text-main-error-dark-red">
          {errorLabel}
          <Button onClick={reloadPage} className="bg-main-dark-green">
            Перезагрузить страницу
          </Button>
        </div>
      )}
    </>
  );
};
