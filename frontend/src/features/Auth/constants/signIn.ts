import { stringValidationRequired } from "@/shared/yup";
import * as yup from "yup";

export const signInSchema = yup.object().shape({
  login: stringValidationRequired,
  password: stringValidationRequired,
});

export type SignInFormValues = yup.InferType<typeof signInSchema>;
