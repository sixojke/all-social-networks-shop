import {
  emailValidationRequired,
  stringValidationRequired,
} from "@/shared/yup";
import * as yup from "yup";

export const signUpSchema = yup.object().shape({
  login: stringValidationRequired.min(
    6,
    "Значение “Логин” должно содержать минимум 6 символов. "
  ),
  email: emailValidationRequired,
  password: stringValidationRequired.min(
    8,
    "Значение “Пароль” должно содержать минимум 8 символов."
  ),
});

export type SignUpFormValues = yup.InferType<typeof signUpSchema>;
