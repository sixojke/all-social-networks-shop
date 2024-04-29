import { useCreateReferralLinkMutation } from "@/entities/referral";
import { FormInput } from "@/shared/components/common/Form/FormInput";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { yupResolver } from "@hookform/resolvers/yup";
import { FormProvider, useForm } from "react-hook-form";
import * as yup from "yup";

const schema = yup.object().shape({
  description: yup.string().required("Заполните поле").nullable(),
});

type FormValues = yup.InferType<typeof schema>;

export const CreateReferralModal = () => {
  const [createReferral] = useCreateReferralLinkMutation();
  const defaultValues: FormValues = { description: null };
  const formApi = useForm({
    mode: "onChange",
    resolver: yupResolver(schema),
    defaultValues,
  });
  const { handleSubmit } = formApi;
  const onSubmit = (data: FormValues) => {
    createReferral({ description: data.description as string });
  };
  return (
    <div className="bg-white p-[1.563vw] rounded-[0.521vw]">
      <FormProvider {...formApi}>
        <form onSubmit={handleSubmit(onSubmit)}>
          <FormInput placeholder="Описание" name="description" />
          <Button className="mt-[0.521vw]" type="submit">
            Создать ссылку
          </Button>
        </form>
      </FormProvider>
    </div>
  );
};
