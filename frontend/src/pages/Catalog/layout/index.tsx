import {
  IGetAllProductsRequest,
  useGetAllProductsQuery,
} from "@/entities/products";
import { ProductsList } from "../components/ProductsList";
import { QueryHandler } from "@/widgets/QueryHandler";
import { Filters } from "../components/Filters";
import { useState } from "react";
import { Pagination } from "@/shared/components/ui/Pagination";
import { AdBlock } from "../components/AdBlock";
import { Footer } from "../components/Footer";
import { useGetAllCategoriesQuery } from "@/entities/categories";

export const CatalogPageLayout = () => {
  const [filters, setFilters] = useState<IGetAllProductsRequest>({});
  const [page, setPage] = useState<number>(1);
  const { data, isLoading, isError } = useGetAllProductsQuery({
    ...filters,
    page,
  });
  const { data: categories, isLoading: categoriesIsLoading } =
    useGetAllCategoriesQuery();

  const onPageChange = (page: number) => {
    setPage(page + 1);
  };

  return (
    <>
      <main className="px-64 mt-10">
        <div className="px-[190px]">
          <AdBlock />
          <div className="mt-7">
            <Filters setOffset={setPage} setFilters={setFilters} />
          </div>
          <QueryHandler
            isError={isError}
            isLoading={isLoading}
            errorLabel="При загрузке товаров произошла ошибка"
          />
          {!data?.Pagination.data && !isError && !isLoading && (
            <div className="w-full flex justify-center flex-col items-center gap-y-9 mt-24 text-3xl font-semibold text-main-dark-green">
              По указанному запросу товаров не найдено
            </div>
          )}
          <ProductsList products={data?.Pagination.data ?? []} />
        </div>
        {!!data?.Pagination && !!data?.Pagination.total_pages && (
          <div className="mt-10 mb-20 flex justify-center">
            <Pagination
              countPage={data.Pagination.total_pages - 1}
              currentPage={page}
              onChange={onPageChange}
            />
          </div>
        )}
      </main>
      <div className="mt-20">
        <Footer
          isLoading={categoriesIsLoading}
          categories={categories?.data ?? []}
        />
      </div>
    </>
  );
};
