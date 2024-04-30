import { useCreateReferralLinkMutation } from "@/entities/referral";
import { FormInput } from "@/shared/components/common/Form/FormInput";
import { FormTextarea } from "@/shared/components/common/Form/FormTextarea";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { yupResolver } from "@hookform/resolvers/yup";
import { FC } from "react";
import { FormProvider, useForm } from "react-hook-form";
import * as yup from "yup";

const schema = yup.object().shape({
  description: yup.string().required("Заполните поле").nullable(),
});

type FormValues = yup.InferType<typeof schema>;

type Props = {
  onHide: () => void;
};

export const CreateReferralModal: FC<Props> = ({ onHide }) => {
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
    onHide();
  };
  return (
    <div className="bg-white w-[22.656vw] flex flex-col justify-center items-center rounded-[0.417vw] py-[1.302vw]">
      <FormProvider {...formApi}>
        <form
          className="flex flex-col justify-center items-center"
          onSubmit={handleSubmit(onSubmit)}
        >
          <p className="text-[1.042vw] text-[#A1A1A1] font-semibold">
            Добавить Описание
          </p>
          <div className="w-[18.854vw]">
            <FormTextarea
              limit={200}
              autoComplete="off"
              className="mt-[0.26vw] h-[5.625vw]"
              name="description"
            />
          </div>
          <Button className="mt-[0.521vw]" type="submit">
            Создать ссылку
          </Button>
        </form>
      </FormProvider>
    </div>
  );
};
