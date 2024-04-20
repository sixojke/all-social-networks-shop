import {
  IGetAllProductsRequest,
  useGetAllProductsQuery,
} from "@/entities/products";
import { ProductsList } from "../components/ProductsList";
import { QueryHandler } from "@/widgets/QueryHandler";
import { Filters } from "../components/Filters";
import { useState } from "react";
import { Pagination } from "@/shared/components/ui/Pagination";

export const CatalogPageLayout = () => {
  const [filters, setFilters] = useState<IGetAllProductsRequest>({});
  const [offset, setOffset] = useState<number>(11);
  const { data, isLoading, isError } = useGetAllProductsQuery({
    ...filters,
    offset,
  });

  const onPageChange = (page: number) => {
    setOffset(page + 1);
  };

  return (
    <>
      <main className="px-64 mt-10">
        <Filters setOffset={setOffset} setFilters={setFilters} />
        <QueryHandler
          isError={isError}
          isLoading={isLoading}
          errorLabel="При загрузке товаров произошла ошибка"
        />
        {!data?.Pagination.data && !isError && !isLoading && (
          <div className="w-full flex justify-center flex-col items-center gap-y-9 mt-24 text-3xl font-bold text-main-dark-blue">
            По указанному запросу товаров не найдено
          </div>
        )}
        <ProductsList products={data?.Pagination.data ?? []} />
        {!!data?.Pagination && !!data?.Pagination.total && (
          <div className="mt-10 mb-20">
            <Pagination
              countPage={data.Pagination.total}
              currentPage={offset - 10}
              onChange={onPageChange}
            />
          </div>
        )}
      </main>
    </>
  );
};
