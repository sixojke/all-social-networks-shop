import classNames from "classnames";
import { ComponentProps, FC, forwardRef } from "react";
import SearchIconImage from "@/assets/icons/search-gray.svg";

type Props = {
  border?: boolean;
  error?: boolean;
  wrapperClassname?: string;
  searchIcon?: boolean;
  rounded?: string;
} & ComponentProps<"input">;

export const Input: FC<Props> = forwardRef(
  (
    {
      placeholder = "Введите",
      rounded,
      wrapperClassname,
      searchIcon,
      className,
      border,
      error,
      ...props
    },
    ref
  ) => {
    return (
      <div
        className={classNames(
          "bg-main-light-gray",
          "placeholder-main-light-green",
          "text-[0.938vw]",
          "flex",
          "items-center",
          "p-0",
          "transition",
          "px-[1.042vw]",
          "font-normal",
          "h-[2.5vw]",
          "py-[0.234vw]",
          "rounded-[0.521vw]",
          "gap-x-3",
          "select-none",
          "w-full",
          "outline-none",
          {
            ["border-solid border-main-dark-green border-[1px]"]:
              border && !error,
            ["border-none"]: !border && !error,
            ["border-solid !border-main-error-dark-red border-[1px] placeholder-main-error-dark-red caret-main-error-dark-red"]:
              error,
          },
          className,
          wrapperClassname
        )}
      >
        {searchIcon && <SearchIconImage />}
        <input
          className={classNames(
            "outline-none placeholder-main-dark-green text-main-dark-green caret-main-green-gray w-full !rounded-none !bg-[#ff000000]",
            className
          )}
          ref={ref}
          placeholder={placeholder}
          {...props}
        />
      </div>
    );
  }
);

Input.displayName = "Input";
