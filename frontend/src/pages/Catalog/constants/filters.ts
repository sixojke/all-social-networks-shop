import { optionValidation } from "@/shared/yup";
import * as yup from "yup";

export const filtersFormSchema = yup.object().shape({
  category: optionValidation,
  subcategory: optionValidation,
  sort: yup
    .object()
    .shape({
      id: yup.number().required(),
      name: yup.string().trim().required(),
      filter: yup.mixed<"asc" | "desc">().oneOf(["asc", "desc"]).required(),
    })
    .nullable(),
  supplier: optionValidation,
});

export type FiltersFormValues = yup.InferType<typeof filtersFormSchema>;
