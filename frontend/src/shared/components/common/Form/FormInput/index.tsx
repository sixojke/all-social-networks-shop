import { Input } from "@/shared/components/ui/Input";
import { ErrorMessage } from "@hookform/error-message";
import { ComponentProps, FC } from "react";
import { useFormContext } from "react-hook-form";

type Props = {
  name: string;
  width?: string;
  isClearable?: boolean;
  border?: boolean;
  isLoading?: boolean;
  label?: string;
} & ComponentProps<"input">;

export const FormInput: FC<Props> = ({
  name,
  className,
  border,
  placeholder = "Введите",
  ...props
}) => {
  const {
    register,
    getFieldState,
    formState: { errors },
  } = useFormContext();
  return (
    <div className="flex flex-col w-full h-[70px]">
      <Input
        error={getFieldState(name).invalid}
        className={className}
        id={name}
        placeholder={placeholder}
        border={border}
        {...register(name)}
        {...props}
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
