import {
  emailValidationRequired,
  stringValidationRequired,
} from "@/shared/yup";
import * as yup from "yup";

export const signUpSchema = yup.object().shape({
  login: stringValidationRequired,
  email: emailValidationRequired,
  password: stringValidationRequired,
});

export type SignUpFormValues = yup.InferType<typeof signUpSchema>;
