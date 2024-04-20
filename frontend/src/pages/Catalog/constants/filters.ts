import { optionValidation } from "@/shared/yup";
import * as yup from "yup";

export const filtersFormSchema = yup.object().shape({
  category: optionValidation,
  subcategory: optionValidation,
  sort: optionValidation,
  supplier: optionValidation,
});

export type FiltersFormValues = yup.InferType<typeof filtersFormSchema>;
