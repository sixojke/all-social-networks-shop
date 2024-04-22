import { Product } from "@/entities/products";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { FC } from "react";
import Image from "next/image";

type Props = {
  product: Product;
};

export const ProductLayout: FC<Props> = ({
  product: { id, img, description, name, price, quantity, quantity_sales },
}) => {
  return (
    <div className="bg-[#FBFCFF] border-main-light-blue border-[1px] rounded-2xl flex justify-between px-3 py-3 items-center">
      <div className="flex items-center gap-x-5">
        <div className="rounded-2xl w-14 h-14 flex justify-center items-center bg-main-dark-blue text-main-white">
          {id}
        </div>
        <div>
          <div className="text-lg text-main-black">{name}</div>
          <div className="text-main-light-gray text-[16px]">{description}</div>
        </div>
      </div>
      <div className="flex items-center gap-x-3 font-semiboldsemi">
        <Button
          disabled
          border
          className="bg-main-white !text-main-dark-blue text-[14.5px] !py-[8px] !px-[12px]"
        >
          {price} руб
        </Button>
        <Button
          disabled
          border
          className="bg-main-white !text-main-dark-blue text-[14.5px] !py-[8px] !px-[12px]"
        >
          {quantity} шт.
        </Button>
        <Button
          disabled
          border
          className="bg-main-white !text-main-dark-blue text-[14.5px] !p-0 !py-[8px] !px-[12px]"
        >
          {quantity_sales}
        </Button>
        <Button className="bg-main-dark-blue text-[14.5px] !py-[8px] !px-[12px]">
          Купить
        </Button>
      </div>
    </div>
  );
};
