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
import searchIcon from "@/assets/icons/search.svg";
import Image from "next/image";

type Props = {
  setFilters: (data: IGetAllProductsRequest) => void;
  setOffset: (page: number) => void;
};

export const Filters: FC<Props> = ({ setFilters, setOffset }) => {
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
    setValue("subcategory", null);
  }, [category?.id]);

  const onSubmit = (data: FiltersFormValues) => {
    setOffset(1);
    setFilters(transformFormFiltersToRequest(data));
  };

  useEffect(() => {
    setTimeout(() => {
      const subscription = formApi.watch(() =>
        formApi.handleSubmit(onSubmit)()
      );
      return () => subscription.unsubscribe();
    }, 300);
  }, [formApi.handleSubmit, formApi.watch]);

  const formReset = () => {
    setOffset(1);
    setFilters({});
    reset();
  };

  return (
    <section>
      <p className="text-main-black font-semibold text-4xl mb-8">Все товары</p>
      <FormProvider {...formApi}>
        <form
          className="flex justify-between h-[110px] border-b-[1px] border-main-dark-gray mb-8"
          onSubmit={handleSubmit(onSubmit)}
        >
          <div className="flex gap-x-6">
            <FormSelect
              name="category"
              isClearable
              width="w-[160px]"
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
              width="w-[225px]"
              disable={!category?.id}
              placeholder={category?.id ? "Выберите" : "Выберите категорию"}
              name="subcategory"
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
              width="w-[170px]"
              label="Сортировка товаров"
              options={[
                { id: 1, name: "По возрастанию", filter: "asc" } as IOption,
                { id: 2, name: "По убыванию", filter: "desc" } as IOption,
              ]}
            />
            <FormSelect
              isClearable
              name="supplier"
              width="w-[170px]"
              label="Выбор поставщика"
              options={[]}
            />
          </div>
          <div className="self-center flex gap-x-3 pb-[15px]">
            <Button
              className="bg-main-dark-green !text-[16px] !py-[8px] !px-[12px] !font-light"
              type="submit"
            >
              <div className="flex items-center gap-x-1">
                <Image src={searchIcon} height={15} width={15} alt="" />
                Найти
              </div>
            </Button>
            <Button
              type="button"
              className="!text-main-dark-green !text-[16px] !py-[8px] !px-[12px] bg-main-light-gray !text-main-black !font-light"
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
