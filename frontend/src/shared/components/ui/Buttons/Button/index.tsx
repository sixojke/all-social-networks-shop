import classNames from "classnames";
import { ComponentProps, FC } from "react";

type Props = {
  border?: boolean;
} & ComponentProps<"button">;

export const Button: FC<Props> = ({
  className,
  children,
  border,
  ...props
}) => {
  return (
    <button
      className={classNames(
        "bg-main-blue",
        "rounded-3xl",
        "text-main-white",
        "text-[18px]",
        "p-0",
        "hover:brightness-110",
        "transition",
        "px-[16px]",
        "font-semibold",
        "py-[4.5px]",
        "select-none",
        "active:brightness-125",
        {
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
