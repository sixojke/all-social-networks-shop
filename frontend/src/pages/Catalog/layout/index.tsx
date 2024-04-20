import {
  IGetAllProductsRequest,
  useGetAllProductsQuery,
} from "@/entities/products";
import { ProductsList } from "../components/ProductsList";
import { QueryHandler } from "@/widgets/QueryHandler";
import { Filters } from "../components/Filters";
import { useState } from "react";

export const CatalogPageLayout = () => {
  const [filters, setFilters] = useState<IGetAllProductsRequest>({});
  const { data, isLoading, isError } = useGetAllProductsQuery(filters);

  return (
    <>
      <main className="px-64 mt-10">
        <Filters setFilters={setFilters} />
        <QueryHandler
          isError={isError}
          isLoading={isLoading}
          errorLabel="При загрузке товаров произошла ошибка"
        />
        <ProductsList products={data?.Pagination.data ?? []} />
      </main>
    </>
  );
};
