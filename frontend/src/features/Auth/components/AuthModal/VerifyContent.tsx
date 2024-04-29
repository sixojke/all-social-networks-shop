import { FormInput } from "@/shared/components/common/Form/FormInput";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { yupResolver } from "@hookform/resolvers/yup";
import { FormProvider, useForm } from "react-hook-form";
import { VerifyFormValues, verifySchema } from "../../constants/verify";
import { FC, useContext } from "react";
import { useSignInMutation, useSignUpVerifyMutation } from "../../service";
import { ModalContext } from "@/shared/contexts/Modal";
import { SignInResponse } from "../../types";
import { useRouter } from "next/navigation";

type Props = {
  userId: number;
  login: string;
  password: string;
  setErrorContent: () => void;
};

export const VerifyContent: FC<Props> = ({
  userId,
  setErrorContent,
  login,
  password,
}) => {
  const router = useRouter();
  const modalContext = useContext(ModalContext);
  const onHide = () => {
    modalContext?.hideModal();
  };
  const [verify] = useSignUpVerifyMutation();
  const [signIn] = useSignInMutation();
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
        signIn({
          password: password,
          username: login,
        })
          .unwrap()
          .then((res: SignInResponse) => {
            localStorage.setItem("accessToken", res.accessToken);
            localStorage.setItem("refreshToken", res.refreshToken);
            router.refresh();
          })
          .catch(() => {
            setErrorContent();
          });
      })
      .catch(() => {
        setErrorContent();
      });
  };
  return (
    <FormProvider {...formApi}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="w-full flex flex-col gap-y-5 items-center justify-center">
          <FormInput border max={6} placeholder="КОД" name="code" />
          <Button type="submit" className="bg-main-dark-green">
            Подтвердить
          </Button>
        </div>
      </form>
    </FormProvider>
  );
};
