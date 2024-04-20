import {
  ComponentProps,
  FC,
  PropsWithChildren,
  useEffect,
  useState,
} from "react";
import { Portal } from "react-portal";
import classNames from "classnames";
import { useDomIsLoaded } from "@/hooks/useDomIsLoaded";

type Props = {
  active: boolean;
  onHide: (state: boolean) => void;
} & ComponentProps<"div">;

export const Modal: FC<PropsWithChildren<Props>> = ({
  active,
  children,
  onHide,
  className,
  ...props
}) => {
  const domIsLoaded = useDomIsLoaded();

  if (!domIsLoaded) return;

  return (
    <>
      <Portal>
        <div
          {...props}
          className={classNames(
            "bg-main-black bg-opacity-40 opacity-0 h-screen w-screen fixed top-0 left-0 flex justify-center items-center pointer-events-none transition-transform z-50",
            className,
            {
              ["opacity-100 pointer-events-auto"]: active,
            }
          )}
          onMouseDown={(e) => {
            if (e.target === e.currentTarget) {
              onHide(false);
            }
          }}
        >
          <div
            className={classNames("w-auto flex transition-transform", {
              ["transition-transform"]: active,
            })}
            onClick={(e) => e.stopPropagation()}
          >
            <div className={"flex-1"}>{children}</div>
          </div>
        </div>
      </Portal>
    </>
  );
};
