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
