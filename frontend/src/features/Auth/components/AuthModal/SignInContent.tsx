import { SignInFormValues, signInSchema } from "@/features/Auth/constants/signIn";
import { FormInput } from "@/shared/components/common/Form/FormInput";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { yupResolver } from "@hookform/resolvers/yup";
import { FormProvider, useForm } from "react-hook-form";

export const SignInContent = () => {
  const defaultValues: SignInFormValues = {
    login: null,
    password: null,
  };
  const formApi = useForm<SignInFormValues>({
    mode: "onChange",
    defaultValues,
    resolver: yupResolver(signInSchema),
  });
  const { handleSubmit } = formApi;
  const onSubmit = (data: SignInFormValues) => {
    console.log(data);
  };
  return (
    <FormProvider {...formApi}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="w-full flex flex-col gap-y-5 items-center justify-center">
          <FormInput placeholder="Логин" name="login" />
          <FormInput placeholder="Пароль" name="password" type="password" />
          <Button type="submit" className="w-32">
            Войти
          </Button>
        </div>
      </form>
    </FormProvider>
  );
};
