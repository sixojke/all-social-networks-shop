import { GroupBase, StylesConfig } from "react-select";
import { IOption } from "./select.types";

export const getSelectStyles = <
  TData extends IOption = IOption
>(): StylesConfig<TData, boolean, GroupBase<TData>> => {
  return {
    container: (baseStyles) => ({
      ...baseStyles,
      width: "100%",
    }),
    control: (baseStyles) => ({
      ...baseStyles,
      height: 32,
      caretColor: "#018476",
      border: "none",
      borderRadius: "10px",
      cursor: "pointer",
      boxShadow: "none",
      color: "#4F4F4F",
      backgroundColor: "#EDEDED",
      "&:hover": {
        borderColor: "",
        color: "#999DA6",
      },
    }),
    singleValue: (baseStyles) => ({
      ...baseStyles,
      color: "#4F4F4F",
      fontSize: 14,
    }),
    input: (baseStyles) => ({
      ...baseStyles,
      color: "#4F4F4F",
    }),
    menu: (baseStyles) => ({
      ...baseStyles,
      zIndex: 10,
      border: "2px solid #E6EDFF",
      borderRadius: "7px",
      boxShadow: "none",
    }),
    noOptionsMessage: (baseStyles) => ({
      ...baseStyles,
      color: "#4F4F4F",
      boxShadow: "none",
      fontWeight: 500,
    }),
    menuList: (baseStyles) => ({
      ...baseStyles,
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
      rowGap: "0",
      padding: "10px 0",
      color: "#4F4F4F",
    }),
    multiValue: (baseStyles) => ({
      ...baseStyles,
      backgroundColor: "#072659",
      color: "#4F4F4F",
    }),
    loadingIndicator: () => ({
      display: "none",
    }),
    placeholder: (baseStyles) => ({
      ...baseStyles,
      color: "#4F4F4F",
      fontSize: 14,
    }),
    multiValueLabel: (baseStyles) => ({
      ...baseStyles,
      color: "4F4F4F",
      fontSize: 14,
    }),
    multiValueRemove: (baseStyles) => ({
      ...baseStyles,
      padding: 0,
      "&:hover": {
        backgroundColor: "white",
        borderRadius: "10px",
        color: "#4F4F4F",
      },
    }),
    dropdownIndicator: (baseStyles) => ({
      ...baseStyles,
      height: "100%",
      borderRadius: "10px",
      color: "#4F4F4F",
      "&:hover": {
        color: "3485FE",
        cursor: "pointer",
      },
    }),
    indicatorSeparator: () => ({}),
    clearIndicator: (baseStyles) => ({
      ...baseStyles,
      paddingRight: 0,
      height: "100%",
      color: "#4F4F4F",
      cursor: "pointer",
      "&:hover": {
        color: "#999DA6",
      },
    }),
    option: (baseStyles) => ({
      ...baseStyles,
      cursor: "pointer",
      width: "85%",
      color: "#4F4F4F",
      backgroundColor: "white",
      fontWeight: 700,
      borderBottom: "1px solid #999DA6",
      fontSize: 14,
      transitionDuration: "0.3s",
      "&:last-child": {
        borderBottom: "1px solid white",
      },
      "&:hover": {
        borderRadius: "5px",
        borderBottom: "1px solid #262626",
        color: "white",
        backgroundColor: "#262626",
      },
    }),
  };
};
