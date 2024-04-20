import { CustomSelect } from "@/shared/components/ui/CustomSelect";
import { IOption } from "@/shared/components/ui/CustomSelect/select.types";
import { ErrorMessage } from "@hookform/error-message";
import { FC } from "react";
import { Controller, useFormContext } from "react-hook-form";

type Props = {
  options: IOption[];
  name: string;
  width?: string;
  isClearable?: boolean;
  isLoading?: boolean;
  label?: string;
  placeholder?: string;
  disable?: boolean;
  classname?: string;
};

export const FormSelect: FC<Props> = ({
  name,
  classname,
  disable,
  label,
  isLoading,
  isClearable,
  placeholder = "Выберите",
  width,
  options,
}) => {
  const {
    control,
    formState: { errors },
  } = useFormContext();

  return (
    <div className="flex flex-col">
      <Controller
        name={name}
        control={control}
        render={({ field: { onChange, value } }) => (
          <CustomSelect
            isClearable={isClearable}
            isLoading={isLoading}
            options={options}
            placeholder={placeholder}
            width={width}
            isDisabled={disable}
            label={label}
            className={classname}
            onChange={onChange}
            value={value}
          />
        )}
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
