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
    <div className="bg-main-light-blue bg-opacity-35 border-main-light-blue border-[1px] rounded-2xl flex justify-between px-4 py-4 items-center">
      <div className="flex items-center gap-x-5">
        <div className="rounded-2xl w-20 h-20 flex justify-center items-center bg-main-dark-blue text-main-white">
          {id}
        </div>
        <div>
          <div className="text-xl text-main-black">{name}</div>
          <div className="text-main-light-gray text-xl">{description}</div>
        </div>
      </div>
      <div className="flex items-center gap-x-3">
        <Button border className="bg-main-white !text-main-dark-blue">
          {price} руб
        </Button>
        <Button border className="bg-main-white !text-main-dark-blue">
          {quantity} шт.
        </Button>
        <Button border className="bg-main-white !text-main-dark-blue">
          {quantity_sales}
        </Button>
        <Button className="bg-main-dark-blue">Купить</Button>
      </div>
    </div>
  );
};
