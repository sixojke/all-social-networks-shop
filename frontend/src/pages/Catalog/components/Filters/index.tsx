import { Button } from "@/shared/components/ui/Buttons/Button";
import { CustomSelect } from "@/shared/components/ui/CustomSelect";
import { yupResolver } from "@hookform/resolvers/yup";
import { FormProvider, useForm } from "react-hook-form";
import { FiltersFormValues, filtersFormSchema } from "../../constants/filters";
import { FormSelect } from "@/shared/components/common/Form/FormSelect";
import { useGetAllCategoriesQuery } from "@/entities/categories";
import { useGetAllSubcategoriesQuery } from "@/entities/subcategories";
import { IGetAllProductsRequest } from "@/entities/products";
import { FC, useEffect } from "react";
import { IOption } from "@/shared/components/ui/CustomSelect/select.types";
import { transformFormFiltersToRequest } from "../../helpers";

type Props = {
  setFilters: (data: IGetAllProductsRequest) => void;
};

export const Filters: FC<Props> = ({ setFilters }) => {
  const formApi = useForm<FiltersFormValues>({
    mode: "onChange",
    defaultValues: {
      category: null,
      sort: null,
      subcategory: null,
      supplier: null,
    },
    resolver: yupResolver(filtersFormSchema),
  });
  const { handleSubmit, watch, reset, setValue } = formApi;

  const category = watch("category");

  const { data: categories, isLoading: categoriesIsLoading } =
    useGetAllCategoriesQuery();

  const { data: subcategories, isLoading: subcategoriesIsLoading } =
    useGetAllSubcategoriesQuery(
      { category_id: watch("category")?.id as number },
      { skip: !category?.id }
    );

  useEffect(() => {
    if (category === null) {
      setValue("subcategory", null);
    }
  }, [category]);

  const onSubmit = (data: FiltersFormValues) => {
    setFilters(transformFormFiltersToRequest(data));
  };

  const formReset = () => {
    setFilters({});
    reset();
  };

  return (
    <section>
      <p className="text-main-black font-semibold text-4xl mb-8">Все товары</p>
      <FormProvider {...formApi}>
        <form
          className="flex justify-between pb-12 border-b-2 border-main-blue-gray mb-8 h-40"
          onSubmit={handleSubmit(onSubmit)}
        >
          <div className="flex gap-x-8">
            <FormSelect
              name="category"
              isClearable
              width="w-[230px]"
              label="Категория"
              isLoading={categoriesIsLoading}
              options={
                categories?.data?.map((item) => ({
                  id: item.id,
                  name: item.name,
                })) ?? []
              }
            />
            <FormSelect
              isClearable
              disable={!category?.id}
              placeholder={category?.id ? "Выберите" : "Выберите категорию"}
              name="subcategory"
              width="w-[290px]"
              label="Подкатегория"
              isLoading={subcategoriesIsLoading}
              options={
                subcategories?.data?.map((item) => ({
                  id: item.id,
                  name: item.name,
                })) ?? []
              }
            />
            <FormSelect
              isClearable
              name="sort"
              width="w-[230px]"
              label="Сортировка товаров по цене"
              options={[
                { id: 1, name: "По возрастанию", filter: "asc" } as IOption,
                { id: 2, name: "По убыванию", filter: "desc" } as IOption,
              ]}
            />
            <FormSelect
              isClearable
              name="supplier"
              width="w-[230px]"
              label="Выбор поставщика"
              options={[]}
            />
          </div>
          <div className="self-center flex gap-x-10">
            <Button className="bg-main-dark-blue" type="submit">
              Найти
            </Button>
            <Button
              className="!text-main-dark-blue bg-main-light-blue"
              onClick={formReset}
            >
              Очистить
            </Button>
          </div>
        </form>
      </FormProvider>
    </section>
  );
};
