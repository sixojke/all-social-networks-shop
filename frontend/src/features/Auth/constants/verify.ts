import {
  numberValidationRequired,
  stringValidationRequired,
} from "@/shared/yup";
import * as yup from "yup";

export const verifySchema = yup.object().shape({
  code: stringValidationRequired,
  id: numberValidationRequired,
});

export type VerifyFormValues = yup.InferType<typeof verifySchema>;
