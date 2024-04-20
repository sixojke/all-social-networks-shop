import classNames from "classnames";
import { ComponentProps, FC } from "react";

type Props = {
  border?: boolean;
} & ComponentProps<"button">;

export const Button: FC<Props> = ({
  className,
  disabled,
  children,
  border,
  ...props
}) => {
  return (
    <button
      disabled={disabled}
      className={classNames(
        "bg-main-blue",
        "text-main-white",
        "rounded-3xl",
        "text-[18px]",
        "p-0",
        "transition",
        "px-[20px]",
        "font-semibold",
        "py-[10px]",
        "select-none",
        {
          ["active:brightness-125"]: !disabled,
          ["hover:bg-main-blue-gray"]: !disabled,
          ["border-solid border-main-blue border-2"]: border,
          ["border-none"]: !border,
        },
        className
      )}
      {...props}
    >
      {children}
    </button>
  );
};
