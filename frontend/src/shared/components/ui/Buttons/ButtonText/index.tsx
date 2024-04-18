import classNames from "classnames";
import { ComponentProps, FC } from "react";

type Props = {} & ComponentProps<"button">;

export const ButtonText: FC<Props> = ({ className, children, ...props }) => {
  return (
    <button
      className={classNames(
        "text-main-black",
        "text-[18px]",
        "p-0",
        "hover:text-main-blue",
        "transition",
        "px-[16px]",
        "font-semibold",
        "py-[4.5px]",
        "select-none",
        "active:brightness-125",
        "border-none",
        className
      )}
      {...props}
    >
      {children}
    </button>
  );
};
