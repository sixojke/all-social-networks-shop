import { useGetAllProductsQuery } from "@/entities/products";
import { Button } from "@/shared/components/ui/Buttons/Button";
import { ButtonText } from "@/shared/components/ui/Buttons/ButtonText";
import { CustomSelect } from "@/shared/components/ui/CustomSelect";
import { Input } from "@/shared/components/ui/Input";
import { HeaderLayout } from "@/widgets/Header";

export const CatalogPageLayout = () => {
  const { data } = useGetAllProductsQuery({ limit: 10, offset: 10 });
  console.log(data);

  return (
    <>
      <HeaderLayout />
    </>
  );
};
