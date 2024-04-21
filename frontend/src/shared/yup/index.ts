import * as yup from "yup";

export const optionValidation = yup
  .object()
  .shape({
    id: yup.number().required(),
    name: yup.string().trim().required(),
  })
  .nullable();

export const optionValidationRequired = yup
  .object()
  .shape({
    id: yup.number().required(),
    name: yup.string().required(),
  })
  .required()
  .nullable()
  .test("nullable", "Выберите", (value) => value !== null);

export const emailValidationRequired = yup
  .string()
  .email("Значение “Email” не является правильным email адресом.")
  .required("Заполните поле")
  .nullable()
  .test("nullable", "Заполните поле", (value) => value !== null);

export const numberValidationRequired = yup
  .number()
  .required("Заполните поле")
  .nullable()
  .test("nullable", "Заполните поле", (value) => value !== null);

export const stringValidationRequired = yup
  .string()
  .required("Заполните поле")
  .nullable()
  .test("nullable", "Заполните поле", (value) => value !== null);
