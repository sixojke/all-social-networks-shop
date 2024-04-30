import classNames from "classnames";
import { isNull, isNumber } from "lodash";
import { ComponentProps, FC, forwardRef } from "react";

type Props = {
  limit?: number;
} & ComponentProps<"textarea">;

export const Textarea: FC<Props> = forwardRef(
  ({ className, limit, value, ...props }, ref) => {
    return (
      <div className="w-full">
        <textarea
          ref={ref}
          value={value ?? ""}
          maxLength={limit}
          className={classNames(
            "w-full min-h-[5.208vw] bg-[#F4F4F4] rounded-[0.573vw] outline-none resize-none px-[0.625vw] py-[0.26vw]",
            className
          )}
          {...props}
        />
        {isNumber(limit) && (
          <div className="mt-[-1.563vw] text-[0.729vw] w-full flex justify-end pb-[0.573vw] pr-[0.26vw] text-[#A1A1A1]">
            {`${isNull(value) ? 0 : String(value).length}/${limit}`}
          </div>
        )}
      </div>
    );
  }
);

Textarea.displayName = "Textarea";
