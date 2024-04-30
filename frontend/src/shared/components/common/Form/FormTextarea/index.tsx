import { Textarea } from "@/shared/components/ui/Textarea";
import { ErrorMessage } from "@hookform/error-message";
import { ComponentProps, FC } from "react";
import { Controller, useFormContext } from "react-hook-form";

type Props = {
  limit?: number;
  name: string;
} & ComponentProps<"textarea">;

export const FormTextarea: FC<Props> = ({
  name,
  className,
  placeholder,
  ...props
}) => {
  const {
    control,
    formState: { errors },
  } = useFormContext();
  return (
    <div className="flex flex-col w-full]">
      <Controller
        name={name}
        control={control}
        render={({ field: { onChange, value } }) => (
          <Textarea
            onChange={onChange}
            value={value}
            className={className}
            id={name}
            placeholder={placeholder}
            {...props}
          />
        )}
      />
      <ErrorMessage
        errors={errors}
        name={name}
        render={({ message }) => (
          <p className="text-main-error-red text-[0.833vw]">{message}</p>
        )}
      />
    </div>
  );
};
