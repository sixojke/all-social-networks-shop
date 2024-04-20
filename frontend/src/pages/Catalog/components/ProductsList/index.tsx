import { Product } from "@/entities/products";
import { ProductLayout } from "@/widgets/Product";
import { FC, Fragment } from "react";

type Props = {
  products: Product[] | null;
};

export const ProductsList: FC<Props> = ({ products }) => {
  return (
    <div className="flex flex-col gap-y-6">
      {products?.map((product) => (
        <Fragment key={product.id}>
          <ProductLayout key={product.id} product={product} />
        </Fragment>
      ))}
    </div>
  );
};
