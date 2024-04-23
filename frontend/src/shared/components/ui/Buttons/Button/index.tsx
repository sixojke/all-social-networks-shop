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
        "bg-main-black",
        "text-main-white",
        "rounded-[10px]",
        "text-[18px]",
        "p-0",
        "duration-150",
        "px-[20px]",
        "font-semibold",
        "py-[10px]",
        "select-none",
        {
          ["active:brightness-125"]: !disabled,
          ["hover:brightness-90"]: !disabled,
          ["border-solid border-main-dark-gray border-opacity-30 border-[1px]"]: border,
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
