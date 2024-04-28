import { FC } from "react";

type Props = {
  breadcrumbs: {
    id: number;
    label: string;
    onClick?: () => void;
  }[];
};

export const Breadcrumbs: FC<Props> = ({ breadcrumbs }) => {
  return (
    <div className="text-main-dark-gray flex text-sm gap-x-1">
      {breadcrumbs.map((item, index) => (
        <p key={item.id} onClick={item.onClick}>
          <span className="select-none">{index > 0 && "/ "}</span>
          <span className="cursor-pointer hover:text-main-black duration-300">
            {item.label}
          </span>
        </p>
      ))}
    </div>
  );
};