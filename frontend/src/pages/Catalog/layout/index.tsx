import { useGetAllProductsQuery } from "@/entities/products";
import { HeaderLayout } from "@/widgets/Header";
import { ProductsList } from "../components/ProductsList";
import { Loader } from "@/shared/components/ui/Loader";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { useRouter } from "next/navigation";
import { QueryHandler } from "@/widgets/QueryHandler";
import { Filters } from "../components/Filters";

export const CatalogPageLayout = () => {
  const { data, isLoading, isError } = useGetAllProductsQuery({});

  return (
    <>
      <HeaderLayout />
      <main className="px-64 mt-10">
        <Filters />
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
