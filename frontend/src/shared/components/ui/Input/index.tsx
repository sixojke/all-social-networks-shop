import classNames from "classnames";
import { isEmpty } from "lodash";
import { ComponentProps, FC, forwardRef } from "react";
import { FieldErrors, FieldValues } from "react-hook-form";

type Props = {
  border?: boolean;
  error?: boolean;
} & ComponentProps<"input">;

export const Input: FC<Props> = forwardRef(
  ({ placeholder = "Введите", className, border, error, ...props }, ref) => {
    return (
      <input
        ref={ref}
        placeholder={placeholder}
        className={classNames(
          "bg-main-light-blue",
          "rounded-xl",
          "placeholder-main-blue",
          "text-main-dark-blue",
          "text-[18px]",
          "p-0",
          "transition",
          "px-[20px]",
          "font-normal",
          "h-12",
          "py-[4.5px]",
          "select-none",
          "w-full",
          "outline-none",
          "caret-main-blue-gray",
          {
            ["border-solid border-main-blue border-2"]: border && !error,
            ["border-none"]: !border && !error,
            ["border-solid !border-main-error-dark-red border-2 placeholder-main-error-dark-red caret-main-error-dark-red"]:
              error,
          },
          className
        )}
        {...props}
      />
    );
  }
);
