import { FormInput } from "@/shared/components/common/Form/FormInput";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { yupResolver } from "@hookform/resolvers/yup";
import { FormProvider, useForm } from "react-hook-form";
import { VerifyFormValues, verifySchema } from "../../constants/verify";
import { FC, useContext } from "react";
import { useSignUpVerifyMutation } from "../../service";
import { ModalContext } from "@/shared/contexts/Modal";

type Props = {
  userId: number;
};

export const VerifyContent: FC<Props> = ({ userId }) => {
  const modalContext = useContext(ModalContext);
  const onHide = () => {
    modalContext?.hideModal();
  };
  const [verify] = useSignUpVerifyMutation();
  const defaultValues: VerifyFormValues = {
    code: null,
    id: userId,
  };
  const formApi = useForm<VerifyFormValues>({
    mode: "onChange",
    defaultValues,
    resolver: yupResolver(verifySchema),
  });
  const { handleSubmit } = formApi;
  const onSubmit = (data: VerifyFormValues) => {
    verify({ id: userId, code: data.code as string })
      .unwrap()
      .then(() => {
        onHide();
      });
  };
  return (
    <FormProvider {...formApi}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="w-full flex flex-col gap-y-5 items-center justify-center">
          <FormInput border max={6} placeholder="КОД" name="code" />
          <Button type="submit" className="bg-main-dark-blue">
            Подтвердить
          </Button>
        </div>
      </form>
    </FormProvider>
  );
};
