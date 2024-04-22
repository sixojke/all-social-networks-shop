import classNames from "classnames";
import { ComponentProps, FC, forwardRef } from "react";
import searchIconImage from "@/assets/icons/search-gray.svg";
import Image from "next/image";

type Props = {
  border?: boolean;
  error?: boolean;
  searchIcon?: boolean;
  rounded?: string;
} & ComponentProps<"input">;

export const Input: FC<Props> = forwardRef(
  (
    {
      placeholder = "Введите",
      rounded,
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
          "bg-[#FBFCFF]",
          "placeholder-main-blue",
          "text-[18px]",
          "flex",
          "items-center",
          "p-0",
          "transition",
          "px-[20px]",
          "font-normal",
          "h-12",
          "py-[4.5px]",
          "gap-x-3",
          "select-none",
          "w-full",
          "outline-none",
          {
            ["border-solid border-main-blue border-[1px]"]: border && !error,
            ["border-none"]: !border && !error,
            ["border-solid !border-main-error-dark-red border-[1px] placeholder-main-error-dark-red caret-main-error-dark-red"]:
              error,
          },
          className
        )}
      >
        {searchIcon && (
          <Image
            className="fill-main-error-red"
            alt=""
            src={searchIconImage}
            width={15}
            height={15}
          />
        )}
        <input
          className={classNames(
            "outline-none placeholder-main-blue text-main-dark-blue bg-[#FBFCFF] caret-main-blue-gray w-full !rounded-none",
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
