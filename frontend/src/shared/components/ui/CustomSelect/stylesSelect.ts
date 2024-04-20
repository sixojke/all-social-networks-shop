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
      minHeight: 34,
      padding: "5px 10px",
      caretColor: "#3485FE",
      border: "none",
      borderRadius: "10px",
      cursor: "pointer",
      boxShadow: "none",
      color: "#3485FE",
      backgroundColor: "#E6F0FF",
      "&:hover": {
        borderColor: "",
        color: "#999DA6",
      },
    }),
    singleValue: (baseStyles) => ({
      ...baseStyles,
      color: "#3485FE",
      fontSize: 18,
    }),
    input: (baseStyles) => ({
      ...baseStyles,
      color: "#3485FE",
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
      color: "#3D73FE",
      boxShadow: "none",
      fontWeight: 500,
    }),
    menuList: (baseStyles) => ({
      ...baseStyles,
      display: "flex",
      flexDirection: "column",
      alignItems: "center",
      rowGap: "10px",
      padding: "10px 0",
      color: "#3485FE",
    }),
    multiValue: (baseStyles) => ({
      ...baseStyles,
      backgroundColor: "#072659",
      color: "3485FE",
    }),
    loadingIndicator: () => ({
      display: "none",
    }),
    placeholder: (baseStyles) => ({
      ...baseStyles,
      color: "3485FE",
      fontSize: 18,
    }),
    multiValueLabel: (baseStyles) => ({
      ...baseStyles,
      color: "3485FE",
      fontSize: 18,
    }),
    multiValueRemove: (baseStyles) => ({
      ...baseStyles,
      padding: 0,
      "&:hover": {
        backgroundColor: "white",
        borderRadius: "10px",
        color: "#3485FE",
      },
    }),
    dropdownIndicator: (baseStyles) => ({
      ...baseStyles,
      height: "100%",
      borderRadius: "10px",
      color: "#3485FE",
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
      color: "#3485FE",
      cursor: "pointer",
      "&:hover": {
        color: "#999DA6",
      },
    }),
    option: (baseStyles) => ({
      ...baseStyles,
      cursor: "pointer",
      width: "85%",
      color: "#3485FE",
      backgroundColor: "white",
      fontWeight: 700,
      borderBottom: "1px solid #999DA6",
      fontSize: 18,
      transitionDuration: "0.3s",
      "&:last-child": {
        borderBottom: "1px solid white",
      },
      "&:hover": {
        borderRadius: "5px",
        borderBottom: "1px solid #3485FE",
        color: "white",
        backgroundColor: "#3485FE",
      },
    }),
  };
};
