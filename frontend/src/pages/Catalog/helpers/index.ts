import { IGetAllProductsRequest } from "@/entities/products";
import { FiltersFormValues } from "../constants/filters";

export const transformFormFiltersToRequest = (
  data: FiltersFormValues
): IGetAllProductsRequest => {
  return {
    category_id: data.category?.id,
    subcategory_id: data.subcategory?.id,
    sort_price: data.sort?.filter,
  };
};
