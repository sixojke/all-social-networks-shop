import {
  SignUpFormValues,
  signUpSchema,
} from "@/features/Auth/constants/signUp";
import { FormInput } from "@/shared/components/common/Form/FormInput";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { yupResolver } from "@hookform/resolvers/yup";
import { FormProvider, useForm } from "react-hook-form";
import { useSignUpMutation } from "../../service";
import { ContentType } from ".";
import { FC } from "react";

type Props = {
  setErrorContent: () => void;
  setContentType: (content: ContentType) => void;
  setUserId: (userId: number) => void;
  setPassword: (value: string) => void;
  setLogin: (value: string) => void;
};

export const SignUpContent: FC<Props> = ({
  setContentType,
  setUserId,
  setLogin,
  setPassword,
  setErrorContent,
}) => {
  const [signUp] = useSignUpMutation();
  const defaultValues: SignUpFormValues = {
    login: null,
    email: null,
    password: null,
  };
  const formApi = useForm<SignUpFormValues>({
    mode: "onChange",
    defaultValues,
    resolver: yupResolver(signUpSchema),
  });
  const { handleSubmit } = formApi;
  const onSubmit = (data: SignUpFormValues) => {
    setContentType("loading");
    signUp({
      email: data.email as string,
      password: data.password as string,
      username: data.login as string,
    })
      .unwrap()
      .then((res) => {
        setUserId(res.id);
        setPassword(data.password as string);
        setLogin(data.login as string);
        setContentType("verify");
      })
      .catch(() => {
        setErrorContent();
      });
  };
  return (
    <FormProvider {...formApi}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="w-full flex flex-col gap-y-5 items-center justify-center">
          <FormInput
            className="bg-main-light-gray"
            border
            placeholder="Логин"
            name="login"
          />
          <FormInput
            className="bg-main-light-gray"
            border
            placeholder="Email"
            name="email"
            type="email"
          />
          <FormInput
            className="bg-main-light-gray"
            border
            placeholder="Пароль"
            name="password"
            type="password"
          />
          <Button type="submit" className="w-32 bg-main-dark-green">
            Создать
          </Button>
        </div>
      </form>
    </FormProvider>
  );
};
