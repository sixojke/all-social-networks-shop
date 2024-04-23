import { Product } from "@/entities/products";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { FC } from "react";
import Facebook from "@/assets/images/facebook.png";
import Image from "next/image";

type Props = {
  product: Product;
};

export const ProductLayout: FC<Props> = ({
  product: { id, img, description, name, price, quantity, quantity_sales },
}) => {
  return (
    <div className="bg-white border-main-dark-gray border-opacity-30 border-[0.5px] rounded-[10px] flex justify-between px-3 py-3 items-center">
      <div className="flex items-center gap-x-5">
        <div className="rounded-[10px] w-14 h-14 flex justify-center items-center bg-main-dark-green text-main-white relative">
          <Image alt="" className="rounded-[10px]" src={Facebook} fill />
        </div>
        <div>
          <div className="text-lg text-main-black">{name}</div>
          <div className="text-main-green-gray text-[16px] max-w-[500px]">
            {description}
          </div>
        </div>
      </div>
      <div className="flex items-center gap-x-3">
        <Button
          disabled
          border
          className="bg-white min-w-[70px] !text-main-black !text-[13px] !py-[8px] !px-[12px] !font-medium"
        >
          {price} руб
        </Button>
        <Button
          disabled
          border
          className="bg-white min-w-[70px] !text-main-black !text-[13px] !py-[8px] !px-[12px] !font-medium"
        >
          {quantity} шт.
        </Button>
        <Button
          disabled
          border
          className="bg-white min-w-[70px] !text-main-black !text-[13px] !p-0 !py-[8px] !px-[12px] !font-medium"
        >
          {quantity_sales}
        </Button>
        <Button className="bg-main-dark-green min-w-[70px] !text-[13px] !py-[8px] !px-[15px] !font-medium">
          Купить
        </Button>
      </div>
    </div>
  );
};
