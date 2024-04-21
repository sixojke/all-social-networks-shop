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
  setContentType: (content: ContentType) => void;
  setUserId: (userId: number) => void;
};

export const SignUpContent: FC<Props> = ({ setContentType, setUserId }) => {
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
        setContentType("verify");
      });
    console.log(data);
  };
  return (
    <FormProvider {...formApi}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="w-full flex flex-col gap-y-5 items-center justify-center">
          <FormInput placeholder="Логин" name="login" />
          <FormInput placeholder="Email" name="email" type="email" />
          <FormInput placeholder="Пароль" name="password" type="password" />
          <Button type="submit" className="w-32 bg-main-dark-blue">
            Создать
          </Button>
        </div>
      </form>
    </FormProvider>
  );
};
