import { CustomSelect } from "@/shared/components/ui/CustomSelect";
import { IOption } from "@/shared/components/ui/CustomSelect/select.types";
import { Input } from "@/shared/components/ui/Input";
import { ErrorMessage } from "@hookform/error-message";
import { FC } from "react";
import { Controller, useFormContext } from "react-hook-form";

type Props = {
  name: string;
  width?: string;
  isClearable?: boolean;
  border?: boolean;
  isLoading?: boolean;
  label?: string;
  placeholder?: string;
  disable?: boolean;
  classname?: string;
};

export const FormInput: FC<Props> = ({
  name,
  classname,
  border,
  placeholder = "Введите",
}) => {
  const {
    register,
    formState: { errors },
  } = useFormContext();

  return (
    <div className="flex flex-col">
      <Input
        {...register(name)}
        className={classname}
        id={name}
        placeholder={placeholder}
        border={border}
      />
      <ErrorMessage
        errors={errors}
        name={name}
        render={({ message }) => (
          <p className="text-main-error-red">{message}</p>
        )}
      />
    </div>
  );
};
