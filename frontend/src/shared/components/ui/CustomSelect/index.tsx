import Select, { GroupBase, Props } from "react-select";
import { IOption } from "./select.types";
import { getSelectStyles } from "./stylesSelect";
import { useId } from "react";

type SelectProps = {
  label?: string;
};

export const CustomSelect = <TData extends IOption = IOption>({
  options,
  label,
  isLoading,
  ...props
}: Props<TData, false, GroupBase<TData>> & SelectProps) => {
  const customSelectStyles = getSelectStyles<TData>();

  return (
    <div>
      <p className="text-main-blue-gray text-base mb-2">{label}</p>
      <Select
        instanceId={useId()}
        options={options}
        isMulti={false}
        isDisabled={isLoading}
        isSearchable={!isLoading}
        isLoading={isLoading}
        getOptionLabel={(label) => label.name}
        getOptionValue={(value) => value.id.toString()}
        placeholder={isLoading ? "Загрузка" : "Выбрать"}
        noOptionsMessage={() => "Список пуст"}
        loadingMessage={() => "Загрузка"}
        isClearable={false}
        styles={customSelectStyles}
        {...props}
      />
    </div>
  );
};
