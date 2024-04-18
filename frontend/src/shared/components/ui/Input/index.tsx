import classNames from "classnames";
import { ComponentProps, FC } from "react";

type Props = {
  border?: boolean;
} & ComponentProps<"input">;

export const Input: FC<Props> = ({
  placeholder = "Введите",
  className,
  children,
  border,
  ...props
}) => {
  return (
    <input
      placeholder={placeholder}
      className={classNames(
        "bg-main-blue",
        "rounded-3xl",
        "placeholder-main-blue-gray",
        "text-main-blue-gray",
        "text-[18px]",
        "p-0",
        "transition",
        "px-[20px]",
        "font-semibold",
        "py-[4.5px]",
        "select-none",
        "w-full",
        "outline-none",
        "caret-main-blue-gray",
        {
          ["border-solid border-main-blue border-2"]: border,
          ["border-none"]: !border,
        },
        className
      )}
      {...props}
    />
  );
};
